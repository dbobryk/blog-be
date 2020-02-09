package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/TinyWarrior/blog-be/internal/api"
	"github.com/TinyWarrior/blog-be/internal/firebaseconfig"
	"github.com/TinyWarrior/blog-be/internal/repo"
	"github.com/TinyWarrior/blog-be/internal/service"
	"github.com/gorilla/mux"
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
