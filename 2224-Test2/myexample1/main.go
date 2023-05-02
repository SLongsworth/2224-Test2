package main

import (
	"log"
	"net/http"
)

func message1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("The cat is on the fence")
		next.ServeHTTP(w, r)
		log.Print("THe cat is walking on the fence")
	})
}

func message2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("The dog is jumping on the fence")
		if r.URL.Path == "/dog" {
			return
		}

		next.ServeHTTP(w, r)
		log.Print("The dog is barking at the cat")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Print("The animals are staring at eachother")
	w.Write([]byte("Animal Behaviour"))
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", message1(message2(finalHandler)))

	log.Print("Listening on :5000...")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)
}
