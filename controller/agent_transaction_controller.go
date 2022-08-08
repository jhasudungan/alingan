package controller

import (
	"alingan/middleware"
	"alingan/service"
	"html/template"
	"log"
	"net/http"
	"path"
)

type AgentTransactionController struct {
	TransactionService service.TransactionService
	ProductService     service.ProductService
	ErrorHandler       middleware.ErrorHandler
}

func (a *AgentTransactionController) ShowCreateTransactionForm(w http.ResponseWriter, r *http.Request) {

	// Get Owner Id & session id from session
	ownerId := "owner-001"
	storeId := "str-001"
	agentId := "AGT454497b9-74d1-4bb0-8753-962a962e31f6"

	products, err := a.ProductService.FindProductByOwnerId(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}

	templateData := make(map[string]interface{})

	templateData["products"] = products
	templateData["storeId"] = storeId
	templateData["agentId"] = agentId

	template, err := template.ParseFiles(path.Join("view", "agent/point_of_sales.html"), path.Join("view", "layout/agent_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), 500)
	}
}
