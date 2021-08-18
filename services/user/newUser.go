package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/models"
	"github.com/mohammedfajer/Storj-REST-API/resources"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	fmt.Println("from CreateUser()")
	
	// set the header to content type x-www-form-urlencoded
	// allow all origin to handle cors issue
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	
	
	// create an empty user of type models.DappUser
	var user models.DappUser
	
	// decode the json request to user
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(user)
	
	// Check user in database
	// insert user in the database
	createdUser := database.Db.Create(&user)
	err = createdUser.Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} 
	
	// format a resposnse object
	res := resources.Response{
			ID:      user.EthereumAddress,
			Message: "User created successfully",
	}
	
	// send the response
	json.NewEncoder(w).Encode(res)
	
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var users models.DappUser 
	database.Db.First(&users, params["id"])
	database.Db.Delete(&users)

	json.NewEncoder(w).Encode(&users)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.DappUser
	database.Db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var user models.DappUser 
	database.Db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func GenerateUserAccessGrant(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from GenerateUserAccessGrant()")
}

