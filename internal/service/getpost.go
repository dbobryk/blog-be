package service

import (
	"context"

	"github.com/TinyWarrior/blog-be/internal/structs"
)

func (s *service) GetPost(ctx context.Context, postID string) (structs.BlogPost, error) {

	post, err := s.repo.GetPost(ctx, postID)
	if err != nil {
		return structs.BlogPost{}, err
	}
	return post, nil

}
