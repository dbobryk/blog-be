package repo

import (
	"context"
	"fmt"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (r *repository) GetPost(ctx context.Context, postID string) (structs.BlogPost, error) {

	var post structs.BlogPost

	dbKey := fmt.Sprintf("/posts/%s", postID)

	if err := r.firebaseClient.NewRef(dbKey).Get(ctx, &post); err != nil {
		return structs.BlogPost{}, err
	}

	return post, nil
}
