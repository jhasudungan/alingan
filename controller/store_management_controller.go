package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"html/template"
	"log"
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

type StoreManagementController struct {
	StoreService   service.StoreService
	AuthMiddleware middleware.AuthMiddleware
	ErrorHandler   middleware.ErrorHandler
}

func (o *StoreManagementController) ShowStoreData(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// ownerId will get from session when authentication is integrated
	ownerId := session.Id

	stores, err := o.StoreService.FindStoreByOwnerId(ownerId)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/store_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	templateData := make(map[string]interface{})

	templateData["stores"] = stores

	err = template.Execute(w, templateData)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

}

func (o *StoreManagementController) ShowStoreInformation(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	storeId := params["storeId"]

	store, err := o.StoreService.FindStoreById(storeId)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/view_store.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	templateData := make(map[string]interface{})

	templateData["store"] = store

	err = template.Execute(w, templateData)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

}

func (o *StoreManagementController) ShowCreateStoreForm(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/create_store.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

}

func (o *StoreManagementController) HandleCreateStoreFormRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	request := model.CreateStoreRequest{}
	// we get owner id from sessions
	request.OwnerId = session.Id
	request.StoreName = r.Form.Get("store-name")
	request.StoreAddress = r.Form.Get("store-address")

	err = o.StoreService.CreateStore(request)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)

}

func (o *StoreManagementController) HandleInactiveStoreRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	storeId := params["storeId"]

	err = o.StoreService.SetStoreInactive(storeId)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)

}

func (o *StoreManagementController) HandleReactiveStoreRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	storeId := params["storeId"]

	err = o.StoreService.SetStoreActive(storeId)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)

}

func (o *StoreManagementController) HandleUpdateStoreRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request := model.UpdateStoreRequest{}
	request.StoreAddress = r.Form.Get("update-store-address")
	request.StoreName = r.Form.Get("update-store-name")

	err = o.StoreService.UpdateStore(request, r.Form.Get("store-id"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/store")
		return
	}

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)

}
