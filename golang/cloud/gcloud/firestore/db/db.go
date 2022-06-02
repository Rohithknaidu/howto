package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/models"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/utils"
)

var dbClient *firestore.Client

func Init(ctx context.Context) {
	projectID := os.Getenv("GCP_PROJECT")
	if projectID == "" {
		//Note: in case of app deployed to app engine on gcp, this line will detect the current project id, so this env need not be set in api.yaml
		projectID = firestore.DetectProjectID
	}
	var err error
	// Get a Firestore client.
	dbClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
}

func Get(ctx context.Context, id string) (models.User, error) {
	doc, err := dbClient.Collection("users").Doc(id).Get(ctx)
	if err != nil {
		return models.User{}, err
	}
	user := models.User{}
	doc.DataTo(&user)

	return user, nil
}

func Add(ctx context.Context, user models.User) (models.User, error) {
	user.UserID = utils.NewID()

	err := dbClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		_, err := dbClient.Collection("users").Doc(user.UserID).Set(ctx, user)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

func Update(ctx context.Context, id string, mapData map[string]interface{}) (models.User, error) {
	batch := dbClient.Batch()
	batch.Set(dbClient.Collection("users").Doc(id), mapData, firestore.MergeAll)
	_, err := batch.Commit(ctx)
	if err != nil {
		return models.User{}, fmt.Errorf("batc Set error :%v", err)
	}

	user, err := Get(ctx, id)
	if err != nil {
		return models.User{}, fmt.Errorf("user return error :%v", err)
	}

	return user, nil
}
