package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", okHandler)

	http.ListenAndServe(":4000", limit(mux))
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
