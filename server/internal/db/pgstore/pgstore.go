package pgstore

import (
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib" // Postgresql driver
)

type PGStore struct {
	db *sql.DB
}

func NewPgStore(dsn string) (*PGStore, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	ls := &PGStore{
		db: db,
	}
	return ls, nil
}

func (pgs *PGStore) Close() {
	pgs.db.Close()
}
