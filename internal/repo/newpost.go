package repo

import (
	"context"
	"fmt"
	"time"
)

func (r *repository) NewPost(ctx context.Context, title string, content string, author string, started time.Time, published bool) error {

	post := struct {
		Title     string
		Content   string
		Author    string
		Started   time.Time
		Published bool
	}{
		Title:     title,
		Content:   content,
		Author:    author,
		Started:   started,
		Published: published,
	}

	dbKey := fmt.Sprintf("/posts/%s", "testing")

	if err := r.firebaseClient.NewRef(dbKey).Set(ctx, post); err != nil {
		return err
	}

	return nil

}
