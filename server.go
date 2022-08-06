package main

import (
	"alingan/controller"
	"alingan/middleware"
	"alingan/model"
	"alingan/repository"
	"alingan/service"
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
	transactionRepo := &repository.TransactionRepositoryImpl{}
	transactionItemRepo := &repository.TransactionItemRepositoryImpl{}
	productImageRepo := &repository.ProductImageRepositoryImpl{}

	// svc
	storeSvc := &service.StoreServiceImpl{
		OwnerRepo: ownerRepo,
		StoreRepo: storeRepo,
	}

	productSvc := &service.ProductServiceImpl{
		OwnerRepo:        ownerRepo,
		ProductRepo:      productRepo,
		ProductImageRepo: productImageRepo,
	}

	agentSvc := &service.AgentServiceImpl{
		OwnerRepo: ownerRepo,
		JoinRepo:  joinRepo,
		AgentRepo: agentRepo,
	}

	transactionSvc := &service.TransactionServiceImpl{
		StoreRepo:           storeRepo,
		OwnerRepo:           ownerRepo,
		ProductRepo:         productRepo,
		TransactionRepo:     transactionRepo,
		JoinRepo:            joinRepo,
		AgentRepo:           agentRepo,
		TransactionItemRepo: transactionItemRepo,
	}

	sessionList := make(map[string]*model.Session)

	authSvc := &service.AuthServiceImpl{
		OwnerRepo:   ownerRepo,
		AgentRepo:   agentRepo,
		SessionList: sessionList,
	}

	fileUploadService := &service.FileUploadServiceImpl{
		ProductImageRepository: productImageRepo,
	}

	authMiddleware := middleware.AuthMiddleware{
		SessionList: sessionList,
	}

	// controller
	storeManagementController := &controller.StoreManagementController{
		AuthMiddleware: authMiddleware,
		StoreService:   storeSvc,
	}

	productManagementController := &controller.ProductManagementController{
		AuthMiddleware: authMiddleware,
		ProductService: productSvc,
	}

	agentManagamentController := &controller.AgentManagamentController{
		AuthMiddleware: authMiddleware,
		AgentService:   agentSvc,
		StoreService:   storeSvc,
	}

	transactionManagementController := &controller.TransactionManagementController{
		AuthMiddleware:     authMiddleware,
		TransactionService: transactionSvc,
		ProductService:     productSvc,
	}

	authController := &controller.AuthController{
		AuthService: authSvc,
	}

	fileUploadController := &controller.FileUploadController{
		FileUploadService: fileUploadService,
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
	r.HandleFunc("/owner/reactive/store/{storeId}", storeManagementController.HandleReactiveStoreRequest).Methods("GET")
	r.HandleFunc("/owner/update/store/submit", storeManagementController.HandleUpdateStoreRequest).Methods("POST")

	r.HandleFunc("/owner/product", productManagementController.ShowProductData).Methods("GET")
	r.HandleFunc("/owner/product/{productId}", productManagementController.ShowProductInformation).Methods("GET")
	r.HandleFunc("/owner/new/product", productManagementController.ShowCreateProductForm).Methods("GET")
	r.HandleFunc("/owner/new/product/submit", productManagementController.HandleCreateProductFormRequest).Methods("POST")
	r.HandleFunc("/owner/inactive/product/{productId}", productManagementController.HandleInactiveProductRequest).Methods("GET")
	r.HandleFunc("/owner/reactive/product/{productId}", productManagementController.HandleReactiveProductRequest).Methods("GET")
	r.HandleFunc("/owner/update/product/submit", productManagementController.HandleUpdateProductRequest).Methods("POST")

	r.HandleFunc("/owner/agent", agentManagamentController.ShowAgentData).Methods("GET")
	r.HandleFunc("/owner/agent/{agentId}", agentManagamentController.ShowAgentInformation).Methods("GET")
	r.HandleFunc("/owner/new/agent", agentManagamentController.ShowCreateAgentForm).Methods("GET")
	r.HandleFunc("/owner/new/agent/submit", agentManagamentController.HandleCreateAgentFormRequest).Methods("POST")
	r.HandleFunc("/owner/inactive/agent/{agentId}", agentManagamentController.HandleSetAgentInactiveRequest).Methods("GET")
	r.HandleFunc("/owner/reactive/agent/{agentId}", agentManagamentController.HandleReactiveActiveRequest).Methods("GET")

	r.HandleFunc("/owner/transaction", transactionManagementController.ShowTransactionData).Methods("GET")
	r.HandleFunc("/owner/new/transaction", transactionManagementController.ShowCreateTransactionForm).Methods("GET")
	r.HandleFunc("/owner/new/transaction/submit", transactionManagementController.HandleCreateTransactionRequest).Methods("POST")
	r.HandleFunc("/owner/transaction/{transactionId}", transactionManagementController.ShowTransactionInformation).Methods("GET")

	r.HandleFunc("/owner/login", authController.ShowLoginForm).Methods("GET")
	r.HandleFunc("/owner/login/submit", authController.HandleLoginFormRequest).Methods("POST")
	r.HandleFunc("/owner/registration", authController.ShowRegistrationForm).Methods("GET")
	r.HandleFunc("/owner/registration/submit", authController.HandleRegistrationFormRequest).Methods("POST")

	r.HandleFunc("/owner/upload/product-image/{productId}", fileUploadController.HandleUploadProductImageRequest).Methods("POST")

	// file server
	assetFileServer := http.FileServer(http.Dir("asset"))
	productImageFileServer := http.FileServer(http.Dir("uploaded/product-image"))

	// file server handler
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static", assetFileServer))
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources", productImageFileServer))

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
