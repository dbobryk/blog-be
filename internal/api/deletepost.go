package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// DeletePost takes a postID and deletes it
func (s *BlogAPIServer) DeletePost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	post := struct {
		PostID string
	}{}

	json.NewDecoder(r.Body).Decode(&post)
	err := s.service.DeletePost(ctx, post.PostID)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Deleted post with ID: %s/n", post.PostID)

}
