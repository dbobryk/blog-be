package repo

import (
	"context"

	"firebase.google.com/go/db"
	"github.com/dbobryk/blog-be/internal/structs"
)

type repository struct {
	firebaseClient *db.Client
}

// Interface defines the interface for Repp
type Interface interface {
	NewPost(context.Context, structs.BlogPost) error
	DeletePost(context.Context, string) error
	UpdatePost(context.Context, structs.BlogPost) error
	GetPost(context.Context, string) (structs.BlogPost, error)
}

// NewRepo returns a new instance of the Repo Interface
func NewRepo(firebaseClient *db.Client) *repository {
	return &repository{
		firebaseClient: firebaseClient,
	}
}
