package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dbobryk/blog-be/internal/structs"
)

func (s *BlogAPIServer) UpdatePost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	var blogPost structs.BlogPost

	json.NewDecoder(r.Body).Decode(&blogPost)

	err := s.service.UpdatePost(ctx, blogPost)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	fmt.Fprintf(w, "Updated post: %s", blogPost.PostID)
}
