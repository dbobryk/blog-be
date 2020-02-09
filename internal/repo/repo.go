package repo

import (
	"context"
	"time"

	"firebase.google.com/go/db"
)

type repo struct {
	firebaseClient *db.Client
}

type RepoInterface interface {
	NewPost(context.Context, string, string, string, time.Time, bool) error
}

func NewRepo(firebaseClient *db.Client) *repo {
	return &repo{
		firebaseClient: firebaseClient,
	}
}

func (r *repo) NewPost(ctx context.Context, title string, content string, author string, started time.Time, published bool) error {

	return nil

}
