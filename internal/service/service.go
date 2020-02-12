package service

import (
	"context"

	"github.com/dbobryk/blog-be/internal/repo"
	"github.com/dbobryk/blog-be/internal/structs"
)

type service struct {
	repo repo.Interface
}

// Interface defines the service interface
type Interface interface {
	NewPost(context.Context, structs.BlogPost) error
	DeletePost(context.Context, string) error
	UpdatePost(context.Context, structs.BlogPost) error
	GetPost(context.Context, string) (structs.BlogPost, error)
}

// NewService returns an instances of service
func NewService(repo repo.Interface) Interface {
	return &service{
		repo: repo,
	}
}
