package main

import (
	"alingan/core/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	publicController := &controller.PublicController{}

	r := &mux.Router{}

	r.HandleFunc("/", publicController.ShowIndexPage)

	assetFileServer := http.FileServer(http.Dir("asset"))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", assetFileServer))

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Println("application started at port :8080")

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
