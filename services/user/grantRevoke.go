package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
	"storj.io/uplink"
)

func share(req resources.ReqShareObject, user models.DappUser, permission uplink.Permission) (string, error) {

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		return "", err 
	}

	objectPrefix := uplink.SharePrefix{Bucket: "app",  Prefix: user.EthereumAddress + req.ObjectPrefix}
	grantAccess, err := userAccess.Share(permission, objectPrefix)
	if err != nil {
		return "", err
	}

	serializedAccess, err :=grantAccess.Serialize()
	if err != nil {
		return "", err
	}

	return serializedAccess, nil 
}

func shares(req resources.ReqShareObjects, user models.DappUser, permission uplink.Permission) (string, error) {

	// Parse The access grant
	userAccess, err := uplink.ParseAccess(req.UserAccessGrant)
	if err != nil {
		return "", err 
	}

	var objectPrefixes []uplink.SharePrefix
	for i:=range req.ObjectsPrefix {
		op := uplink.SharePrefix{Bucket: "app",  Prefix: user.EthereumAddress + req.ObjectsPrefix[i]}
		objectPrefixes = append(objectPrefixes, op)

	}

	grantAccess, err := userAccess.Share(permission, objectPrefixes...)
	if err != nil {
		return "", err
	}

	serializedAccess, err :=grantAccess.Serialize()
	if err != nil {
		return "", err
	}

	return serializedAccess, nil 
}

func Share(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json; charset=utf-8")


	var req resources.ReqShareObject
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	var permission uplink.Permission
	
	switch req.PermissionType {
	case "ReadOnly": 
		permission = uplink.Permission{AllowDownload: true, NotBefore: time.Now().Add( -2 * time.Minute )}
		// accessGrant, err := share(req, user, permission)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// grant := resources.AccessGrant{Grant: accessGrant}
		// if err := json.NewEncoder(w).Encode(&grant); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	case "WriteOnly":
		permission = uplink.Permission{AllowUpload: true, NotBefore: time.Now().Add( -2 * time.Minute )}
		// accessGrant, err := share(req, user, permission)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// grant := resources.AccessGrant{Grant: accessGrant}
		// if err := json.NewEncoder(w).Encode(&grant); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	case "FullAccess":
		permission = uplink.Permission{AllowDownload: true, AllowUpload: true, AllowList: true, AllowDelete: true, NotBefore: time.Now().Add( -2 * time.Minute ) }
		// accessGrant, err := share(req, user, permission)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// grant := resources.AccessGrant{Grant: accessGrant}
		// if err := json.NewEncoder(w).Encode(&grant); err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }
	default:
		http.Error(w, "Request body invalid permission", http.StatusBadRequest)
		return
	}

	accessGrant, err := share(req, user, permission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	grant := resources.AccessGrant{Grant: accessGrant}
	if err := json.NewEncoder(w).Encode(&grant); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	


}

func Shares(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")


	var req resources.ReqShareObjects
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Get User from database
	params := mux.Vars(r)
	var user models.DappUser 
	database.Db.First(&user, params["id"])

	var permission uplink.Permission

	switch req.PermissionType {
		case "ReadOnly": 
				permission = uplink.Permission{AllowDownload: true, NotBefore: time.Now().Add( -2 * time.Minute )}
				// accessGrant, err := shares(req, user, permission)
				// if err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// 	return
				// }
				// grant := resources.AccessGrant{Grant: accessGrant}
				// if err := json.NewEncoder(w).Encode(&grant); err != nil {
				// 	http.Error(w, err.Error(), http.StatusInternalServerError)
				// }
		case "WriteOnly":
			permission = uplink.Permission{AllowUpload: true, NotBefore: time.Now().Add( -2 * time.Minute )}
			// accessGrant, err := shares(req, user, permission)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }
			// grant := resources.AccessGrant{Grant: accessGrant}
			// if err := json.NewEncoder(w).Encode(&grant); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
			
		case "FullAccess":
			permission = uplink.Permission{AllowDownload: true, AllowUpload: true, AllowList: true, AllowDelete: true, NotBefore: time.Now().Add( -2 * time.Minute ) }
			// accessGrant, err := shares(req, user, permission)
			// if err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// 	return
			// }
			// grant := resources.AccessGrant{Grant: accessGrant}
			// if err := json.NewEncoder(w).Encode(&grant); err != nil {
			// 	http.Error(w, err.Error(), http.StatusInternalServerError)
			// }
		default:
			http.Error(w, "Request body invalid permission", http.StatusBadRequest)
			return
		}

		accessGrant, err := shares(req, user, permission)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		grant := resources.AccessGrant{Grant: accessGrant}
		if err := json.NewEncoder(w).Encode(&grant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
}

func RevokeGrant(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Get User Request Data
	var req resources.ReqRevokeAccess
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
		return 
	}

	// Open up the project we will be working with
	project, err := uplink.OpenProject(ctx, userAccess)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}
	defer project.Close()

	// Get access grant 
	access, err := uplink.ParseAccess(req.AccessGrant)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	// Revoke Access
	err = project.RevokeAccess(ctx, access)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return 
	}

	res := resources.SuccessMessage{Message: "Successfully revoked acccess grant!"}
	if err := json.NewEncoder(w).Encode(&res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}


}


// func RevokeGrants(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "from RevokeGrants()")
// }