// Package kafka contains outbound adapters for chat event publication.
package kafka

// Publisher is the Kafka adapter boundary for durable chat events.
//
// Phase 7 wires this type to a concrete Kafka producer after message creation exists.
type Publisher struct {
	brokers []string
}

// NewPublisher records broker configuration without opening a global producer.
func NewPublisher(brokers []string) *Publisher {
	copied := append([]string{}, brokers...)

	return &Publisher{
		brokers: copied,
	}
}

// Brokers returns the configured broker list for composition smoke checks.
func (p *Publisher) Brokers() []string {
	if p == nil {
		return []string{}
	}

	return append([]string{}, p.brokers...)
}
