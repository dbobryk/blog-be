package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tinywarrior/blog-be/internal/service"
)

// BlogAPIServer defines the structure for the api server
type BlogAPIServer struct {
	service service.Interface
}

// NewBlogAPIServer returns an instance of the BlogApiServer
func NewBlogAPIServer(service service.Interface) *BlogAPIServer {

	return &BlogAPIServer{
		service: service,
	}
}

// NewPost creates a new blog post
func (s *BlogAPIServer) NewPost(w http.ResponseWriter, req *http.Request) {

	ctx := context.Background()

	PostRequest := struct {
		Title     string
		Content   string
		Author    string
		Started   time.Time
		Published bool
	}{}

	json.NewDecoder(req.Body).Decode(&PostRequest)

	err := service.NewPost(ctx, PostRequest.Title, PostRequest.Content, PostRequest.Author, PostRequest.Started, PostRequest.Published)
	if err != nil {
		fmt.Fprintf(w, err)
	}

	fmt.Fprintf(w, "Post saved successfully.")

}
