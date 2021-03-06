package store

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/catcurd/models"
)

type Store interface {
	Get(ctx *gofr.Context) ([]models.Cat, error)
	GetByID(ctx *gofr.Context, id string) (models.Cat, error)
	Create(ctx *gofr.Context, c models.Cat) (models.Cat, error)
	Update(ctx *gofr.Context, c models.Cat) (models.Cat, error)
	Delete(ctx *gofr.Context, id string) error
}
