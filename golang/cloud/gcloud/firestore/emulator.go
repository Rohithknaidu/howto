// TODO: example to demonstrate using firestore emulator instance which is running locally

package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
)

func main() {

	ctx := context.Background()
	dbClient, err := firestore.NewClient(ctx, "*detect-project-id*")
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	id := "id123"
	rec := struct{ A string }{A: "my data"}
	// rec := "string data"

	// ctx = context.Background()
	_, err = dbClient.Collection("checkLog").Doc(id).Set(ctx, rec)
	if err != nil {
		log.Fatalf("Failed to set firestore client record: %v", err)
	}

	// doc := dbClient.Collection("whitelistVideos").Doc(id)
	// data, err := doc.Get(ctx)
	// if err != nil {
	// 	log.Fatalf("Failed to get firestore client record: %v", err)
	// }
	// fmt.Println(data)

}

// FIRESTORE_EMULATOR_HOST=::1:8948 go run local.go
