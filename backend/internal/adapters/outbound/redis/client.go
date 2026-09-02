// Package redis contains Redis outbound adapters for presence, sessions, and counters.
package redis

// Client is the Redis adapter boundary for presence/session state and sequence counters.
//
// Phase 6 wires this type to a concrete Redis client when WebSocket sessions land.
type Client struct {
	addr string
}

// NewClient records Redis connection configuration without opening a global client.
func NewClient(addr string) *Client {
	return &Client{
		addr: addr,
	}
}

// Addr returns the configured Redis address for composition smoke checks.
func (c *Client) Addr() string {
	if c == nil {
		return ""
	}

	return c.addr
}
