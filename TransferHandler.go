package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type requestTransfer struct {
	User   uint `json:"user,int"`
	From   uint `json:"from,int"`
	To     uint `json:"to,int"`
	Amount int  `json:"amount,int"`
}

//TransferHandler ...
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	request := &requestTransfer{}
	err = json.Unmarshal(body, request)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}

	//Перевод
	//...

	//Вывод
	fmt.Fprintln(w, "TransferHandler")
	fmt.Fprintln(w, request.User, request.From, request.To, request.Amount)
	fmt.Fprintln(w, http.StatusOK, "OK")
}
