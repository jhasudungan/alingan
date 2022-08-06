package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type TransactionManagementController struct {
	TransactionService service.TransactionService
	ProductService     service.ProductService
	AuthMiddleware     middleware.AuthMiddleware
}

type WebResponse struct {
	Status      int64  `json:"status"`
	Description string `json:"description"`
}

func (t *TransactionManagementController) ShowTransactionData(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := t.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	if isAuthenticated == false {
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	ownerId := session.Id

	transactions, err := t.TransactionService.FindTransactionByOwner(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/transaction_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["transactions"] = transactions

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (t *TransactionManagementController) ShowCreateTransactionForm(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := t.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	if isAuthenticated == false {
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	ownerId := session.Id

	products, err := t.ProductService.FindProductByOwnerId(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/point_of_sales.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["products"] = products

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (t *TransactionManagementController) HandleCreateTransactionRequest(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	requestBody := model.CreateTransactionRequest{}
	json.NewDecoder(r.Body).Decode(&requestBody)

	err := t.TransactionService.CreateTransaction(requestBody)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	response := &WebResponse{Status: int64(200), Description: "success submit transaction"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}

func (t *TransactionManagementController) ShowTransactionInformation(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := t.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	if isAuthenticated == false {
		http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
		return
	}

	params := mux.Vars(r)
	transactionId := params["transactionId"]

	transaction, err := t.TransactionService.GetTransactionInformation(transactionId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["transaction"] = transaction

	template, err := template.ParseFiles(path.Join("view", "owner/view_transaction.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}
