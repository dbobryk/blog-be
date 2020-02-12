package service

import (
	"context"
	"errors"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (s *service) NewPost(ctx context.Context, blogPost structs.BlogPost) error {

	if blogPost.Title == "" {
		return errors.New("title must not be empty")
	}

	if blogPost.Content == "" {
		return errors.New("content must not be empty")
	}

	if blogPost.Author == "" {
		return errors.New("author must not be empty")
	}

	blogPost.PostID = NewShortUUID()

	err := s.repo.NewPost(ctx, blogPost)
	if err != nil {
		return err
	}

	return nil

}
