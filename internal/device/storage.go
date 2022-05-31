package device

import "context"

type Repository interface {
	Create(ctx context.Context, device *Device) (string, error)
	FindAll(ctx context.Context) ([]Device, error)
	FindOne(ctx context.Context, key string, value string) (*Device, error)
	Update(ctx context.Context, device *Device) error
	Delete(ctx context.Context, id string) error
}
