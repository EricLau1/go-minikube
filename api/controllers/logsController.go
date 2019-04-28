package controllers

import (
	"go-minikube/api/models"
	"go-minikube/api/utils"
	"net/http"
)

func GetLogs(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	logs, err := models.PaginateLogs(utils.PageRequest(r, 5))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, logs, http.StatusOK)
}
