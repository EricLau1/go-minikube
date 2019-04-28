package controllers

import (
	"go-minikube/api/utils"
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, utils.INFO, http.StatusOK)
}

func GetHelp(w http.ResponseWriter, r *http.Request) {
	utils.ToJson(w, utils.HELPS, http.StatusOK)
}
