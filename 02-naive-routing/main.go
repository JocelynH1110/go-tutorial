package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", http.HandlerFunc(HandleRequest)))
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		w.Write([]byte(`
			<h1>Welcome to Jocelyn's website</h1>
			<a href="/contact">contact</a>
		`))
	case "/contact":
		w.Write([]byte(`
			<h1>information</h1>
			<a href="/">Back to HomePage</a> 
		`))
	default:
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`
			<h1>404 Not Found</h1>
			<p> Can't find this page!!</p>
			<a href="/">Back to homepage</a>
		`))
	}
}
