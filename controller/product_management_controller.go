package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"fmt"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductManagementController struct {
	ProductService service.ProductService
	AuthMiddleware middleware.AuthMiddleware
	ErrorHandler   middleware.ErrorHandler
}

func (p *ProductManagementController) ShowProductData(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// ownerId will get from session when authentication is integrated
	ownerId := session.Id

	products, err := p.ProductService.FindProductByOwnerId(ownerId)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/product_list.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	templateData := make(map[string]interface{})

	templateData["products"] = products

	err = template.Execute(w, templateData)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}
}

func (p *ProductManagementController) ShowProductInformation(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	productId := params["productId"]

	product, err := p.ProductService.FindProductById(productId)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/view_product.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	templateData := make(map[string]interface{})

	templateData["product"] = product

	err = template.Execute(w, templateData)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

}

func (p *ProductManagementController) ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/create_product.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

}

func (p *ProductManagementController) HandleCreateProductFormRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	request := model.CreateProductRequest{}

	// we get owner id from sessions
	request.OwnerId = session.Id

	request.ProductName = r.Form.Get("product-name")
	request.ProductMeasurementUnit = r.Form.Get("product-measurement-unit")

	data, err := strconv.ParseFloat(r.Form.Get("product-price"), 64)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	request.ProductPrice = data

	err = p.ProductService.CreateProduct(request)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	http.Redirect(w, r, "/owner/product", http.StatusSeeOther)

}

func (p *ProductManagementController) HandleInactiveProductRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	productId := params["productId"]

	err = p.ProductService.SetProductInactive(productId)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	redirectUrl := fmt.Sprintf("/owner/product/%v", productId)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

}

func (p *ProductManagementController) HandleReactiveProductRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	// storeId
	params := mux.Vars(r)
	productId := params["productId"]

	err = p.ProductService.SetProductActive(productId)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	redirectUrl := fmt.Sprintf("/owner/product/%v", productId)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

}

func (p *ProductManagementController) HandleUpdateProductRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := p.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		p.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	request := model.UpdateProductRequest{}

	request.ProductName = r.Form.Get("update-product-name")
	request.ProductMeasurementUnit = r.Form.Get("update-product-measurement-unit")

	data, err := strconv.ParseFloat(r.Form.Get("update-product-price"), 64)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	request.ProductPrice = data
	productId := r.Form.Get("product-id")

	err = p.ProductService.UpdateProduct(request, productId)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/product")
		return
	}

	redirectUrl := fmt.Sprintf("/owner/product/%v", productId)
	http.Redirect(w, r, redirectUrl, http.StatusSeeOther)

}
