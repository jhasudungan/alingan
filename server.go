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
	reportRepo := &repository.ReportRepositoryImpl{}

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

	reportService := &service.ReportServiceImpl{
		ReportRepository: reportRepo,
	}

	sessionList := make(map[string]*model.Session)
	agentSessionList := make(map[string]*model.AgentSession)

	authSvc := &service.AuthServiceImpl{
		OwnerRepo:        ownerRepo,
		AgentRepo:        agentRepo,
		SessionList:      sessionList,
		AgentSessionList: agentSessionList,
		JoinRepo:         joinRepo,
	}

	fileUploadService := &service.FileUploadServiceImpl{
		ProductImageRepository: productImageRepo,
	}

	authMiddleware := middleware.AuthMiddleware{
		SessionList:      sessionList,
		AgentSessionList: agentSessionList,
	}

	errorHandler := middleware.ErrorHandler{}

	// controller
	storeManagementController := &controller.StoreManagementController{
		AuthMiddleware: authMiddleware,
		StoreService:   storeSvc,
		ErrorHandler:   errorHandler,
	}

	productManagementController := &controller.ProductManagementController{
		AuthMiddleware: authMiddleware,
		ProductService: productSvc,
		ErrorHandler:   errorHandler,
	}

	agentManagamentController := &controller.AgentManagamentController{
		AuthMiddleware: authMiddleware,
		AgentService:   agentSvc,
		StoreService:   storeSvc,
		ErrorHandler:   errorHandler,
	}

	transactionManagementController := &controller.TransactionManagementController{
		AuthMiddleware:     authMiddleware,
		TransactionService: transactionSvc,
		StoreService:       storeSvc,
		ProductService:     productSvc,
		AgentService:       agentSvc,
		ErrorHandler:       errorHandler,
	}

	authController := &controller.AuthController{
		AuthService:    authSvc,
		ErrorHandler:   errorHandler,
		AuthMiddleware: authMiddleware,
	}

	fileUploadController := &controller.FileUploadController{
		FileUploadService: fileUploadService,
		ErrorHandler:      errorHandler,
	}

	agentTransactionController := &controller.AgentTransactionController{
		TransactionService: transactionSvc,
		ProductService:     productSvc,
		AuthMiddleware:     authMiddleware,
		ErrorHandler:       errorHandler,
	}

	ownerController := &controller.OwnerController{
		ReportService:  reportService,
		ErrorHandler:   errorHandler,
		AuthMiddleware: authMiddleware,
	}

	publicController := &controller.PublicController{
		ErrorHandler: errorHandler,
	}

	// router and handler
	r := &mux.Router{}
	r.HandleFunc("/", publicController.ShowIndexPage).Methods("GET")
	r.HandleFunc("/about", publicController.ShowAboutPage).Methods("GET")
	r.NotFoundHandler = http.HandlerFunc(publicController.ShowNotFoundPage)

	r.HandleFunc("/owner/guide", ownerController.ShowGuidePage).Methods("GET")
	r.HandleFunc("/owner/mobile", ownerController.ShowMobileAppPage).Methods("GET")

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
	r.HandleFunc("/owner/reactive/agent/{agentId}", agentManagamentController.HandleReactiveAgentRequest).Methods("GET")
	r.HandleFunc("/owner/update/agent/submit", agentManagamentController.HandleUpdateAgentRequest).Methods("POST")

	r.HandleFunc("/owner/transaction", transactionManagementController.ShowTransactionData).Methods("GET")
	r.HandleFunc("/owner/new/transaction", transactionManagementController.ShowCreateTransactionForm).Methods("GET")
	r.HandleFunc("/owner/new/transaction/submit", transactionManagementController.HandleCreateTransactionRequest).Methods("POST")
	r.HandleFunc("/owner/transaction/{transactionId}", transactionManagementController.ShowTransactionInformation).Methods("GET")

	r.HandleFunc("/owner/dashboard", ownerController.ShowDashboard).Methods("GET")
	r.HandleFunc("/owner/login", authController.ShowLoginForm).Methods("GET")
	r.HandleFunc("/owner/login/submit", authController.HandleLoginFormRequest).Methods("POST")
	r.HandleFunc("/owner/registration", authController.ShowRegistrationForm).Methods("GET")
	r.HandleFunc("/owner/registration/submit", authController.HandleRegistrationFormRequest).Methods("POST")
	r.HandleFunc("/owner/registration/submit/success", authController.ShowRegistrationSuccessPage).Methods("GET")
	r.HandleFunc("/owner/profile", authController.ShowOwnerProfilePage).Methods("GET")
	r.HandleFunc("/owner/update/profile/submit", authController.HandleOwnerUpdateProfileRequest).Methods("POST")
	r.HandleFunc("/owner/update/password", authController.ShowUpdatePasswordPage).Methods("GET")
	r.HandleFunc("/owner/update/password/submit", authController.HandleOwnerUpdatePasswordRequest).Methods("POST")

	r.HandleFunc("/owner/upload/product-image/{productId}", fileUploadController.HandleUploadProductImageRequest).Methods("POST")

	r.HandleFunc("/agent/login", authController.ShowAgentLoginForm).Methods("GET")
	r.HandleFunc("/agent/login/submit", authController.HandleAgentLoginFormRequest).Methods("POST")
	r.HandleFunc("/agent/new/transaction", agentTransactionController.ShowCreateTransactionForm).Methods("GET")
	r.HandleFunc("/agent/transaction", agentTransactionController.ShowAgentAndStoreTransactionList).Methods("GET")
	r.HandleFunc("/agent/transaction/{transactionId}", agentTransactionController.ShowTransactionInformation).Methods("GET")

	r.HandleFunc("/owner/logout", authController.HandleOwnerLogOutRequest).Methods("GET")
	r.HandleFunc("/agent/logout", authController.HandleAgentLogOutRequest).Methods("GET")

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
