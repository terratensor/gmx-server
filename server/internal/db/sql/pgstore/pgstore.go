package pgstore

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/terratensor/gmx-server/server/internal/app/repos/entity"
	"time"
)

var _ entity.StoreEntityInterface = &Entities{}

type DBPgEntity struct {
	ID              uuid.UUID  `db:"id"`
	Filename        string     `db:"filename"`
	Name            string     `db:"name"`
	Description     string     `db:"description"`
	Longitude       float64    `db:"longitude"`
	Latitude        float64    `db:"latitude"`
	Height          float64    `db:"height"`
	DescriptionJson string     `db:"description_json"`
	CellID          uint64     `db:"cell_id"`
	Geohash         string     `db:"geohash"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       time.Time  `db:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at"`
}

type Entities struct {
	db *sql.DB
}
