package firebaseconfig

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
)

// NewFirebaseConnection returns a new firebase connection
func NewFirebaseConnection(ctx context.Context) (*db.Client, error) {

	databaseURL := os.Getenv("DB_NAME")

	config := &firebase.Config{
		DatabaseURL: databaseURL,
	}

	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		return nil, err
	}

	client, err := app.Database(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}
