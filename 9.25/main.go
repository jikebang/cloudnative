package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	version := os.Getenv("VERSION")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		for k, vv := range r.Header {
			for _, v := range vv {
				w.Header().Add(k, v)
			}
		}
		w.Header().Set("VERSION", version)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
