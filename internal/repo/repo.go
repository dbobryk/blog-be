package repo

import (
	"context"
	"time"

	"firebase.google.com/go/db"
	"github.com/TinyWarrior/blog-be/internal/structs"
)

type repository struct {
	firebaseClient *db.Client
}

// Interface defines the interface for Repp
type Interface interface {
	NewPost(context.Context, string, string, string, string, time.Time, bool) error
	DeletePost(context.Context, string) error
	UpdatePost(context.Context, string, string, string, string, bool) error
	GetPost(context.Context, string) (structs.BlogPost, error)
}

// NewRepo returns a new instance of the Repo Interface
func NewRepo(firebaseClient *db.Client) *repository {
	return &repository{
		firebaseClient: firebaseClient,
	}
}
