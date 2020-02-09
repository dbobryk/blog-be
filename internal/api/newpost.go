package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// NewPost creates a new blog post
func (s *BlogAPIServer) NewPost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	PostRequest := struct {
		Title     string
		Content   string
		Author    string
		Started   time.Time
		Published bool
	}{}

	json.NewDecoder(r.Body).Decode(&PostRequest)

	err := s.service.NewPost(ctx, PostRequest.Title, PostRequest.Content, PostRequest.Author, PostRequest.Started, PostRequest.Published)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Post saved successfully.\n")

}
