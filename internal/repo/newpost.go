package repo

import (
	"context"
	"fmt"
	"time"
)

func (r *repository) NewPost(ctx context.Context, postID string, title string, content string, author string, started time.Time, published bool) error {

	post := struct {
		PostID    string
		Title     string
		Content   string
		Author    string
		Started   time.Time
		Published bool
	}{
		PostID:    postID,
		Title:     title,
		Content:   content,
		Author:    author,
		Started:   started,
		Published: published,
	}

	dbKey := fmt.Sprintf("/posts/p-%s", post.PostID)

	if err := r.firebaseClient.NewRef(dbKey).Set(ctx, post); err != nil {
		return err
	}

	return nil

}
