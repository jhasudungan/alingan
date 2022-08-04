package service

import (
	"alingan/entity"
	"alingan/repository"
	"alingan/util"
	"errors"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

type FileUploadService interface {
	UploadIProductmage(productId string, r *http.Request) (string, error)
}

type FileUploadServiceImpl struct {
	ProductImageRepository repository.ProductImageRepository
}

func (f *FileUploadServiceImpl) UploadIProductmage(productId string, r *http.Request) (string, error) {
	location := ""

	// Define the maximum upload size
	maximumUpload := 2 * 1024 * 1024 // 2 MB maximum file upload

	// Check if the form can be parsed
	err := r.ParseMultipartForm(int64(maximumUpload))

	if err != nil {
		return location, err
	}

	// Get The File
	file, fileHeader, err := r.FormFile("upload-product-image")
	defer file.Close()

	if err != nil {
		return location, err
	}

	// Get The File Size
	fileSize := fileHeader.Size

	// read the file into bytes
	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {
		return location, err
	}

	if fileSize > int64(maximumUpload) {
		return location, errors.New("file too large")
	}

	// detect the file type
	contentType := http.DetectContentType(fileBytes)

	if contentType != "image/png" && contentType != "image/jpeg" {
		return location, err
	}

	// directory uploaded
	dirUploaded := "./uploaded/product-image"

	// new file name
	newName := util.GenerateId("PRDIMG")

	// get the file extension
	readContentTypes, err := mime.ExtensionsByType(contentType)

	if err != nil {
		return location, err
	}

	fileExtension := readContentTypes[0]

	// saving path : directory/newName.fileExtension
	savingPath := filepath.Join(dirUploaded, newName+fileExtension)

	// create the file
	newFile, err := os.Create(savingPath)
	defer newFile.Close()

	if err != nil {
		return location, err
	}

	// write the bytes to that file
	_, err = newFile.Write(fileBytes)

	if err != nil {
		return location, err
	}

	location = "http://localhost:8080/resources/" + newName + fileExtension

	productImage := entity.ProductImage{}
	productImage.ProductId = productId
	productImage.LocationPath = location
	productImage.ProductImageId = newName

	err = f.ProductImageRepository.Insert(productImage)

	if err != nil {
		return location, err
	}

	return location, nil
}
