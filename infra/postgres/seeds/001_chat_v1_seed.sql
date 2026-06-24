BEGIN;

INSERT INTO users (
    id,
    username,
    display_name,
    email,
    initials,
    avatar_color,
    role,
    last_seen_at,
    created_at,
    updated_at
) VALUES
    ('usr_alice', 'alice', 'Alice Tran', 'alice@nexus.local', 'AT', 'teal', 'Owner', NULL, '2026-06-24T08:00:00+07:00', '2026-06-24T08:00:00+07:00'),
    ('usr_bob', 'bob', 'Bob Nguyen', 'bob@nexus.local', 'BN', 'indigo', 'Member', '2026-06-24T08:32:00+07:00', '2026-06-24T08:00:00+07:00', '2026-06-24T08:32:00+07:00'),
    ('usr_charlie', 'charlie', 'Charlie Le', 'charlie@nexus.local', 'CL', 'orange', 'Test', '2026-06-23T22:10:00+07:00', '2026-06-24T08:00:00+07:00', '2026-06-23T22:10:00+07:00')
ON CONFLICT (id) DO UPDATE SET
    username = EXCLUDED.username,
    display_name = EXCLUDED.display_name,
    email = EXCLUDED.email,
    initials = EXCLUDED.initials,
    avatar_color = EXCLUDED.avatar_color,
    role = EXCLUDED.role,
    last_seen_at = EXCLUDED.last_seen_at,
    updated_at = EXCLUDED.updated_at;

INSERT INTO conversations (
    id,
    type,
    title,
    description,
    system_tag,
    created_by,
    created_at,
    updated_at
) VALUES
    ('conv_bob', 'ONE_TO_ONE', 'Bob Nguyen', '1:1 product implementation thread', 'Direct', 'usr_alice', '2026-06-24T10:30:00+07:00', '2026-06-24T10:42:00+07:00'),
    ('conv_infra', 'GROUP', 'Infra Review Room', 'Kafka, Redis, Postgres, and WebSocket diagnostics', 'Group', 'usr_alice', '2026-06-24T09:00:00+07:00', '2026-06-24T09:18:00+07:00'),
    ('conv_media', 'GROUP', 'Media QA', 'Upload previews, MIME checks, and local storage', 'Media', 'usr_alice', '2026-06-23T18:00:00+07:00', '2026-06-23T18:22:00+07:00')
ON CONFLICT (id) DO UPDATE SET
    type = EXCLUDED.type,
    title = EXCLUDED.title,
    description = EXCLUDED.description,
    system_tag = EXCLUDED.system_tag,
    created_by = EXCLUDED.created_by,
    updated_at = EXCLUDED.updated_at;

INSERT INTO conversation_members (
    conversation_id,
    user_id,
    role,
    joined_at,
    left_at,
    last_seen_sequence
) VALUES
    ('conv_bob', 'usr_alice', 'ADMIN', '2026-06-24T10:30:00+07:00', NULL, 41),
    ('conv_bob', 'usr_bob', 'MEMBER', '2026-06-24T10:30:00+07:00', NULL, 38),
    ('conv_infra', 'usr_alice', 'ADMIN', '2026-06-24T09:00:00+07:00', NULL, 57),
    ('conv_infra', 'usr_bob', 'MEMBER', '2026-06-24T09:00:00+07:00', NULL, 55),
    ('conv_infra', 'usr_charlie', 'MEMBER', '2026-06-24T09:00:00+07:00', NULL, 49),
    ('conv_media', 'usr_alice', 'ADMIN', '2026-06-23T18:00:00+07:00', NULL, 16),
    ('conv_media', 'usr_charlie', 'MEMBER', '2026-06-23T18:00:00+07:00', NULL, 14)
ON CONFLICT (conversation_id, user_id) DO UPDATE SET
    role = EXCLUDED.role,
    left_at = EXCLUDED.left_at,
    last_seen_sequence = EXCLUDED.last_seen_sequence;

INSERT INTO media_objects (
    id,
    owner_user_id,
    original_filename,
    mime_type,
    size_bytes,
    size_label,
    storage_path,
    created_at
) VALUES
    ('media_schema', 'usr_alice', 'message-timeline-schema.sql', 'text/sql', 9216, '9 KB', 'data/uploads/media_schema/original', '2026-06-24T09:12:00+07:00')
ON CONFLICT (id) DO UPDATE SET
    owner_user_id = EXCLUDED.owner_user_id,
    original_filename = EXCLUDED.original_filename,
    mime_type = EXCLUDED.mime_type,
    size_bytes = EXCLUDED.size_bytes,
    size_label = EXCLUDED.size_label,
    storage_path = EXCLUDED.storage_path;

INSERT INTO messages (
    id,
    conversation_id,
    sender_id,
    client_msg_id,
    sequence,
    message_kind,
    body,
    media_id,
    status,
    created_at,
    updated_at
) VALUES
    ('msg_39', 'conv_bob', 'usr_bob', NULL, 39, 'TEXT', 'The conversation list feels stronger when unread and presence are visible at the same time.', NULL, 'DELIVERED', '2026-06-24T10:36:00+07:00', '2026-06-24T10:36:00+07:00'),
    ('msg_40', 'conv_bob', 'usr_alice', NULL, 40, 'TEXT', 'Agreed. I also want sequence numbers exposed so reconnect behavior is obvious during demos.', NULL, 'DELIVERED', '2026-06-24T10:39:00+07:00', '2026-06-24T10:39:00+07:00'),
    ('msg_41', 'conv_bob', 'usr_bob', NULL, 41, 'TEXT', 'Let''s keep the ACK payload visible in the UI.', NULL, 'DELIVERED', '2026-06-24T10:42:00+07:00', '2026-06-24T10:42:00+07:00'),
    ('msg_54', 'conv_infra', 'usr_charlie', NULL, 54, 'TEXT', 'Redis presence TTL is still the fast path. Postgres only answers last-seen fallback now.', NULL, 'DELIVERED', '2026-06-24T09:04:00+07:00', '2026-06-24T09:04:00+07:00'),
    ('msg_55', 'conv_infra', 'usr_alice', NULL, 55, 'FILE', 'Attached the Postgres timeline migration sketch for review.', 'media_schema', 'DELIVERED', '2026-06-24T09:12:00+07:00', '2026-06-24T09:12:00+07:00'),
    ('msg_56', 'conv_infra', 'usr_bob', NULL, 56, 'TEXT', 'The UI can sort by sequence and keep realtime arrival as just another insert path.', NULL, 'DELIVERED', '2026-06-24T09:16:00+07:00', '2026-06-24T09:16:00+07:00'),
    ('msg_57', 'conv_infra', 'usr_alice', NULL, 57, 'TEXT', 'Postgres history is now the durable recovery path.', NULL, 'SENT', '2026-06-24T09:18:00+07:00', '2026-06-24T09:18:00+07:00'),
    ('msg_15', 'conv_media', 'usr_charlie', NULL, 15, 'TEXT', 'Image previews should not shift the composer.', NULL, 'DELIVERED', '2026-06-23T18:22:00+07:00', '2026-06-23T18:22:00+07:00')
ON CONFLICT (id) DO UPDATE SET
    conversation_id = EXCLUDED.conversation_id,
    sender_id = EXCLUDED.sender_id,
    client_msg_id = EXCLUDED.client_msg_id,
    sequence = EXCLUDED.sequence,
    message_kind = EXCLUDED.message_kind,
    body = EXCLUDED.body,
    media_id = EXCLUDED.media_id,
    status = EXCLUDED.status,
    updated_at = EXCLUDED.updated_at;

INSERT INTO message_deliveries (
    message_id,
    recipient_id,
    status,
    delivered_at,
    read_at,
    updated_at
) VALUES
    ('msg_39', 'usr_alice', 'DELIVERED', '2026-06-24T10:36:01+07:00', NULL, '2026-06-24T10:36:01+07:00'),
    ('msg_40', 'usr_bob', 'DELIVERED', '2026-06-24T10:39:01+07:00', NULL, '2026-06-24T10:39:01+07:00'),
    ('msg_41', 'usr_alice', 'DELIVERED', '2026-06-24T10:42:01+07:00', NULL, '2026-06-24T10:42:01+07:00'),
    ('msg_54', 'usr_alice', 'DELIVERED', '2026-06-24T09:04:01+07:00', NULL, '2026-06-24T09:04:01+07:00'),
    ('msg_54', 'usr_bob', 'DELIVERED', '2026-06-24T09:04:01+07:00', NULL, '2026-06-24T09:04:01+07:00'),
    ('msg_55', 'usr_bob', 'DELIVERED', '2026-06-24T09:12:01+07:00', NULL, '2026-06-24T09:12:01+07:00'),
    ('msg_55', 'usr_charlie', 'DELIVERED', '2026-06-24T09:12:01+07:00', NULL, '2026-06-24T09:12:01+07:00'),
    ('msg_56', 'usr_alice', 'DELIVERED', '2026-06-24T09:16:01+07:00', NULL, '2026-06-24T09:16:01+07:00'),
    ('msg_56', 'usr_charlie', 'DELIVERED', '2026-06-24T09:16:01+07:00', NULL, '2026-06-24T09:16:01+07:00'),
    ('msg_57', 'usr_bob', 'SENT', NULL, NULL, '2026-06-24T09:18:00+07:00'),
    ('msg_57', 'usr_charlie', 'SENT', NULL, NULL, '2026-06-24T09:18:00+07:00'),
    ('msg_15', 'usr_alice', 'DELIVERED', '2026-06-23T18:22:01+07:00', NULL, '2026-06-23T18:22:01+07:00')
ON CONFLICT (message_id, recipient_id) DO UPDATE SET
    status = EXCLUDED.status,
    delivered_at = EXCLUDED.delivered_at,
    read_at = EXCLUDED.read_at,
    updated_at = EXCLUDED.updated_at;

COMMIT;
