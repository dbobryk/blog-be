package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/TinyWarrior/blog-be/internal/api"
	"github.com/TinyWarrior/blog-be/internal/firebaseconfig"
	"github.com/TinyWarrior/blog-be/internal/repo"
	"github.com/TinyWarrior/blog-be/internal/service"
	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()

	firebaseClient, err := firebaseconfig.NewFirebaseConnection(ctx)
	if err != nil {
		log.Fatalln("Error initializing firebase: %s", err)
	}

	repo := repo.NewRepo(firebaseClient)
	service := service.NewService(repo)
	apiServer := api.NewBlogAPIServer(service)

	r := mux.NewRouter()
	r.HandleFunc("/health", health)
	r.HandleFunc("/NewPost", apiServer.NewPost)

	http.ListenAndServe(":8080", r)
}

func health(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All is okay.\n")

}
