package controller

import (
	"alingan/model"
	"alingan/service"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductManagementController struct {
	ProductService service.ProductService
}

func (p *ProductManagementController) ShowProductData(w http.ResponseWriter, r *http.Request) {

	// ownerId will get from session when authentication is integrated
	ownerId := "owner-001"

	products, err := p.ProductService.FindProductByOwnerId(ownerId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/product_list.html"), path.Join("view", "layout/owner_layout.html"))

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

func (p *ProductManagementController) ShowProductInformation(w http.ResponseWriter, r *http.Request) {

	// storeId
	params := mux.Vars(r)
	productId := params["productId"]

	product, err := p.ProductService.FindProductById(productId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/view_product.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	templateData := make(map[string]interface{})

	templateData["product"] = product

	err = template.Execute(w, templateData)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (p *ProductManagementController) ShowCreateProductForm(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "owner/create_product.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

}

func (p *ProductManagementController) HandleCreateProductFormRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request := model.CreateProductRequest{}

	// we get owner id from sessions
	request.OwnerId = "owner-001"

	request.ProductName = r.Form.Get("product-name")
	request.ProductMeasurementUnit = r.Form.Get("product-measurement-unit")

	data, err := strconv.ParseFloat(r.Form.Get("product-price"), 64)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request.ProductPrice = data

	// stock deactived later
	request.ProductStock = int64(0)

	err = p.ProductService.CreateProduct(request)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/product", http.StatusSeeOther)

}

func (p *ProductManagementController) HandleInactiveProductRequest(w http.ResponseWriter, r *http.Request) {

	// storeId
	params := mux.Vars(r)
	productId := params["productId"]

	err := p.ProductService.SetProductInactive(productId)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/product", http.StatusSeeOther)

}

func (p *ProductManagementController) HandleUpdateProductRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request := model.UpdateProductRequest{}

	request.ProductName = r.Form.Get("update-product-name")
	request.ProductMeasurementUnit = r.Form.Get("update-product-measurement-unit")

	data, err := strconv.ParseFloat(r.Form.Get("update-product-price"), 64)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	request.ProductPrice = data

	// stock deactived later
	request.ProductStock = int64(0)

	err = p.ProductService.UpdateProduct(request, r.Form.Get("product-id"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/product", http.StatusSeeOther)

}
