package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestDeposit struct {
	User   uint `json:"user,int"`
	Amount int  `json:"amount,int"`
}

//DepositHandler ...
func DepositHandler(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	request := &requestDeposit{}
	err = json.Unmarshal(body, request)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}

	//Работа с базой данных
	//...

	//Вывод
	fmt.Fprintln(w, "DepositHandler")
	fmt.Fprintln(w, request.User, request.Amount)
	fmt.Fprintln(w, http.StatusOK, "OK")
}
