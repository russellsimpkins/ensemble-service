package main

import (
	"fmt"
	"github.com/russellsimpkins/ensemble"
	"net/http"
	"time"
)

func StartListener() {
	http.HandleFunc("/svc/v1/magic", ensemble.Handle)

	srv := &http.Server{
		Handler:        nil,
		Addr:           ":8002",
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		MaxHeaderBytes: 32768,
	}

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}

func main() {
	StartListener()
}
