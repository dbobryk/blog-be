package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dbobryk/blog-be/internal/structs"
)

// NewPost creates a new blog post
func (s *BlogAPIServer) NewPost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	var blogPost structs.BlogPost

	json.NewDecoder(r.Body).Decode(&blogPost)

	err := s.service.NewPost(ctx, blogPost)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Post saved successfully.\n")

}
