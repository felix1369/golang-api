package usecase

import (
	"context"
	"time"

	"github.com/felix1369/golang-api/model"
	"github.com/felix1369/golang-api/model/entity"
)

type articleUsecase struct {
	articleRepo    entity.RoleInfrastructure
	contextTimeout time.Duration
}

// NewArticleUsecase will create new an articleUsecase object representation of domain.ArticleUsecase interface
func NewArticleUsecase(a entity.RoleInfrastructure, ar entity.Role, timeout time.Duration) entity.RoleApplication {
	return &articleUsecase{
		articleRepo:    a,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) Fetch(c context.Context) (res []entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.Fetch(ctx, "")
	if err != nil {
		return nil, err
	}

	return
}

func (a *articleUsecase) GetByID(c context.Context, id int64) (res entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *articleUsecase) Update(c context.Context, ar *entity.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.articleRepo.Update(ctx, ar)
}

func (a *articleUsecase) GetByName(c context.Context, title string) (res entity.Role, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	res, err = a.articleRepo.GetByName(ctx, title)
	if err != nil {
		return
	}
	return
}

func (a *articleUsecase) Store(c context.Context, m *entity.Role) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, _ := a.GetByName(ctx, m.Name)
	if existedArticle != (entity.Role{}) {
		return model.ErrConflict
	}

	err = a.articleRepo.Store(ctx, m)
	return
}

func (a *articleUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedArticle, err := a.articleRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedArticle == (entity.Role{}) {
		return model.ErrNotFound
	}
	return a.articleRepo.Delete(ctx, id)
}
