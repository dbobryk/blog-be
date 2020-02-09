package repo

import (
	"context"
	"fmt"
	"time"
)

func (r *repository) UpdatePost(ctx context.Context, postID string, title string, author string, content string, published bool) error {

	dbKey := fmt.Sprintf("/posts/%s", postID)

	toUpdate := map[string]interface{}{
		"Title":     title,
		"Author":    author,
		"Content":   content,
		"Published": published,
		"Updated":   time.Now(),
	}

	if err := r.firebaseClient.NewRef(dbKey).Update(ctx, toUpdate); err != nil {
		return err
	}

	return nil

}
