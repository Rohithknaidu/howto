package service

import (
	"context"

	"github.com/muly/howto/golang/cloud/gcloud/firestore/db"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/models"
)

func AddUser(ctx context.Context, user models.User) (models.User, error) {
	return db.Add(ctx, user)
}

func GetUser(ctx context.Context, id string) (models.User, error) {
	return db.Get(ctx, id)
}

func UpdateUser(ctx context.Context, id string, mapData map[string]interface{}) (models.User, error) {
	return db.Update(ctx, id, mapData)
}
