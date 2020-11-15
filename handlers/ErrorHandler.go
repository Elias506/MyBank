package handlers

import (
	"fmt"
	"net/http"
)

//ErrorHandler ...
func ErrorHandler(w http.ResponseWriter, url string, err error, status int){
	fmt.Printf(">Resource %s Error:\n\t %s\n", url, err)
	http.Error(w, err.Error(), status)
}