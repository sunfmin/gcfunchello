package gcfunchello

import (
	"fmt"
	"net/http"
)

func Mux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello")
	if err != nil {
		panic(err)
	}
}
