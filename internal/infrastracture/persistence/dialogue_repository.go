package persistence

import (
	"context"
	dialogue2 "github.com/iota-agency/iota-erp/internal/domain/entities/dialogue"
	"github.com/iota-agency/iota-erp/sdk/composables"
	"github.com/iota-agency/iota-erp/sdk/service"
)

type GormDialogueRepository struct {
}

func NewDialogueRepository() dialogue2.Repository {
	return &GormDialogueRepository{}
}

func (g *GormDialogueRepository) GetPaginated(ctx context.Context, limit, offset int, sortBy []string) ([]*dialogue2.Dialogue, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var uploads []*dialogue2.Dialogue
	q := tx.Limit(limit).Offset(offset)
	for _, s := range sortBy {
		q = q.Order(s)
	}
	if err := q.Find(&uploads).Error; err != nil {
		return nil, err
	}
	return uploads, nil
}

func (g *GormDialogueRepository) Count(ctx context.Context) (int64, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return 0, service.ErrNoTx
	}
	var count int64
	if err := tx.Model(&dialogue2.Dialogue{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (g *GormDialogueRepository) GetAll(ctx context.Context) ([]*dialogue2.Dialogue, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var entities []*dialogue2.Dialogue
	if err := tx.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (g *GormDialogueRepository) GetByID(ctx context.Context, id int64) (*dialogue2.Dialogue, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var entity dialogue2.Dialogue
	if err := tx.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

func (g *GormDialogueRepository) GetByUserID(ctx context.Context, userID int64) ([]*dialogue2.Dialogue, error) {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return nil, service.ErrNoTx
	}
	var entities []*dialogue2.Dialogue
	if err := tx.Where("user_id = ?", userID).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (g *GormDialogueRepository) Create(ctx context.Context, data *dialogue2.Dialogue) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	if err := tx.Create(data).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormDialogueRepository) Update(ctx context.Context, data *dialogue2.Dialogue) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	if err := tx.Save(data).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormDialogueRepository) Delete(ctx context.Context, id int64) error {
	tx, ok := composables.UseTx(ctx)
	if !ok {
		return service.ErrNoTx
	}
	if err := tx.Delete(&dialogue2.Dialogue{}, id).Error; err != nil {
		return err
	}
	return nil
}
