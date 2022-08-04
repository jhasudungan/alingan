package controller

import (
	"alingan/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type FileUploadController struct {
	FileUploadService service.FileUploadService
}

func (f *FileUploadController) HandleUploadProductImageRequest(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	productId := params["productId"]

	_, err := f.FileUploadService.UploadIProductmage(productId, r)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Exceute Render", 500)
		return
	}

	http.Redirect(w, r, "/owner/product", http.StatusSeeOther)

}
