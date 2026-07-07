package store

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)



// func NewRequest() (r *Request) {
//     return &Request{}
// }

// type Endpoint struct {
// 	id string
//     token_hash string
//     name string
//     created_at time.Time
// }
type PostgresDBConnection struct {
	pool *pgxpool.Pool
	ctx *context.Context
}

func NewPostgresDBConnection(p *pgxpool.Pool, c *context.Context) *PostgresDBConnection {
	return &PostgresDBConnection{pool: p, ctx: c}
}

func (db *PostgresDBConnection) Save(r *Request) error {
	query := "insert into requests(method, path, query, headers, body, body_size_bytes, body_truncrated, content_type, source_ip, received_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"

	if _, err := db.pool.Exec(*db.ctx, query,
		r.Method,
		r.Path,
		r.Query,
		r.Headers,
		r.Body,
		r.Body_size_bytes,
		r.Body_truncated,
		r.Content_type,
		r.Source_ip,
		r.Received_at); err != nil {
		return err
	}
	
	return nil
}
