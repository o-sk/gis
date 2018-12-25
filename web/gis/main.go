package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/o-sk/gis"
)

func handler(w http.ResponseWriter, r *http.Request) {
	q := r.FormValue("q")
	if q == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	i, err := gis.Search(q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	j, err := json.Marshal(i)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(j)
}

func main() {
	port := os.Args[1]
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
