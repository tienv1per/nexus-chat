# Chat V1 PostgreSQL Query Paths

Phase 3 keeps the query paths explicit so later Go adapters can stay simple and
use parameterized SQL without ORM-generated surprises.

## Membership Check

Used by `SendMessage`, media access checks, and conversation reads.

```sql
SELECT 1
FROM conversation_members
WHERE conversation_id = $1
  AND user_id = $2
  AND left_at IS NULL;
```

Relevant access path:

- Primary key: `(conversation_id, user_id)`
- Active user lookup: `idx_conversation_members_user_active`

## User Conversation List

Start from membership because the current user is the filter.

```sql
SELECT
    c.id,
    c.type,
    c.title,
    c.description,
    c.system_tag,
    cm.role,
    cm.last_seen_sequence
FROM conversation_members cm
JOIN conversations c ON c.id = cm.conversation_id
WHERE cm.user_id = $1
  AND cm.left_at IS NULL
ORDER BY c.updated_at DESC
LIMIT $2;
```

Relevant access path:

- `idx_conversation_members_user_active`
- `conversations` primary key

## Message History By Sequence Cursor

The UI loads a bounded page before a sequence cursor. The caller should enforce a
maximum limit, such as `50`.

```sql
SELECT
    id,
    conversation_id,
    sender_id,
    client_msg_id,
    sequence,
    message_kind,
    body,
    media_id,
    status,
    created_at
FROM messages
WHERE conversation_id = $1
  AND sequence < $2
ORDER BY sequence DESC
LIMIT $3;
```

Recent messages use the same index without the `sequence < $2` predicate:

```sql
SELECT *
FROM messages
WHERE conversation_id = $1
ORDER BY sequence DESC
LIMIT $2;
```

Relevant access path:

- Unique ordering: `idx_messages_conversation_sequence_unique`
- Timeline page: `idx_messages_history`

## Idempotent Send Retry

Used before inserting a duplicate message for the same sender retry.

```sql
SELECT id, sequence, status
FROM messages
WHERE conversation_id = $1
  AND sender_id = $2
  AND client_msg_id = $3;
```

Relevant access path:

- `idx_messages_client_msg_id_unique`

## Media Owner Lookup

Used by upload ownership and later authorization checks.

```sql
SELECT id, owner_user_id, original_filename, mime_type, size_bytes, storage_path
FROM media_objects
WHERE owner_user_id = $1
ORDER BY created_at DESC
LIMIT $2;
```

Relevant access path:

- `idx_media_owner_created`

## Delivery Status Debugging

Used by status consumers and admin/debug pages.

```sql
SELECT message_id, recipient_id, status, delivered_at, read_at, updated_at
FROM message_deliveries
WHERE recipient_id = $1
  AND status = $2
ORDER BY updated_at DESC
LIMIT $3;
```

Relevant access path:

- `idx_message_deliveries_recipient_status`
