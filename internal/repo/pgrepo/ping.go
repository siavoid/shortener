package pgrepo

import "context"

func (p *PostgresRepo) Ping(ctx context.Context) error {
	return p.db.Pool.Ping(ctx)
}
