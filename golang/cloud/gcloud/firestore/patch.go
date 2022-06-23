//TODO: example code to demonstrate how to patch the existing record.
// i.e. given an existing record (with 10 fields) in firestore,
//if we want to pass only 6 fields with the updated data,
//this code should take those 6 fields and update the existing record,
//while retaining the values for the other 4 fields which are not passed.
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/exec"

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

var (
	dbClient       *firestore.Client
	ctx            context.Context
	collectionName = "users"
	id             = "1234"
)

type User struct {
	UserID     string `json:"id" firestore:"id"`
	Name       string `json:"name" firestore:"name"`
	Email      string `json:"email" firestore:"email"`
	Department string `json:"department" firestore:"department"`
}

func Init() {
	ctx = context.Background()
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		projectID = firestore.DetectProjectID
	}
	var err error
	dbClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	// post one user for patch
	user := User{
		UserID:     id,
		Name:       "test_01",
		Email:      "test@email.com",
		Department: "testing",
	}
	_, err = dbClient.Collection(collectionName).Doc(user.UserID).Set(ctx, user)
	if err != nil {
		log.Panic("postUser: post error:", err)
	}
	log.Println("Existing user details:", user)
}

func main() {
	Init()
	router := mux.NewRouter()

	router.HandleFunc("/users", patchUser).Methods("PATCH")
	
	// go routine to send patch request
	go func() {
		cmd := exec.Command("curl", "--location", "--request", "PATCH", "http://localhost:8080/users")
		if err := cmd.Run(); err != nil {
			log.Println("go(): exec command run error:", err)
		}

	}()

	log.Fatal(http.ListenAndServe(":8080", router))
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	updateInput := map[string]interface{}{
		"name": "test_01_testing",
	}

	// batch commit method
	batch := dbClient.Batch()
	batch.Set(dbClient.Collection(collectionName).Doc(id), updateInput, firestore.MergeAll)
	if _, err := batch.Commit(ctx); err != nil {
		http.Error(w, "patchUser: batch commit error", http.StatusInternalServerError)
		return
	}

	// get updated user details
	doc, err := dbClient.Collection(collectionName).Doc(id).Get(r.Context())
	if err != nil {
		http.Error(w, "patchUser: get id error", http.StatusInternalServerError)
		return
	}
	updatedUser := User{}
	doc.DataTo(&updatedUser)
	response, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "updatedUser marshalling error", http.StatusInternalServerError)
		return
	}
	if len(response) == 0 {
		response = []byte("no record is returned")
	}
	log.Println("Updated user details:", updatedUser)

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
