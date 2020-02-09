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
	SavePost(context.Context, string, string, string, string, time.Time, bool) error
}

func NewRepo(firebaseClient *db.Client) *repo {
	return &repo{
		firebaseClient: firebaseClient,
	}
}
