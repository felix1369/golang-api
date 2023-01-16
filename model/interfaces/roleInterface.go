package interfaces

import (
	"context"

	"github.com/felix1369/golang-api/model/entities"
)

// RoleRepository represent the role's infrastructure contract
type RoleRepository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []entities.Role, err error)
	GetByID(ctx context.Context, id uint) (entities.Role, error)
	GetByName(ctx context.Context, title string) (entities.Role, error)
	Update(ctx context.Context, ar *entities.Role) error
	Store(ctx context.Context, a *entities.Role) error
	Delete(ctx context.Context, id uint) error
}

// RoleUseCase represent the role's application contract
type RoleUseCase interface {
	Fetch(ctx context.Context) ([]entities.Role, error)
	GetByID(ctx context.Context, id uint) (entities.Role, error)
	Update(ctx context.Context, ar *entities.Role) error
	GetByName(ctx context.Context, title string) (entities.Role, error)
	Store(context.Context, *entities.Role) error
	Delete(ctx context.Context, id uint) error
}
