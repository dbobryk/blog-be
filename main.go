package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tinywarrior/blog-be/internal/api"
	"github.com/tinywarrior/blog-be/internal/firebaseconfig"
	"github.com/tinywarrior/blog-be/internal/repo"
	"github.com/tinywarrior/blog-be/internal/service"
)

func main() {

	ctx := context.Background()

	firebaseClient := firebaseconfig.NewFirebaseClient(ctx)

	repo := repo.NewRepo(firebaseClient)
	service := service.NewService(repo)
	apiServer := api.NewBlogApiServer(service)

	r := mux.NewRouter()
	r.HandleFunc("/health", health)
	r.HandleFunc("/NewPost", apiServer.NewPost)

	http.ListenAndServe(":8080", r)
}

func health(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All is okay.\n")

}
