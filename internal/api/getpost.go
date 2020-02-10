package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// GetPost gets a single post by post_id and returns it
func (s *BlogAPIServer) GetPost(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	vars := mux.Vars(r)
	postID := vars["postID"]

	post, err := s.service.GetPost(ctx, postID)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	out, err := json.Marshal(post)
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(out))

}
