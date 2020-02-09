package service

import (
	"context"
	"errors"
	"time"

	"github.com/TinyWarrior/blog-be/internal/repo"
)

type service struct {
	repo repo.RepoInterface
}

type ServiceInterface interface {
	NewPost(context.Context, string, string, string, time.Time, bool) error
}

func NewService(repo repo.RepoInterface) ServiceInterface {
	return &service{
		repo: repo,
	}
}

func (s *service) NewPost(ctx context.Context, title string, content string, author string, started time.Time, published bool) error {

	if title == "" {
		return errors.New("Title must not be empty")
	}

	if content == "" {
		return errors.New("Content must not be empty")
	}

	if author == "" {
		return errors.New("Author must not be empty")
	}

	err := s.repo.NewPost(ctx, title, content, author, started, published)
	if err != nil {
		return err
	}

	return nil

}
