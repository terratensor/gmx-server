package pgstore

import (
	"context"
	"database/sql"
	"encoding/json"
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

func NewEntities(dsn string) (*Entities, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	ls := &Entities{
		db: db,
	}
	return ls, nil
}

func (es *Entities) Close() {
	es.db.Close()
}

func (es *Entities) ReadByCellID(ctx context.Context, cellID uint64) (*entity.Entity, error) {

	dbe := &DBPgEntity{}

	rows, err := es.db.QueryContext(
		ctx,
		`SELECT id, filename, name, description, longitude, latitude, height, description_json, cell_id, geohash, created_at, updated_at, deleted_at
	FROM db_entities WHERE cell_id = $1`,
		cellID,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&dbe.ID,
			&dbe.Filename,
			&dbe.Name,
			&dbe.Description,
			&dbe.Longitude,
			&dbe.Latitude,
			&dbe.Height,
			&dbe.DescriptionJson,
			&dbe.CellID,
			&dbe.Geohash,
			&dbe.CreatedAt,
			&dbe.UpdatedAt,
			&dbe.DeletedAt,
		); err != nil {
			return nil, err
		}
	}

	var descJsonMap map[string]interface{}
	json.Unmarshal([]byte(dbe.DescriptionJson), &descJsonMap)

	return &entity.Entity{
		ID:              dbe.ID,
		Filename:        dbe.Filename,
		Name:            dbe.Name,
		Description:     dbe.Description,
		Longitude:       dbe.Longitude,
		Latitude:        dbe.Latitude,
		Height:          dbe.Height,
		DescriptionJson: descJsonMap,
		CellID:          dbe.CellID,
		Geohash:         dbe.Geohash,
		CreatedAt:       dbe.CreatedAt,
		UpdatedAt:       dbe.UpdatedAt,
	}, nil
}
