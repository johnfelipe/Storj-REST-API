package routers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mohammedfajer/Storj-REST-API/services/user"
)

func SetupRouters() {
	router  := mux.NewRouter()

	router.Schemes("https")

	// router.HandleFunc("/login", iam.Login)
	// router.HandleFunc("/home", iam.Home)
	// router.HandleFunc("/refresh", iam.Refresh)

	router.HandleFunc("/api/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/api/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", user.DeleteUser).Methods("DELETE")

	// router.HandleFunc("/api/users/{id}", user.GenerateUserAccessGrant).Methods("GET")
	// router.HandleFunc("/api/users/{id}/lists", user.List).Methods("GET")
	// router.HandleFunc("/api/users/{id}/downloads", user.Download).Methods("GET")
	// router.HandleFunc("/api/users/{id}/shares", user.Grant).Methods("GET")
	
	// router.HandleFunc("/api/users/{id}/uploads", user.Upload).Methods("POST")
	// router.HandleFunc("/api/users/{id}/revokes", user.Revoke).Methods("POST")
	// router.HandleFunc("/api/users/{id}/deletes", user.Delete).Methods("DELETE")


	fmt.Println("\t* Running on http://localhost:8000/ (Prese CTRL+C to quit) ")
	log.Fatal(http.ListenAndServe(":8000", router))
}