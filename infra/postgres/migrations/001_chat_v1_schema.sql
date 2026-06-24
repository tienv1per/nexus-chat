BEGIN;

DO $$
BEGIN
    CREATE TYPE conversation_type AS ENUM ('ONE_TO_ONE', 'GROUP');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
    CREATE TYPE member_role AS ENUM ('ADMIN', 'MEMBER');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
    CREATE TYPE message_kind AS ENUM ('TEXT', 'IMAGE', 'VIDEO', 'FILE');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

DO $$
BEGIN
    CREATE TYPE delivery_status AS ENUM ('PENDING', 'SENT', 'DELIVERED', 'FAILED');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS users (
    id text PRIMARY KEY CHECK (id <> ''),
    username text NOT NULL UNIQUE CHECK (username <> ''),
    display_name text NOT NULL CHECK (display_name <> ''),
    email text NOT NULL UNIQUE CHECK (email <> ''),
    initials text NOT NULL CHECK (initials <> ''),
    avatar_color text NOT NULL CHECK (avatar_color <> ''),
    role text NOT NULL CHECK (role IN ('Owner', 'Member', 'Test')),
    last_seen_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS conversations (
    id text PRIMARY KEY CHECK (id <> ''),
    type conversation_type NOT NULL,
    title text NOT NULL CHECK (title <> ''),
    description text NOT NULL DEFAULT '',
    system_tag text NOT NULL DEFAULT '',
    created_by text NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS conversation_members (
    conversation_id text NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    user_id text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role member_role NOT NULL DEFAULT 'MEMBER',
    joined_at timestamptz NOT NULL DEFAULT now(),
    left_at timestamptz,
    last_seen_sequence bigint NOT NULL DEFAULT 0 CHECK (last_seen_sequence >= 0),
    PRIMARY KEY (conversation_id, user_id)
);

CREATE TABLE IF NOT EXISTS media_objects (
    id text PRIMARY KEY CHECK (id <> ''),
    owner_user_id text NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    original_filename text NOT NULL CHECK (original_filename <> ''),
    mime_type text NOT NULL CHECK (mime_type <> ''),
    size_bytes bigint NOT NULL CHECK (size_bytes > 0),
    size_label text NOT NULL DEFAULT '',
    storage_path text NOT NULL CHECK (storage_path <> ''),
    created_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS messages (
    id text PRIMARY KEY CHECK (id <> ''),
    conversation_id text NOT NULL REFERENCES conversations(id) ON DELETE CASCADE,
    sender_id text NOT NULL REFERENCES users(id) ON DELETE RESTRICT,
    client_msg_id text,
    sequence bigint NOT NULL CHECK (sequence > 0),
    message_kind message_kind NOT NULL,
    body text NOT NULL DEFAULT '',
    media_id text REFERENCES media_objects(id) ON DELETE SET NULL,
    status delivery_status NOT NULL DEFAULT 'SENT',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CHECK (client_msg_id IS NULL OR client_msg_id <> ''),
    CHECK (message_kind = 'TEXT' OR media_id IS NOT NULL)
);

CREATE TABLE IF NOT EXISTS message_deliveries (
    message_id text NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    recipient_id text NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status delivery_status NOT NULL DEFAULT 'PENDING',
    delivered_at timestamptz,
    read_at timestamptz,
    updated_at timestamptz NOT NULL DEFAULT now(),
    PRIMARY KEY (message_id, recipient_id),
    CHECK (read_at IS NULL OR delivered_at IS NOT NULL)
);

CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);

CREATE INDEX IF NOT EXISTS idx_conversation_members_user
    ON conversation_members (user_id, conversation_id);

CREATE INDEX IF NOT EXISTS idx_conversation_members_user_active
    ON conversation_members (user_id, conversation_id)
    WHERE left_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_conversations_created_by_created
    ON conversations (created_by, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_media_owner_created
    ON media_objects (owner_user_id, created_at DESC);

CREATE UNIQUE INDEX IF NOT EXISTS idx_messages_conversation_sequence_unique
    ON messages (conversation_id, sequence);

CREATE UNIQUE INDEX IF NOT EXISTS idx_messages_client_msg_id_unique
    ON messages (conversation_id, sender_id, client_msg_id)
    WHERE client_msg_id IS NOT NULL;

CREATE INDEX IF NOT EXISTS idx_messages_history
    ON messages (conversation_id, sequence DESC, id);

CREATE INDEX IF NOT EXISTS idx_messages_conversation_created
    ON messages (conversation_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_messages_sender_created
    ON messages (sender_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_message_deliveries_recipient_status
    ON message_deliveries (recipient_id, status, updated_at DESC);

CREATE INDEX IF NOT EXISTS idx_message_deliveries_status
    ON message_deliveries (status, updated_at DESC);

COMMIT;
