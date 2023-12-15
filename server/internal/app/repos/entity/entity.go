package entity

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Entity struct {
	ID              uuid.UUID
	Filename        string
	Name            string
	Description     string
	Longitude       float64
	Latitude        float64
	Height          float64
	DescriptionJson struct{}
	CellID          uint64
	Geohash         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type StoreEntityInterface interface {
	Create(ctx context.Context, entity Entity) (*Entity, error)
	Read(ctx context.Context, ID uuid.UUID) (*Entity, error)
	Neighbours(ctx context.Context, CellID uint64) (chan Entity, error)
}

type Entities struct {
	entityStore StoreEntityInterface
}

func NewEntities(entityStore StoreEntityInterface) *Entities {
	return &Entities{
		entityStore: entityStore,
	}
}
