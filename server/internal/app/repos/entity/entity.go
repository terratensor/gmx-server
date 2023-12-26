package entity

import (
	"context"
	"fmt"
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
	DescriptionJson map[string]interface{}
	CellID          uint64
	Geohash         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// StoreEntityInterface
// Create(ctx context.Context, entity Entity) (*Entity, error)
// Read(ctx context.Context, ID uuid.UUID) (*Entity, error)
// Neighbours(ctx context.Context, cellID uint64) (chan Entity, error)
type StoreEntityInterface interface {
	ReadByCellID(ctx context.Context, cellID uint64) (*Entity, error)
}

type Entities struct {
	entityStore StoreEntityInterface
}

func NewEntities(entityStore StoreEntityInterface) *Entities {
	return &Entities{
		entityStore: entityStore,
	}
}

func (es *Entities) Read(ctx context.Context, cellID uint64) (*Entity, error) {
	entity, err := es.entityStore.ReadByCellID(ctx, cellID)
	if err != nil {
		return nil, fmt.Errorf("read entity error %w", err)
	}
	return entity, err
}
