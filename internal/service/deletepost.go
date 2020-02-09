package service

import (
	"context"
	"errors"
)

func (s *service) DeletePost(ctx context.Context, postID string) error {

	if postID == "" {
		return errors.New("postID must have a value")
	}

	err := s.repo.DeletePost(ctx, postID)
	if err != nil {
		return err
	}

	return nil
}
