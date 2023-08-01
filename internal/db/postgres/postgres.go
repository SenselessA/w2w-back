package postgres

import (
	"github.com/SenselessA/w2w_backend/internal/db"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Open(cfg *db.ConfigDB) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", cfg.String())
	//return pgx.Connect(context.Background(), cfg.String())
}
