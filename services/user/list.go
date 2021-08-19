package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)

func List(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from List()")



	// set the header to content type x-www-form-urlencoded
	// allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	// Get User Request Data
	var objReq resources.ObjectsReq
	err := json.NewDecoder(r.Body).Decode(&objReq)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()

	ctx := context.Background()

	userAccess, err := uplink.ParseAccess(objReq.UserAccessGrant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		log.Println("Can't open project with user Access")
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer project.Close()
	objects := project.ListObjects(ctx, "app", &uplink.ListObjectsOptions{Prefix: user.EthereumAddress + objReq.UserPrefix})
	var listObjects resources.ListData


	for objects.Next() {
		item := objects.Item()
		listObjects.Objects = append(listObjects.Objects, item.Key)
		fmt.Println(item.IsPrefix, item.Key)
	}
	if err := objects.Err(); err != nil {
		log.Println("Can't access objects")
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


	json.NewEncoder(w).Encode(listObjects)
	// fmt.Fprintln(w, listObjects)
}