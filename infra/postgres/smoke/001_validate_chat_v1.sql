DO $$
DECLARE
    expected_table_count integer;
    expected_seed_user_count integer;
    expected_seed_conversation_count integer;
    expected_index_count integer;
BEGIN
    SELECT count(*)
    INTO expected_table_count
    FROM information_schema.tables
    WHERE table_schema = 'public'
      AND table_name IN (
          'users',
          'conversations',
          'conversation_members',
          'media_objects',
          'messages',
          'message_deliveries'
      );

    IF expected_table_count <> 6 THEN
        RAISE EXCEPTION 'expected 6 chat tables, found %', expected_table_count;
    END IF;

    SELECT count(*)
    INTO expected_index_count
    FROM pg_class
    WHERE relkind = 'i'
      AND relname IN (
          'idx_conversation_members_user',
          'idx_conversation_members_user_active',
          'idx_media_owner_created',
          'idx_messages_conversation_sequence_unique',
          'idx_messages_client_msg_id_unique',
          'idx_messages_history',
          'idx_message_deliveries_recipient_status'
      );

    IF expected_index_count <> 7 THEN
        RAISE EXCEPTION 'expected 7 required chat indexes, found %', expected_index_count;
    END IF;

    SELECT count(*)
    INTO expected_seed_user_count
    FROM users
    WHERE username IN ('alice', 'bob', 'charlie');

    IF expected_seed_user_count <> 3 THEN
        RAISE EXCEPTION 'expected seed users alice/bob/charlie, found %', expected_seed_user_count;
    END IF;

    SELECT count(*)
    INTO expected_seed_conversation_count
    FROM conversations
    WHERE id IN ('conv_bob', 'conv_infra');

    IF expected_seed_conversation_count <> 2 THEN
        RAISE EXCEPTION 'expected direct and group seed conversations, found %', expected_seed_conversation_count;
    END IF;

    PERFORM 1
    FROM conversation_members
    WHERE conversation_id = 'conv_bob'
      AND user_id = 'usr_alice'
      AND left_at IS NULL;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'expected usr_alice to be an active member of conv_bob';
    END IF;

    PERFORM 1
    FROM messages
    WHERE conversation_id = 'conv_bob'
      AND sequence = 41
      AND id = 'msg_41';

    IF NOT FOUND THEN
        RAISE EXCEPTION 'expected message history seed msg_41 at sequence 41';
    END IF;
END $$;

SELECT
    c.id AS conversation_id,
    m.id AS message_id,
    m.sequence,
    m.message_kind,
    m.status
FROM messages m
JOIN conversations c ON c.id = m.conversation_id
WHERE m.conversation_id = 'conv_bob'
  AND m.sequence <= 41
ORDER BY m.sequence DESC
LIMIT 3;
