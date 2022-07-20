package controller

import (
	"alingan/core/service"
	"html/template"
	"log"
	"net/http"
	"path"
)

type OwnerController struct {
	StoreService service.StoreService
}

func (o *OwnerController) ShowStoreData(w http.ResponseWriter, r *http.Request) {

	// ownerId will get from session when authentication is integrated
	ownerId := "owner-001"

	stores, err := o.StoreService.FindStoreByOwnerId(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/store_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})
	templateData["stores"] = stores

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}
