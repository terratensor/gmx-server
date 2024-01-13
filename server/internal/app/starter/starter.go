package starter

import (
	"github.com/terratensor/gmx-server/server/internal/app/repos/entity"
	"github.com/terratensor/gmx-server/server/internal/db/pgstore"
	"github.com/terratensor/gmx-server/server/internal/db/pgstore/pgentity"
)

type App struct {
	ent *entity.Entities
}

type APIServer interface {
	Start(entities *entity.Entity)
	Stop()
}

func NewApp(pgs *pgstore.PGStore) *App {

	pgEnt := pgentity.NewEntities(pgs.DB)

	a := &App{
		ent: entity.NewEntities(pgEnt),
	}

	return a
}
