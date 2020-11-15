package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestWithdraw struct {
	User   uint `json:"user,int"`
	Amount int `json:"amount,int"`
}

//WithdrawHandler ...
func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()	
	request := &requestWithdraw{}
	err = json.Unmarshal(body, request)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	//Работа с базой данных
	//...

	//Вывод
	fmt.Fprintln(w, "WithdrawHandler")
	fmt.Fprintln(w, request.User, request.Amount)
	fmt.Fprintln(w, http.StatusOK, "OK")
}