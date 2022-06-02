package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/models"
	"github.com/muly/howto/golang/cloud/gcloud/firestore/service"
)

var ctx context.Context

func PostUser(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	defer r.Body.Close()
	reqUser := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		http.Error(w, "error decoding reqest body", http.StatusBadRequest)
		return
	}

	resUser, err := service.AddUser(ctx, reqUser)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx = context.Background()
	params := mux.Vars(r)
	id := params["id"]

	resUser, err := service.GetUser(ctx, id)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}

func PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id := params["id"]
	defer r.Body.Close()

	reqUser := map[string]interface{}{}
	if err := json.NewDecoder(r.Body).Decode(&reqUser); err != nil {
		http.Error(w, "error decoding reqest body", http.StatusBadRequest)
		return
	}

	resUser, err := service.UpdateUser(ctx, id, reqUser)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, "error adding user", http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(response)
}
