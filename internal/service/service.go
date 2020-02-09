package service

import (
	"context"
	"time"

	"github.com/TinyWarrior/blog-be/internal/repo"
)

type service struct {
	repo repo.Interface
}

// Interface defines the service interface
type Interface interface {
	NewPost(context.Context, string, string, string, time.Time, bool) error
}

// NewService returns an instances of service
func NewService(repo repo.Interface) Interface {
	return &service{
		repo: repo,
	}
}
