package user

import (
	"bytes"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)



func UploadIdentity(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqUploadObjectIdentity
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])
	
	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + req.ObjectKey)
	upload, err := project.UploadObject(ctx, "app", key , nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var network bytes.Buffer        // Stand-in for a network connection
  	enc := gob.NewEncoder(&network) // Will write to network.
	// Encode (send) the value.
	err = enc.Encode(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Copy the data to the upload.
	d := network.Bytes()
	buf := bytes.NewBuffer(d)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
	resp := struct {
		Message string 
	} {
		Message: "Successfully uploaded object",
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UploadRecord(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqUploadObjectRecord
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])
	
	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + req.ObjectKey)
	upload, err := project.UploadObject(ctx, "app", key , nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// var network bytes.Buffer        // Stand-in for a network connection
  	// enc := gob.NewEncoder(&network) // Will write to network.
	// // Encode (send) the value.

	// // Strut to Json to String
	// b, err := json.Marshal(req.Data)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// err = enc.Encode(string(b))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }

	// // Copy the data to the upload.
	// d := network.Bytes()

	
	d, err := json.Marshal(req.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(d)
	fmt.Println(string(d) == `{"patient name":"Mohammed Fajer","record name":"Blood Test","provider":"London Hospital","date":"01/01/2024","patient address":"0x89205A3A3b2A69De6Dbf7f01ED13B2108B2c43e7","doctor address":"0x29203K3A3b2A69De6DbkO2X21D13B2108B2c43e7","doctor note":"super secret"}`)
	fmt.Println([]byte(string(d)))
	buf := bytes.NewBuffer([]byte(string(d)))

	  
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
  
	resp := struct {
		Message string 
	} {
		Message: "Successfully uploaded object",
	}

	if err := json.NewEncoder(w).Encode(&resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}





