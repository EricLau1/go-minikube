package controllers

import (
	"encoding/json"
	"go-minikube/api/models"
	"go-minikube/api/utils"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetWallets(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	wallets, err := models.PaginateWallets(utils.PageRequest(r, 5))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallets, http.StatusOK)
}

func GetWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	wallet, err := models.GetWalletById(utils.ExtractId(r))
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusBadRequest)
		return
	}
	utils.ToJson(w, wallet, http.StatusOK)
}

func PutWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	body, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	wallet.PublicKey = utils.Trim(wallet.PublicKey)
	rows, err := models.UpdateWallet(wallet, models.SUM)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusOK)
}

func PostWallet(w http.ResponseWriter, r *http.Request) {
	utils.HttpInfo(r)
	body, _ := ioutil.ReadAll(r.Body)
	var wallet models.Wallet
	err := json.Unmarshal(body, &wallet)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	vars := mux.Vars(r)
	rows, err := models.Transfer([]string{utils.Trim(wallet.PublicKey), utils.Trim(vars["public_key"])}, wallet.Cash)
	if err != nil {
		utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}
	utils.ToJson(w, rows, http.StatusCreated)
}
