package controller

import (
	"alingan/core/model"
	"alingan/core/service"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TransactionManagementController struct {
	TransactionService service.TransactionService
	ProductService     service.ProductService
}

type WebResponse struct {
	Status int64  `json:"status"`
	Desc   string `json: "desc"`
}

func (t *TransactionManagementController) ShowTransactionData(w http.ResponseWriter, r *http.Request) {

	ownerId := "owner-001"

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

	ownerId := "owner-001"

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

	response := &WebResponse{Status: int64(200), Desc: "success submit transaction"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
