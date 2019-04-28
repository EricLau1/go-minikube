package main

import (
	"go-minikube/api"
	"go-minikube/api/models"
	"go-minikube/gen"
)

func main() {

	db := models.Connect()
	db.Debug().DropTableIfExists(&models.Wallet{}, &models.Owner{}, &models.Log{})
	db.Debug().AutoMigrate(&models.Owner{}, &models.Wallet{}, &models.Log{})
	db.Close()
	gen.GenerateData()

	api.Run()
}
