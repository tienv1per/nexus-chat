#!/usr/bin/env bash
set -euo pipefail

bootstrap_server="${KAFKA_BOOTSTRAP_SERVER:-kafka:29092}"
retention_ms="${KAFKA_RETENTION_MS:-86400000}"

create_topic() {
  local topic="$1"
  local partitions="$2"

  kafka-topics.sh \
    --bootstrap-server "${bootstrap_server}" \
    --create \
    --if-not-exists \
    --topic "${topic}" \
    --partitions "${partitions}" \
    --replication-factor 1 \
    --config "retention.ms=${retention_ms}"
}

create_topic "${KAFKA_TOPIC_MESSAGE_CREATED:-chat.message.created}" 3
create_topic "${KAFKA_TOPIC_MESSAGE_DELIVERED:-chat.message.delivered}" 3

kafka-topics.sh \
  --bootstrap-server "${bootstrap_server}" \
  --describe \
  --topic "${KAFKA_TOPIC_MESSAGE_CREATED:-chat.message.created}"

kafka-topics.sh \
  --bootstrap-server "${bootstrap_server}" \
  --describe \
  --topic "${KAFKA_TOPIC_MESSAGE_DELIVERED:-chat.message.delivered}"
