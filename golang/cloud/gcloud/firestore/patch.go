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

	"cloud.google.com/go/firestore"
	"github.com/gorilla/mux"
)

var (
	dbClient       *firestore.Client
	ctx            context.Context
	collectionName = "users"
	idConst        = "123"
)

type User struct {
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}


func Init() {
	ctx = context.TODO()
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		projectID = firestore.DetectProjectID
	}
	var err error
	dbClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}

	user := User{
		UserID:     idConst,
		Name:       "test_01",
		Email:      "test@email.com",
		Department: "testing",
	}
	_, err = dbClient.Collection(collectionName).Doc(user.UserID).Set(ctx, user)
	if err != nil {
		log.Println("error adding user")

	}
}

func main() {
	Init()
	router := mux.NewRouter()

	router.HandleFunc("/users", patchUser).Methods("PATCH")

	log.Println("Listening and Serving @8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}

func patchUser(w http.ResponseWriter, r *http.Request) {
	updateInput := map[string]interface{}{
		"name": "test_01_testing",
	}
	updatedUser, err := patch(w, r, idConst, updateInput)
	if err != nil {
		http.Error(w, "error patching user", http.StatusInternalServerError)
	}
	response, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "error marshalling updatedUser", http.StatusInternalServerError)
		return
	}
	if len(response) == 0 {
		response = []byte("no record is returned")
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func patch(w http.ResponseWriter, r *http.Request, id string, input map[string]interface{}) (User, error) {
	// batch commit method
	batch := dbClient.Batch()
	docRef := dbClient.Collection(collectionName).Doc(id)
	batch.Set(docRef, input, firestore.MergeAll)
	if _, err := batch.Commit(ctx); err != nil {
		return User{}, err
	}

	// return updated user details
	updatedUser := User{}
	doc, err := dbClient.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		return User{}, err
	}
	doc.DataTo(&updatedUser)

	return updatedUser, nil
}
