package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var gitSHA string

func middleware(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-App-Version", gitSHA)
		f(w, r)
	}
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func main() {
	http.HandleFunc("/hello_world", middleware(helloWorldHandler))
	log.Info("Ready to serve!")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
