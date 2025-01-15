package db

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"io"
	"log"
	"os"
)

func CreateClient() *firestore.Client {
	keyFile, err := os.Open(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
	if err != nil {
		log.Println(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"))
		log.Fatalf("Unable to open GOOGLE_APPLICATION_CREDENTIALS: %v", err)
	}

	defer func(keyFile *os.File) {
		err := keyFile.Close()
		if err != nil {
			log.Fatalf("Unable to close key file: %v", err)
		}
	}(keyFile)

	data, err := io.ReadAll(keyFile)

	sa := option.WithCredentialsJSON(data)
	config := &firebase.Config{ProjectID: os.Getenv("GOOGLE_PROJECT_ID")}

	app, err := firebase.NewApp(context.Background(), config, sa)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	client, err := app.Firestore(context.Background())

	if err != nil {
		log.Fatalf("error getting client: %v\n", err)
	}

	return client
}
