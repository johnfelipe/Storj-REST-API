package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)


func Download(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqDownloadObject
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

	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, "app", key, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// f, err := os.Create("test.txt")
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }
	

	// xx := bytes.NewBuffer(receivedContents)

    // l, err := f.WriteString(xx.String())
    // if err != nil {
    //     fmt.Println(err)
    //     f.Close()
    //     return
    // }
    // fmt.Println(l, "bytes written successfully")
    // err = f.Close()
    // if err != nil {
    //     fmt.Println(err)
    //     return
    // }


	// println("Here at upload")
	// println(receivedContents)
	// println(string(receivedContents))
	// println(receivedContents)

	// String to json
	var data resources.Record
	json.Unmarshal(receivedContents, &data)
	fmt.Println(data)
	

	// json.NewDecoder(bytes.NewReader(receivedContents)).Decode(&data)
	
	d, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(string(d))




	// fmt.Println(string(receivedContents))

	// reader := bytes.NewReader(receivedContents)

	

	// var data resources.Record
	// json.NewDecoder(reader).Decode(&data)

	
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}


func Downloads(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Downloads()")
}
