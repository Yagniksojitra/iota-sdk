package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"github.com/iota-agency/iota-erp/internal/domain/entities/user"
	model "github.com/iota-agency/iota-erp/internal/interfaces/graph/gqlmodels"
	"github.com/iota-agency/iota-erp/sdk/mapper"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (*model.User, error) {
	u := &user.User{}
	if err := mapper.LenientMapping(&input, u); err != nil {
		return nil, err
	}
	if input.Password != nil {
		if err := u.SetPassword(*input.Password); err != nil {
			return nil, err
		}
	}
	if err := r.app.UserService.Create(ctx, u); err != nil {
		return nil, err
	}
	return u.ToGraph(), nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int64, input model.UpdateUser) (*model.User, error) {
	entity, err := r.app.UserService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if err := mapper.LenientMapping(&input, entity); err != nil {
		return nil, err
	}
	if input.Password != nil {
		if err := entity.SetPassword(*input.Password); err != nil {
			return nil, err
		}
	}
	if err := r.app.UserService.Update(ctx, entity); err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int64) (*model.User, error) {
	entity, err := r.app.UserService.Delete(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int64) (*model.User, error) {
	entity, err := r.app.UserService.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedUsers, error) {
	entities, err := r.app.UserService.GetPaginated(ctx, limit, offset, sortBy)
	if err != nil {
		return nil, err
	}
	result := make([]*model.User, len(entities))
	for i, entity := range entities {
		result[i] = entity.ToGraph()
	}
	total, err := r.app.UserService.Count(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PaginatedUsers{
		Data:  result,
		Total: total,
	}, nil
}

// UserCreated is the resolver for the userCreated field.
func (r *subscriptionResolver) UserCreated(ctx context.Context) (<-chan *model.User, error) {
	ch := make(chan *model.User)
	r.app.EventPublisher.Subscribe(func(evt *user.Created) {
		ch <- evt.Result.ToGraph()
	})
	return ch, nil
}

// UserUpdated is the resolver for the userUpdated field.
func (r *subscriptionResolver) UserUpdated(ctx context.Context) (<-chan *model.User, error) {
	ch := make(chan *model.User)
	r.app.EventPublisher.Subscribe(func(evt *user.Updated) {
		ch <- evt.Result.ToGraph()
	})
	return ch, nil
}

// UserDeleted is the resolver for the userDeleted field.
func (r *subscriptionResolver) UserDeleted(ctx context.Context) (<-chan *model.User, error) {
	ch := make(chan *model.User)
	r.app.EventPublisher.Subscribe(func(evt *user.Deleted) {
		ch <- evt.Result.ToGraph()
	})
	return ch, nil
}

// Avatar is the resolver for the avatar field.
func (r *userResolver) Avatar(ctx context.Context, obj *model.User) (*model.Media, error) {
	if obj.AvatarID == nil {
		return nil, nil
	}
	upload, err := r.app.UploadService.GetUploadByID(ctx, *obj.AvatarID)
	if err != nil {
		return nil, err
	}
	return upload.ToGraph(), nil
}

// User returns UserResolver implementation.
func (r *Resolver) User() UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
