package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
)

func NewDBClient() *firestore.Client {
	projectID := "tasks-api-dev" // add gcp project id here

	if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/Users/muly/Downloads/tasks-api-dev-398f3d05127b.json"); err != nil { // set the path to the credentials json
		panic(err)
	}

	dbClient, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
	return dbClient
}

func main() {
	dbConn := NewDBClient()
	ctx := context.Background()
	collection := "tasksbytags"

	// // scenario 1: filter on document id
	doc, err := dbConn.Collection(collection).Doc("tag1").Get(ctx)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(doc.Data(), doc.Ref.ID)
	fmt.Println()

	// // scenario 2: filter on document id using a list
	// in clause is not possible on the doc id. workaround is to add the doc id as a field and apply the where clause "in" filter on that field.

	// // scenario 3: filter on the array field
	// the taskDocIds field is an array, and we are checking if this array contains any of the given strings
	docs1, err := dbConn.Collection(collection).Query.Where("taskDocIds", "array-contains-any", []string{"ask1", "task-A"}).Documents(ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, doc := range docs1 {
		fmt.Println(doc.Data(), doc.Ref.ID)
	}
	fmt.Println()

	// // scenario 4: filter on a string field to search for any match in a list of values
	// below query checks for records which has any of those values in the tag field.
	docs2, err := dbConn.Collection(collection).Query.Where("tag", "in", []string{"tag2", "tag3"}).Documents(ctx).GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, doc := range docs2 {
		fmt.Println(doc.Data(), doc.Ref.ID)
	}
	fmt.Println()
}
