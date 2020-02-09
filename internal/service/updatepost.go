package service

import "context"

func (s *service) UpdatePost(ctx context.Context, postID string, title string, author string, content string, published bool) error {

	err := s.repo.UpdatePost(ctx, postID, title, author, content, published)
	if err != nil {
		return err
	}

	return nil
}
