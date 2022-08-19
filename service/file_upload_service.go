package service

import (
	"alingan/entity"
	"alingan/repository"
	"alingan/util"
	"context"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/joho/godotenv"
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
	multipartFile, fileHeader, err := r.FormFile("upload-product-image")
	defer multipartFile.Close()

	if err != nil {
		return location, err
	}

	actualFile, err := fileHeader.Open()
	defer actualFile.Close()

	if err != nil {
		return location, err
	}

	// Get The File Size
	fileSize := fileHeader.Size

	// read the file into bytes
	multipartFileBytes, err := ioutil.ReadAll(multipartFile)

	if err != nil {
		return location, err
	}

	if fileSize > int64(maximumUpload) {
		return location, errors.New("file too large")
	}

	// detect the file type
	contentType := http.DetectContentType(multipartFileBytes)

	if contentType != "image/png" && contentType != "image/jpeg" {
		return location, errors.New("image should be jpeg/png format")
	}

	// create cloudinary context
	ctx := context.Background()

	// load env for cloudinary
	err = godotenv.Load()

	if err != nil {
		return location, err
	}

	cldClouName := os.Getenv("CLOUDINARY_CLOUD_NAME")
	cldApiKey := os.Getenv("CLOUDINARY_API_KEY")
	cldApiSecret := os.Getenv("CLOUDINARY_API_SECRET")
	cldTargetFolder := os.Getenv("CLOUDINARY_PRODUCT_IMAGE_FOLDER")

	// create cloudinary instance
	cldInstance, err := cloudinary.NewFromParams(cldClouName, cldApiKey, cldApiSecret)

	if err != nil {
		return location, err
	}

	// Upload "actual file (actualFile)" into cloudinary
	cloudinaryResponse, err := cldInstance.Upload.Upload(
		ctx,
		actualFile,
		uploader.UploadParams{Folder: cldTargetFolder})

	if err != nil {
		return location, err
	}

	log.Print(cloudinaryResponse)

	// save data to db
	productImage := entity.ProductImage{}
	productImage.LocationPath = cloudinaryResponse.SecureURL
	productImage.ProductId = productId
	productImage.ProductImageId = util.GenerateId("PRD-IMG")

	err = f.ProductImageRepository.Insert(productImage)

	if err != nil {
		return location, err
	}

	return location, nil
}
