package gcfunchello

import (
	"fmt"
	"net/http"
)

func Mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/", home)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello\n")
	if err != nil {
		panic(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Home\n")
	if err != nil {
		panic(err)
	}
}
