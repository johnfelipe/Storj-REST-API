package main

import (

	// "context"
	// "log"
	// "os"

	"github.com/mohammedfajer/Storj-REST-API/database"
	"github.com/mohammedfajer/Storj-REST-API/helpers"
	"github.com/mohammedfajer/Storj-REST-API/routers"
	// "github.com/mohammedfajer/Storj-REST-API/services/app"
)

func main() {

	helpers.LoadEnv()

	// Generate restricted app access grant
	// err := app.AccessGrant(context.Background(), os.Getenv("SATELLITEADDRESS"), 
	// 	os.Getenv("APIKEY"), os.Getenv("APPPASSPHRASE"))

	// if err != nil {
	// 	log.Fatal("There was an error generating the app access grant")
	// }

	database.Setup()
	
	routers.SetupRouters()

	database.CloseDatabase()

	
}