package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
}

func Double(n int) int {
	if n == 0 {
		return 0
	} else {
		return n * 2
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test", TestHandler)
	http.Handle("/", router)
	fmt.Println("Everything is set up!")
}
