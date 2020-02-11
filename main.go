package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/dbobryk/blog-be/internal/api"
	"github.com/dbobryk/blog-be/internal/firebaseconfig"
	"github.com/dbobryk/blog-be/internal/repo"
	"github.com/dbobryk/blog-be/internal/service"
	"github.com/gorilla/mux"
)

func main() {

	ctx := context.Background()

	firebaseClient, err := firebaseconfig.NewFirebaseConnection(ctx)
	if err != nil {
		log.Fatalf("Error initializing firebase: %s", err)
	}

	repo := repo.NewRepo(firebaseClient)
	service := service.NewService(repo)
	apiServer := api.NewBlogAPIServer(service)

	r := mux.NewRouter()
	r.HandleFunc("/health", health).Methods("GET")
	r.HandleFunc("/newpost", apiServer.NewPost).Methods("POST")
	r.HandleFunc("/deletepost", apiServer.DeletePost).Methods("POST")
	r.HandleFunc("/updatepost", apiServer.UpdatePost).Methods("POST")
	r.HandleFunc("/post/{postID}", apiServer.GetPost).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func health(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All is okay.\n")

}
