package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"
	"fmt"

	model "github.com/iota-agency/iota-erp/graph/gqlmodels"
)

// DeleteUpload is the resolver for the deleteUpload field.
func (r *mutationResolver) DeleteUpload(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUpload - deleteUpload"))
}

// Upload is the resolver for the upload field.
func (r *queryResolver) Upload(ctx context.Context, id int64) (*model.Upload, error) {
	entity, err := r.app.UploadService.GetUploadByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return entity.ToGraph(), nil
}

// Uploads is the resolver for the uploads field.
func (r *queryResolver) Uploads(ctx context.Context, offset int, limit int, sortBy []string) (*model.PaginatedUploads, error) {
	uploads, err := r.app.UploadService.GetUploadsPaginated(ctx, limit, offset, sortBy)
	if err != nil {
		return nil, err
	}
	result := make([]*model.Upload, len(uploads))
	for _, upload := range uploads {
		result = append(result, upload.ToGraph())
	}
	total, err := r.app.UploadService.GetUploadsCount(ctx)
	if err != nil {
		return nil, err
	}
	return &model.PaginatedUploads{
		Data:  result,
		Total: total,
	}, nil
}

// UploadCreated is the resolver for the uploadCreated field.
func (r *subscriptionResolver) UploadCreated(ctx context.Context) (<-chan *model.Upload, error) {
	panic(fmt.Errorf("not implemented: UploadCreated - uploadCreated"))
}

// UploadUpdated is the resolver for the uploadUpdated field.
func (r *subscriptionResolver) UploadUpdated(ctx context.Context) (<-chan *model.Upload, error) {
	panic(fmt.Errorf("not implemented: UploadUpdated - uploadUpdated"))
}

// UploadDeleted is the resolver for the uploadDeleted field.
func (r *subscriptionResolver) UploadDeleted(ctx context.Context) (<-chan int64, error) {
	panic(fmt.Errorf("not implemented: UploadDeleted - uploadDeleted"))
}
