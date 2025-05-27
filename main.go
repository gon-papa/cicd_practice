package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// クライアントにレスポンスを返す
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloHandler)
	port := "8080"
	log.Printf("cpi-fetcher is running on port %s...\n", port)

	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
