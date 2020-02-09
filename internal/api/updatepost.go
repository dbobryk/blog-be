package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func (s *BlogAPIServer) UpdatePost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	post := struct {
		PostID    string
		Title     string
		Author    string
		Content   string
		Published bool
	}{}

	json.NewDecoder(r.Body).Decode(&post)

	err := s.service.UpdatePost(ctx, post.PostID, post.Title, post.Author, post.Content, post.Published)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Updated post: %s", post.PostID)
}
