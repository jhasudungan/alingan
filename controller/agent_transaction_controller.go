package controller

import (
	"alingan/middleware"
	"alingan/service"
	"html/template"
	"net/http"
	"path"
)

type AgentTransactionController struct {
	TransactionService service.TransactionService
	ProductService     service.ProductService
	AuthMiddleware     middleware.AuthMiddleware
	ErrorHandler       middleware.ErrorHandler
}

func (a *AgentTransactionController) ShowCreateTransactionForm(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := a.AuthMiddleware.AuthenticateAgent(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForAgentAuthMiddleware(&w, "authentication failed")
		return
	}

	// Get Owner Id & session id from session
	ownerId := session.OwnerId
	storeId := session.StoreId
	agentId := session.Id

	products, err := a.ProductService.FindProductByOwnerId(ownerId)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPrivateRoute(&w, err.Error(), "/agent/new/transaction")
		return
	}

	templateData := make(map[string]interface{})

	templateData["products"] = products
	templateData["storeId"] = storeId
	templateData["agentId"] = agentId

	template, err := template.ParseFiles(path.Join("view", "agent/point_of_sales.html"), path.Join("view", "layout/agent_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPrivateRoute(&w, err.Error(), "/agent/new/transaction")
		return
	}

	err = template.Execute(w, templateData)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPrivateRoute(&w, err.Error(), "/agent/new/transaction")
		return
	}
}
