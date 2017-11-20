package main

import (
	"fmt"
	"flag"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, %s!", r.URL.Path[1:])
}

func main() {
	httpPort := flag.Int("port", 6821, "HTTP Server Listening Port")

	flag.Parse();

	fmt.Printf("Working on : %d", *httpPort)

	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%d", *httpPort), nil)
}