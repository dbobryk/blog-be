package repo

import (
	"context"
	"fmt"
)

func (r *repository) DeletePost(ctx context.Context, postID string) error {

	dbKey := fmt.Sprintf("/posts/%s", postID)

	if err := r.firebaseClient.NewRef(dbKey).Delete(ctx); err != nil {
		return fmt.Errorf("error deleting post %s: %s", postID, err.Error())
	}

	return nil

}
