package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)


func downloadIdentity(ctx context.Context, req resources.ReqDownloadObject, user models.DappUser) (resources.Identity, error) {

	var data resources.Identity

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		return resources.Identity{}, err 
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		return resources.Identity{}, err 
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + req.ObjectKey)
	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, "app", key, nil)
	if err != nil {
		return resources.Identity{}, err 
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return resources.Identity{}, err 
	}

	json.Unmarshal(receivedContents, &data)
	
	return data, nil
}


func downloadRecord(ctx context.Context, req resources.ReqDownloadObject, user models.DappUser) (resources.Record, error) {

	var data resources.Record

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		return resources.Record{}, err 
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		return resources.Record{}, err 
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + req.ObjectKey)
	// Initiate a download of the same object again
	download, err := project.DownloadObject(ctx, "app", key, nil)
	if err != nil {
		return resources.Record{}, err 
	}
	defer download.Close()

	// Read everything from the download stream
	receivedContents, err := ioutil.ReadAll(download)
	if err != nil {
		return resources.Record{}, err 
	}

	json.Unmarshal(receivedContents, &data)
	
	return data, nil
}


func oldDownload(w http.ResponseWriter, r *http.Request) {
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


	// String to json
	result := strings.Contains(req.ObjectKey, "identity")
	if result {
		var data resources.Identity
		json.Unmarshal(receivedContents, &data)
		fmt.Println(data)
	
		d, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(d))
		if err := json.NewEncoder(w).Encode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		var data resources.Record
		json.Unmarshal(receivedContents, &data)
		fmt.Println(data)
		

		
		d, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(d))
		if err := json.NewEncoder(w).Encode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}


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

	// String to json
	result := strings.Contains(req.ObjectKey, "identity")
	if result {
		data, err := downloadIdentity(ctx, req, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		data, err := downloadRecord(ctx, req, user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(&data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}




func oldDownloads(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	var req resources.ReqDownloadObjects
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

	var identites []resources.Identity
	var records []resources.Record
	for i:=range req.ObjectKeys {
		// Intitiate the upload of our Object to the specified bucket and key.
		key := (user.EthereumAddress + req.ObjectKeys[i])

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


	

		// String to json
		result := strings.Contains(req.ObjectKeys[i], "identity")
		if result {
			var data resources.Identity
			json.Unmarshal(receivedContents, &data)
			fmt.Println(data)
		
			d, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println(string(d))

			identites = append(identites, data)

			// if err := json.NewEncoder(w).Encode(&data); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
		} else {
			var data resources.Record
			json.Unmarshal(receivedContents, &data)
			fmt.Println(data)
			

			
			d, err := json.Marshal(data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			fmt.Println(string(d))

			records = append(records, data)

			// if err := json.NewEncoder(w).Encode(&data); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
		}
	}


	if err := json.NewEncoder(w).Encode(&identites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// for i:= range identites {
	// 	if err := json.NewEncoder(w).Encode(&identites[i]); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// }

	// for i:= range records {
	// 	if err := json.NewEncoder(w).Encode(&records[i]); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// }
}


func Downloads(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var req resources.ReqDownloadObjects
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	var identites []resources.Identity
	var records []resources.Record

	for i:=range req.ObjectKeys {
		// String to json
		result := strings.Contains(req.ObjectKeys[i], "identity")
		if result {
			data, err := downloadIdentity(ctx, resources.ReqDownloadObject{UserAccessGrant: req.UserAccessGrant, ObjectKey: req.ObjectKeys[i]}, user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			identites = append(identites, data)
		} else {
			data, err := downloadRecord(ctx, resources.ReqDownloadObject{UserAccessGrant: req.UserAccessGrant, ObjectKey: req.ObjectKeys[i]}, user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			records = append(records, data)
		}
	}

	if err := json.NewEncoder(w).Encode(&identites); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := json.NewEncoder(w).Encode(&records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}	
}