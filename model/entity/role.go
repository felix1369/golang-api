package entity

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name  string `json:"name"`
	Level string `json:"level"`
}

// RoleInfrastructure represent the role's infrastructure contract
type RoleInfrastructure interface {
	Fetch(ctx context.Context, query string, args ...interface{}) (res []Role, err error)
	GetByID(ctx context.Context, id int64) (Role, error)
	GetByName(ctx context.Context, title string) (Role, error)
	Update(ctx context.Context, ar *Role) error
	Store(ctx context.Context, a *Role) error
	Delete(ctx context.Context, id int64) error
}

// RoleApplication represent the role's application contract
type RoleApplication interface {
	Fetch(ctx context.Context) ([]Role, error)
	GetByID(ctx context.Context, id int64) (Role, error)
	Update(ctx context.Context, ar *Role) error
	GetByName(ctx context.Context, title string) (Role, error)
	Store(context.Context, *Role) error
	Delete(ctx context.Context, id int64) error
}
