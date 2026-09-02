// Package postgres contains PostgreSQL outbound adapters.
package postgres

// Repository is the PostgreSQL adapter boundary for metadata and timeline stores.
//
// Phase 5 wires this type to database/sql or pgx after REST/gRPC contracts land.
type Repository struct {
	dsn string
}

// NewRepository records PostgreSQL connection configuration without opening a global client.
func NewRepository(dsn string) *Repository {
	return &Repository{
		dsn: dsn,
	}
}

// DSN returns the configured connection string for composition smoke checks.
func (r *Repository) DSN() string {
	if r == nil {
		return ""
	}

	return r.dsn
}
