package user

import "context"

type Service interface {
	Create(ctx context.Context) error
	Get(ctx context.Context) error
	GetId(ctx context.Context) error
	Delete(ctx context.Context) error
	Patch(ctx context.Context) error
}

type service struct {
}

func NewService() *service {
	return &service{}
}
