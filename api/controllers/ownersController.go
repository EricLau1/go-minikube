package controllers

import (
	"encoding/json"
	"go-minikube/api/models"
	"go-minikube/api/models/validators"
	"go-minikube/api/utils"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetOwners(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	owners, err := models.PaginateOwners(utils.PageRequest(r, 5))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, owners, http.StatusOK)
}

func GetOwner(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	owner, err := models.GetOwnerById(utils.ExtractId(r))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, owner, http.StatusOK)
}

func PostOwner(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	body, _ := ioutil.ReadAll(r.Body)
	var owner models.Owner
	err := json.Unmarshal(body, &owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	owner, err = validators.ValidateOwner(owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	response, err := models.NewOwner(owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, response, http.StatusCreated)
}

func PutOwner(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	keys := r.URL.Query()
	disable, _ := strconv.ParseBool(keys.Get("disable"))
	if disable {
		rows, err := models.DisableOwner(utils.ExtractId(r))
		if err != nil {
			utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		utils.ToJson(w, rows, http.StatusOK)
		return
	}
	var owner models.Owner
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	owner, err = validators.ValidateOwner(owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	owner.ID = utils.ExtractId(r)
	rows, err := models.UpdateOwner(owner)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusOK)
}
