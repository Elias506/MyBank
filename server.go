package main

import (
	_ "database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type user struct {
	ID      int `json:"id,int,omitempty"`
	Balance int `json:"balance,int"`
}

type transfer struct {
	User   int `json:"user,int,omitempty"`
	From   int `json:"from,int,omitempty"`
	To     int `json:"to,int,omitempty"`
	Amount int `json:"amount,int"`
}

func handleError(w http.ResponseWriter, url string, err error, status int){
	fmt.Printf("resource %s err: %s\n", url, err)
	http.Error(w, "remote resource error", status)
}

func main() {
	/*connStr := "user=postgres password=mypass dbname=productdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DataBased error: ", err)
	}
	defer db.Close()*/
	http.HandleFunc("/balance",  balanceHandler)
	http.HandleFunc("/deposit",  depositHandler)
	http.HandleFunc("/withdraw", withdrawHandler)
	http.HandleFunc("/transfer", transferHandler)

	fmt.Println("Server is listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("user")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "User Id is %d\n", id)

	//Поиск id в БД
	fmt.Fprintln(w, "Search in DataBase...")
	//Если такого нет, создаем новый

	//Инициализируем баланс
	balance := 21
	//Подготовка к выводу

	u := &user{Balance: balance}
	uJson, err := json.Marshal(u)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, http.StatusOK, "OK")
	fmt.Fprintf(w, "%s\n", string(uJson))
}

func depositHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	t := &transfer{}
	err = json.Unmarshal(body, t)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusInternalServerError)
		return
	}
	if t.User == 0 {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, t.User, t.Amount)
	//Работа с базой данных

	fmt.Fprintln(w, http.StatusOK, "OK")
}

func withdrawHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	t := &transfer{}
	err = json.Unmarshal(body, t)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusInternalServerError)
		return
	}
	if t.User == 0 {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, t.User, t.Amount)
	//Работа с базой данных

	fmt.Fprintln(w, http.StatusOK, "OK")
}

func transferHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "transferHandler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	t := &transfer{}
	err = json.Unmarshal(body, t)
	if err != nil {
		handleError(w, r.URL.String(), err, http.StatusInternalServerError)
		return
	}
	if t.From == 0 || t.To == 0 {
		handleError(w, r.URL.String(), err, http.StatusBadRequest)
		return
	}

	//Манипуляции с переводом
	//Затем вывод

	fmt.Fprintln(w, http.StatusOK, "OK")
}
