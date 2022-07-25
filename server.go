package main

import (
	"alingan/core/controller"
	"alingan/core/repository"
	"alingan/core/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// repository
	storeRepo := &repository.StoreRepositoryImpl{}
	ownerRepo := &repository.OwnerRepositoryImpl{}

	// svc
	storeSvc := &service.StoreServiceImpl{
		OwnerRepo: ownerRepo,
		StoreRepo: storeRepo,
	}

	// controller
	storeManagementController := &controller.StoreManagementController{
		StoreService: storeSvc,
	}
	publicController := &controller.PublicController{}

	// router and handler
	r := &mux.Router{}
	r.HandleFunc("/", publicController.ShowIndexPage)
	r.HandleFunc("/owner/store", storeManagementController.ShowStoreData).Methods("GET")
	r.HandleFunc("/owner/store/{storeId}", storeManagementController.ShowStoreInformation).Methods("GET")
	r.HandleFunc("/owner/new/store", storeManagementController.ShowCreateStoreForm).Methods("GET")
	r.HandleFunc("/owner/new/store/submit", storeManagementController.HandleCreateStoreFormRequest).Methods("POST")
	r.HandleFunc("/owner/inactive/store/{storeId}", storeManagementController.HandleInactiveStoreRequest).Methods("GET")
	r.HandleFunc("/owner/update/store/submit", storeManagementController.HandleUpdateStoreRequest).Methods("POST")

	// file server
	assetFileServer := http.FileServer(http.Dir("asset"))

	// file server handler
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", assetFileServer))

	// web server setup
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
