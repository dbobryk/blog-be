package service

import (
	"context"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (s *service) UpdatePost(ctx context.Context, blogPost structs.BlogPost) error {

	err := s.repo.UpdatePost(ctx, blogPost)
	if err != nil {
		return err
	}

	return nil
}
