package controller

import (
	"alingan/core/service"
	"html/template"
	"log"
	"net/http"
	"path"
)

type TransactionManagementController struct {
	TransactionService service.TransactionService
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
