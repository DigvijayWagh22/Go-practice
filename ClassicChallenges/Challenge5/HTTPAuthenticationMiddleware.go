package main

import (
	"fmt"
	"net/http"
)

const validToken = "secret"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("X-Auth-Token")
		if token != validToken {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})

}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You are authorized!")
}

func SetupServer() http.Handler {
	mux := http.NewServeMux()
	//public
	mux.HandleFunc("/hello", helloHandler)

	secureRoute := http.HandlerFunc(secureHandler)
	mux.Handle("/secure", AuthMiddleware(secureRoute))
	return mux
}
func main() {
	http.ListenAndServe(":8080", SetupServer())
}
