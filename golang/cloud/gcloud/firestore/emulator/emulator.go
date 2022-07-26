// example to demonstrate using firestore emulator instance which is running locally

package main

import (
	"context"
	"fmt"
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
	collectionName := "mycollection"

	type mystruct struct{ A string }
	rec := mystruct{A: "my data"}

	_, err = dbClient.Collection(collectionName).Doc(id).Set(ctx, rec)
	if err != nil {
		log.Fatalf("Failed to set firestore client record: %v", err)
	}

	doc, err := dbClient.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		log.Fatalf("Failed to get firestore client record: %v", err)
	}

	out := mystruct{}
	err = doc.DataTo(&out)
	if err != nil {
		log.Fatalf("failed to convert firestore record to struct: %v", err)
	}

	fmt.Printf("data read from firestore :::::::::::::::::::: %+v\n", out)
}

// // installing firestore emulator
// TODO: need to add the installation steps here

// // to tun this program"
// ./emulator.sh

// // possible errors
// 1. if you see error saying that the credentials are missing, it means the code is trying to connect to the firestore in the gcp instead of connecting to firestore emulator.
// this could happen when the emulator is not setup or is not running when the go code is run
