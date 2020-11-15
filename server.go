package main

import (
	"./handlers"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/balance",  handlers.BalanceHandler)
	http.HandleFunc("/deposit",  handlers.DepositHandler)
	http.HandleFunc("/withdraw", handlers.WithdrawHandler)
	http.HandleFunc("/transfer", handlers.TransferHandler)

	fmt.Println(">Server is listening...")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("ListenAndServe error: ", err)
	}
}
