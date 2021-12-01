package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	fileName := "files/" + pageID + ".html"
	_,err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}

	http.ServeFile(w, r, fileName)
}

func Double(n int) int {
	if n == 0 {
		return 0
	} else {
		return n * 2
	}
}

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler)
	rtr.HandleFunc("/homepage", pageHandler)
	rtr.HandleFunc("/contact", pageHandler)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
