package persistence

import (
	"context"

	"github.com/iota-agency/iota-erp/internal/domain/entities/user"
	"github.com/iota-agency/iota-erp/sdk/composables"
	"github.com/iota-agency/iota-erp/sdk/graphql/helpers"
	"github.com/iota-agency/iota-erp/sdk/service"
)

func NewUserRepository() user.Repository {
	return &GormUserRepository{}
}

type GormUserRepository struct{}

func (g *GormUserRepository) GetPaginated(
	ctx context.Context,
	limit, offset int,
	sortBy []string,
) ([]*user.User, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	q := tx.Limit(limit).Offset(offset)
	q, err := helpers.ApplySort(q, sortBy, &user.User{}) //nolint:exhaustruct
	if err != nil {
		return nil, err
	}
	var users []*user.User
	if err := q.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (g *GormUserRepository) Count(ctx context.Context) (int64, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return 0, service.ErrNoTx
	}
	var count int64
	if err := tx.Model(&user.User{}).Count(&count).Error; err != nil { //nolint:exhaustruct
		return 0, err
	}
	return count, nil
}

func (g *GormUserRepository) GetAll(ctx context.Context) ([]*user.User, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var users []*user.User
	if err := tx.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (g *GormUserRepository) GetByID(ctx context.Context, id int64) (*user.User, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var entity user.User
	if err := tx.Preload("Roles").First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (g *GormUserRepository) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var entity user.User
	if err := tx.First(&entity, "email = ?", email).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (g *GormUserRepository) CreateOrUpdate(ctx context.Context, user *user.User) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	return tx.Save(user).Error
}

func (g *GormUserRepository) Create(ctx context.Context, user *user.User) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	return tx.Create(user).Error
}

func (g *GormUserRepository) Update(ctx context.Context, user *user.User) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	if err := tx.Model(user).Association("Roles").Replace(user.Roles); err != nil {
		return err
	}
	return tx.Save(user).Error
}

func (g *GormUserRepository) UpdateLastLogin(ctx context.Context, id int64) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	return tx.Model(&user.User{}).Where("id = ?", id).Update("last_login", "NOW()").Error //nolint:exhaustruct
}

func (g *GormUserRepository) UpdateLastAction(ctx context.Context, id int64) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	return tx.Model(&user.User{}).Where("id = ?", id).Update("last_action", "NOW()").Error //nolint:exhaustruct
}

func (g *GormUserRepository) Delete(ctx context.Context, id int64) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	return tx.Delete(&user.User{}, id).Error //nolint:exhaustruct
}
