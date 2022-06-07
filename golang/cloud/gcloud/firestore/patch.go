//TODO: example code to demonstrate how to patch the existing record.
// i.e. given an existing record (with 10 fields) in firestore,
//if we want to pass only 6 fields with the updated data,
//this code should take those 6 fields and update the existing record,
//while retaining the values for the other 4 fields which are not passed.
package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	id             = "123"
)

type User struct {
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Department string `json:"department"`
}

type MapData map[string]interface{}

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
		UserID:     id,
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
	updateInput := MapData{
		"name": "test_01_testing",
	}
	updatedUser := patch(w, r, ctx, id, updateInput)
	response, err := json.Marshal(updatedUser)
	if err != nil {
		http.Error(w, "error updating user", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func patch(w http.ResponseWriter, r *http.Request, ctx context.Context, id string, input MapData) (updatedUser User) {
	err := dbClient.RunTransaction(ctx, func(ctx context.Context, t *firestore.Transaction) error {
		doc, _ := t.Get(dbClient.Collection(collectionName).Doc(id))
		if doc.Ref.ID != id {
			return fmt.Errorf("ids do not match...%v", id)
		}
		log.Println("ids matched...", id)
		batch := dbClient.Batch()
		batch.Set(dbClient.Collection(collectionName).Doc(id), input, firestore.MergeAll)
		_, err := batch.Commit(ctx)
		if err != nil {
			// ERROR: com.google.cloud.datastore.emulator.impl.util.WrappedStreamObserver onError
			log.Println("batch commit error :", err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Println("transaction batch set error :", err)
		http.Error(w, "batch Set error ", http.StatusInternalServerError)
		return User{}
	}

	doc, err := dbClient.Collection(collectionName).Doc(id).Get(ctx)
	if err != nil {
		http.Error(w, "error getting user from patch", http.StatusInternalServerError)
		return User{}
	}
	err = doc.DataTo(&updatedUser)
	if err != nil {
		http.Error(w, "error mapping doc to user", http.StatusInternalServerError)
		return User{}
	}

	return updatedUser
}
