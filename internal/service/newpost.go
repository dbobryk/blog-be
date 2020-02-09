package service

import (
	"context"
	"errors"
	"time"
)

func (s *service) NewPost(ctx context.Context, title string, content string, author string, started time.Time, published bool) error {

	if title == "" {
		return errors.New("title must not be empty")
	}

	if content == "" {
		return errors.New("content must not be empty")
	}

	if author == "" {
		return errors.New("author must not be empty")
	}

	postID := NewShortUUID()

	err := s.repo.NewPost(ctx, postID, title, content, author, started, published)
	if err != nil {
		return err
	}

	return nil

}
