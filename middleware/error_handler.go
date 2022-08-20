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
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}
}

func (e *ErrorHandler) WebErrorHandlerForOwnerAuthMiddleware(w *http.ResponseWriter, errorMessage string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/owner/public_error.html"))

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}
}

func (e *ErrorHandler) WebErrorHandlerForOwnerPrivateRoute(w *http.ResponseWriter, errorMessage string, backUrl string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/owner/private_error.html"), path.Join("view", "/layout/owner_layout.html"))

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage
	data["backurl"] = backUrl

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

}

func (e *ErrorHandler) WebErrorHandlerForAgentPublicRoute(w *http.ResponseWriter, errorMessage string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/agent/public_error.html"))

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}
}

func (e *ErrorHandler) WebErrorHandlerForAgentAuthMiddleware(w *http.ResponseWriter, errorMessage string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/agent/public_error.html"))

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}
}

func (e *ErrorHandler) WebErrorHandlerForAgentPrivateRoute(w *http.ResponseWriter, errorMessage string, backUrl string) {

	templateErrorPublic, err := template.ParseFiles(path.Join("view", "/agent/private_error.html"), path.Join("view", "/layout/owner_layout.html"))

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

	data := make(map[string]interface{})
	data["error"] = errorMessage
	data["backurl"] = backUrl

	err = templateErrorPublic.Execute(*w, data)

	if err != nil {
		http.Error(*w, err.Error(), 500)
	}

}

func (e *ErrorHandler) FinalUnhandlerError(w *http.ResponseWriter, err error) {
	http.Error(*w, err.Error(), 500)
}
