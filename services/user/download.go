package user

import (
	"fmt"
	"net/http"
)


func Download(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "from Download()")

	// v:= r.URL.Query()
	
	// var data interface {}
	
	// switch v.Get("data") {
	// case "identity":
	// 	 data = resources.Identity{}
	// case "record":
	// 	 data = resources.Record{}	
	// }


	
	// ctx := context.Background()

	// w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "GET")
	// w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// params := mux.Vars(r)
	// var user models.DappUser 
	// database.Db.First(&user, params["id"])
	
	// // Get User Request Data
	// var objReq resources.ObjectsReq
	// err := json.NewDecoder(r.Body).Decode(&objReq)
	// if err != nil  {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return 
	// }
	// defer r.Body.Close()

	

	// // Parse The access grant
	// userAccess, err := uplink.ParseAccess(objReq.UserAccessGrant)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 
	// }
	// // fmt.Println(userAccess)

	// // Open up the project we will be working with
	// project, err := uplink.OpenProject(ctx, userAccess)
	// if err != nil {
		
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 
	// }
	// defer project.Close()

	// key := (user.EthereumAddress + objReq.ObjectKey[0])
	// // Initiate a download of the same object again
	// download, err := project.DownloadObject(ctx, "app", key, nil)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 
	// }
	// defer download.Close()

	// // Read everything from the download stream
	// receivedContents, err := ioutil.ReadAll(download)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 
	// }

	
	// network := bytes.NewBuffer(receivedContents)
	// // Decode 
	// var d resources.ObjectsReq
	// dec := gob.NewDecoder(network)
	// err = dec.Decode(&d)
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return 
    // }

	// // if err := json.Unmarshal(receivedContents, &data); err != nil {
    // //     http.Error(w, err.Error(), http.StatusInternalServerError)
	// // 	return 
    // // }

	// // b, err := base64.StdEncoding.DecodeString(string(receivedContents))
	// // if err != nil {
	// // 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// // 	return 
	// // }
	// fmt.Println(d)


	// json.NewEncoder(w).Encode(d)
}


func Downloads(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from Downloads()")
}
