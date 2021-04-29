package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

//go:embed assets/shipwright-logo-lightbg-512.png
var asset []byte

func main() {
	port := 8080
	if strValue, ok := os.LookupEnv("PORT"); ok {
		if intValue, err := strconv.Atoi(strValue); err == nil {
			port = intValue
		}
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "image/png")

		if _, err := w.Write(asset); err != nil {
			log.Printf("failed to write to HTTP response: %v", err)
		}
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
