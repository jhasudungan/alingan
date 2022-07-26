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
	productRepo := &repository.ProductRepositoryImpl{}
	joinRepo := &repository.JoinRepositoryImpl{}
	agentRepo := &repository.AgentRepositoryImpl{}

	// svc
	storeSvc := &service.StoreServiceImpl{
		OwnerRepo: ownerRepo,
		StoreRepo: storeRepo,
	}

	productSvc := &service.ProductServiceImpl{
		OwnerRepo:   ownerRepo,
		ProductRepo: productRepo,
	}

	agentSvc := &service.AgentServiceImpl{
		OwnerRepo: ownerRepo,
		JoinRepo:  joinRepo,
		AgentRepo: agentRepo,
	}

	transactionSvc := &service.TransactionServiceImpl{
		OwnerRepo:   ownerRepo,
		JoinRepo:    joinRepo,
		AgentRepo:   agentRepo,
		ProductRepo: productRepo,
		StoreRepo:   storeRepo,
	}

	// controller
	storeManagementController := &controller.StoreManagementController{
		StoreService: storeSvc,
	}

	productManagementController := &controller.ProductManagementController{
		ProductService: productSvc,
	}

	agentManagamentController := &controller.AgentManagamentController{
		AgentService: agentSvc,
		StoreService: storeSvc,
	}

	transactionManagementController := &controller.TransactionManagementController{
		TransactionService: transactionSvc,
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

	r.HandleFunc("/owner/product", productManagementController.ShowProductData).Methods("GET")
	r.HandleFunc("/owner/product/{productId}", productManagementController.ShowProductInformation).Methods("GET")
	r.HandleFunc("/owner/new/product", productManagementController.ShowCreateProductForm).Methods("GET")
	r.HandleFunc("/owner/new/product/submit", productManagementController.HandleCreateProductFormRequest).Methods("POST")
	r.HandleFunc("/owner/inactive/product/{productId}", productManagementController.HandleInactiveProductRequest).Methods("GET")
	r.HandleFunc("/owner/update/product/submit", productManagementController.HandleUpdateProductRequest).Methods("POST")

	r.HandleFunc("/owner/agent", agentManagamentController.ShowAgentData).Methods("GET")
	r.HandleFunc("/owner/agent/{agentId}", agentManagamentController.ShowAgentInformation).Methods("GET")
	r.HandleFunc("/owner/new/agent", agentManagamentController.ShowCreateAgentForm).Methods("GET")
	r.HandleFunc("/owner/new/agent/submit", agentManagamentController.HandleCreateAgentFormRequest).Methods("POST")
	r.HandleFunc("/owner/inactive/agent/{agentId}", agentManagamentController.HandleSetAgentInactiveRequest).Methods("GET")

	r.HandleFunc("/owner/transaction", transactionManagementController.ShowTransactionData).Methods("GET")

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
