package repo

import (
	"context"
	"time"

	"firebase.google.com/go/db"
)

type repository struct {
	firebaseClient *db.Client
}

// Interface defines the interface for Repp
type Interface interface {
	NewPost(context.Context, string, string, string, string, time.Time, bool) error
}

// NewRepo returns a new instance of the Repo Interface
func NewRepo(firebaseClient *db.Client) *repository {
	return &repository{
		firebaseClient: firebaseClient,
	}
}
