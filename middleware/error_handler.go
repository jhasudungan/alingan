package middleware

import (
	"net/http"
	"path"
	"text/template"
)

type ErrorHandler struct{}

func (e *ErrorHandler) WebErrorHandlerForOwnerPublicRoute(w *http.ResponseWriter, errorMessage string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/owner/public_error.html"))

	if err != nil {
		panic(err.Error())
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		panic(err.Error())
	}
}

func (e *ErrorHandler) WebErrorHandlerForOwnerAuthMiddleware(w *http.ResponseWriter, errorMessage string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/owner/public_error.html"))

	if err != nil {
		panic(err.Error())
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		panic(err.Error())
	}
}

func (e *ErrorHandler) WebErrorHandlerForOwnerPrivateRoute(w *http.ResponseWriter, errorMessage string, backUrl string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/owner/private_error.html"), path.Join("view", "/layout/owner_layout.html"))

	if err != nil {
		panic(err.Error())
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage
	data["backurl"] = backUrl

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		panic(err.Error())
	}

}
