package repo

import (
	"context"
	"fmt"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (r *repository) NewPost(ctx context.Context, blogPost structs.BlogPost) error {

	dbKey := fmt.Sprintf("/posts/p-%s", blogPost.PostID)

	if err := r.firebaseClient.NewRef(dbKey).Set(ctx, blogPost); err != nil {
		return err
	}

	return nil

}
