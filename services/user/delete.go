package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)

func delete(ctx context.Context, req resources.ReqDeleteObject, user models.DappUser) (error) {

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		return err 
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		return err 
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + req.ObjectKey)

	_, err = project.DeleteObject(ctx, "app", key)
	if err != nil {
		return err 
	}
	
	return nil
}

func Delete(w http.ResponseWriter, r *http.Request) {
	

	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqDeleteObject
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	// Delete object
	err := delete(ctx, req, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send Response back to client
	status := resources.SuccessMessage{Message: "Successfully deleted object"}
	if err := json.NewEncoder(w).Encode(&status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
}

func Deletes(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqDeleteObjects
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()


	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	for i:= range req.ObjectKeys {
		// Delete object
		err := delete(ctx, resources.ReqDeleteObject{UserAccessGrant: req.UserAccessGrant, ObjectKey: req.ObjectKeys[i]}, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Send Response back to client
	status := resources.SuccessMessage{Message: "Successfully deleted objects"}
	if err := json.NewEncoder(w).Encode(&status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
