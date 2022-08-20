package controller

import (
	"alingan/middleware"
	"alingan/service"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type FileUploadController struct {
	FileUploadService service.FileUploadService
	ErrorHandler      middleware.ErrorHandler
}

func (f *FileUploadController) HandleUploadProductImageRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	productId := params["productId"]
	redirectLink := fmt.Sprintf("/owner/product/%v", productId)

	_, err := f.FileUploadService.UploadIProductmage(productId, r)

	if err != nil {
		f.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), redirectLink)
		return
	}

	http.Redirect(w, r, redirectLink, http.StatusSeeOther)

}
