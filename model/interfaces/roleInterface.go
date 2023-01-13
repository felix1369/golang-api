package interfaces

import (
	"context"

	"github.com/felix1369/golang-api/model/entity"
)

// RoleRepository represent the role's infrastructure contract
type RoleRepository interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []entity.Role, err error)
	GetByID(ctx context.Context, id int64) (entity.Role, error)
	GetByName(ctx context.Context, title string) (entity.Role, error)
	Update(ctx context.Context, ar *entity.Role) error
	Store(ctx context.Context, a *entity.Role) error
	Delete(ctx context.Context, id int64) error
}

// RoleUseCase represent the role's application contract
type RoleUseCase interface {
	Fetch(ctx context.Context) ([]entity.Role, error)
	GetByID(ctx context.Context, id int64) (entity.Role, error)
	Update(ctx context.Context, ar *entity.Role) error
	GetByName(ctx context.Context, title string) (entity.Role, error)
	Store(context.Context, *entity.Role) error
	Delete(ctx context.Context, id int64) error
}
