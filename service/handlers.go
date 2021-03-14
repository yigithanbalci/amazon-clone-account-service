package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yigithanbalci/amazon-clone-account-service/dbclient"
)

var Dbclient dbclient.IBoltClient

func GetAccount(w http.ResponseWriter, r *http.Request) {

	accountID := mux.Vars(r)["accountID"]

	account, err := Dbclient.QueryAccount(accountID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	data, _ := json.Marshal(account)
	w.Header().Set("content-type", "application/json")
	w.Header().Set("content-length", strconv.Itoa(len(data)))
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
