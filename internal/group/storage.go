package group

import "context"

type Repository interface {
	Create(ctx context.Context, group *Group) (string, error)
	FindAll(ctx context.Context) ([]Group, error)
	FindOne(ctx context.Context, id string) (Group, error)
	Update(ctx context.Context, group *Group) error
	Delete(ctx context.Context, id string) error
}
