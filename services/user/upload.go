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


func Upload(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from Upload()")
	ctx := context.Background()
	// set the header to content type x-www-form-urlencoded
	// allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	fmt.Println(user)

	// Get User Request Data
	var objReq resources.ObjectsReq
	err := json.NewDecoder(r.Body).Decode(&objReq)
	if err != nil  {
		fmt.Println("here0")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	defer r.Body.Close()

	fmt.Println(objReq)

	

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(objReq.UserAccessGrant)
	if err != nil {
		fmt.Println("here1")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	fmt.Println(userAccess)

	
	// appAccessStr := os.Getenv("APPACCESS")
	// appAccess, err := uplink.ParseAccess(appAccessStr)
	// if err != nil {
	// 	fmt.Println("error :", err)
	// 	return
	// }

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		fmt.Println("here2")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer project.Close()


	

	// Ensure the desired Bucket within the Project is created.
	// _, err = project.EnsureBucket(ctx, "app")
	// if err != nil {
	// 	fmt.Println("error : ", err)
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return

	// }

	// // fmt.Println(p)
	// Intitiate the upload of our Object to the specified bucket and key.
	key := (user.EthereumAddress + objReq.ObjectKey[0])
	fmt.Println("key ", key)
	upload, err := project.UploadObject(ctx, "app", key , nil)
	if err != nil {
		fmt.Println("here3")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	var network bytes.Buffer        // Stand-in for a network connection
    enc := gob.NewEncoder(&network) // Will write to network.
	// Encode (send) the value.
	err = enc.Encode(objReq)
	if err != nil {
		fmt.Println("here4")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// Copy the data to the upload.
	d := network.Bytes()

	buf := bytes.NewBuffer(d)
	_, err = io.Copy(upload, buf)
	if err != nil {
		_ = upload.Abort()
		fmt.Println("here5")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// Commit the uploaded object.
	err = upload.Commit()
	if err != nil {
		fmt.Println("here6")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	resp := struct {
		Message string 
	} {
		Message: "Successfully uploaded object",
	}

	json.NewEncoder(w).Encode(&resp)

	// // w.WriteHeader(http.StatusCreated)
	// // fmt.Fprintf(w, "Upload sucess!")

}

func Uploads(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	fmt.Println(user)

	// Get User Request Data
	var objReq resources.ObjectsReq
	err := json.NewDecoder(r.Body).Decode(&objReq)
	if err != nil  {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}
	defer r.Body.Close()

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(objReq.UserAccessGrant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	fmt.Println(userAccess)

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer project.Close()

	// Intitiate the upload of our Object to the specified bucket and key.
	for i:= range objReq.ObjectKey {
		key := (user.EthereumAddress + objReq.ObjectKey[i])
		fmt.Println("key ", key)
		upload, err := project.UploadObject(ctx, "app", key , nil)
		if err != nil {
			
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var network bytes.Buffer        // Stand-in for a network connection
		enc := gob.NewEncoder(&network) // Will write to network.
		// Encode (send) the value.
		err = enc.Encode(objReq.Data[i])
		if err != nil {
		
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}

		// Copy the data to the upload.
		d := network.Bytes()

		buf := bytes.NewBuffer(d)
		_, err = io.Copy(upload, buf)
		if err != nil {
			_ = upload.Abort()

			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}

		// Commit the uploaded object.
		err = upload.Commit()
		if err != nil {
		
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return 
		}
	}
	resp := struct {
		Message string 
	} {
		Message: "Successfully uploaded object",
	}



	json.NewEncoder(w).Encode(&resp)
}