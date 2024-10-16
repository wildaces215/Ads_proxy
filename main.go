package main

import (
	"fmt"
	"net/http"
)

func main() {
	proxy := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf("Recieved request: %s %s\n", r.Method, r.URL, r.Body)
			proxyReq, err := http.NewRequest(r.Method, r.URL.String(), r.Body)
			if err != nil {
				http.Error(w, "Error creating proxy request", http.StatusInternalServerError)
			}
		}),
	}
}
