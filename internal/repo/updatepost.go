package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (r *repository) UpdatePost(ctx context.Context, blogPost structs.BlogPost) error {

	dbKey := fmt.Sprintf("/posts/%s", blogPost.PostID)

	toUpdate := map[string]interface{}{
		"Title":     blogPost.Title,
		"Author":    blogPost.Author,
		"Content":   blogPost.Content,
		"Published": blogPost.Published,
		"Updated":   time.Now(),
	}

	if err := r.firebaseClient.NewRef(dbKey).Update(ctx, toUpdate); err != nil {
		return err
	}

	return nil

}
