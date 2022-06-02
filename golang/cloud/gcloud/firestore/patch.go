//TODO: example code to demonstrate how to patch the existing record.
// i.e. given an existing record (with 10 fields) in firestore, if we want to pass only 6 fields with the updated data, this code should take those 6 fields and update the existing record, while retaining the values for the other 4 fields which are not passed.
package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/db"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/handlers"
)

var ctx context.Context

func main() {
	db.Init(ctx)

	router := mux.NewRouter()

	router.HandleFunc("/users", handlers.PostUser).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", handlers.PatchUser).Methods("PATCH")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	log.Println("Listen to server @:", port)
	if err := httpServer.ListenAndServe(); err != nil {
		panic(err)
	}

}
