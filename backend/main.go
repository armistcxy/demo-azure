package main

import (
	"log"
	"net/http"
)

func greetHandle(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte("hi " + name))
}

func main() {
	http.HandleFunc("GET /greet/{name}", greetHandle)

	allowedOrgins := []string{"http://localhost:3000"}

	log.Fatal(http.ListenAndServe(":8080", CORS(http.DefaultServeMux, allowedOrgins)))
}

func CORS(next http.Handler, allowedOrigins []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
