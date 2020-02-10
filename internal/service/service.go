package service

import (
	"context"
	"time"

	"github.com/TinyWarrior/blog-be/internal/repo"
	"github.com/TinyWarrior/blog-be/internal/structs"
)

type service struct {
	repo repo.Interface
}

// Interface defines the service interface
type Interface interface {
	NewPost(context.Context, string, string, string, time.Time, bool) error
	DeletePost(context.Context, string) error
	UpdatePost(context.Context, string, string, string, string, bool) error
	GetPost(context.Context, string) (structs.BlogPost, error)
}

// NewService returns an instances of service
func NewService(repo repo.Interface) Interface {
	return &service{
		repo: repo,
	}
}
