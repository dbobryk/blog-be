package api

import (
	"github.com/dbobryk/blog-be/internal/service"
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
