package starter

import "github.com/terratensor/gmx-server/server/internal/app/repos/entity"

type APIServer interface {
	Start(entities *entity.Entity)
	Stop()
}
