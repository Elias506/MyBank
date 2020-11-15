package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type responseBalance struct {
	Balance int `json:"balance,int"`
}

//BalanceHandler ...
func BalanceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("user")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User Id is %d\n", id)

	//Поиск id в БД
	fmt.Fprintln(w, "Search in DataBase...")
	//Если такого нет, создаем новый

	//Инициализируем баланс
	balance := 21
	
	//Подготовка к выводу
	
	response := &responseBalance{Balance: balance}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		ErrorHandler(w, r.URL.String(), err, http.StatusInternalServerError)
		return
	}
	
	//Вывод
	fmt.Fprintln(w, "BalanceHandler")
	fmt.Fprintln(w, string(responseJSON))
	fmt.Fprintln(w, http.StatusOK, "OK")
	
}
