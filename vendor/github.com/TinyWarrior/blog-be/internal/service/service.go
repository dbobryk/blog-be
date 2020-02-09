package service

import (
	"github.com/TinyWarrior/blog-be/internal/repo"
)


type service struct {
	repo repo.RepoInterface
}

type ServiceInterface interface {
	NewPost(context.Context, string, string, string, string, bool) error
}

func NewService(repo repo.Interface) {
	return service{
		repo: repo.Interface
	}
}


func (s *service) NewPost (ctx context.Context, title string, content string, author string, started time.Time, published bool) error {

	if title == "" {
		return errors.New("Title must not be empty")
	}

	if content == "" {
		return errors.New("Content must not be empty")
	}

	if author == "" {
		return errors.New("Author must not be empty")
	}

	if started == time.Time {
		return errors.New("Started time must not be empty")
	}

	if published == nil {
		return errors.New("Published must have a true/false value")
	}

	s.repo.SavePost()

}
