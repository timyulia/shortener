//go:generate go run github.com/golang/mock/mockgen -destination=./mock/service_gen.go -source=service.go -package=mock service
package service

import "context"

type Repo interface {
	AddLinksPair(ctx context.Context, short, long string) error
	GetLongURL(ctx context.Context, short string) (string, error)
}

type Service struct {
	repo Repo
}

func New(r Repo) *Service {
	return &Service{
		repo: r,
	}
}
