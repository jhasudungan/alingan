package controller

import (
	"alingan/middleware"
	"alingan/model"
	"alingan/service"
	"html/template"
	"net/http"
	"path"
)

type AuthController struct {
	AuthService  service.AuthService
	ErrorHandler middleware.ErrorHandler
}

func (a *AuthController) ShowLoginForm(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "owner/login.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

}

func (a *AuthController) ShowRegistrationForm(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "owner/registration.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

}

func (a *AuthController) HandleLoginFormRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	request := model.OwnerLoginRequest{}
	request.OwnerEmail = r.Form.Get("owner-email-login")
	request.OwnerPassword = r.Form.Get("owner-email-password")

	session, err := a.AuthService.OwnerLogin(request)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "alingan-session",
		Value:   session.Token,
		Expires: session.Expiry,
		Path:    "/",
	})

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)
}

func (a *AuthController) HandleRegistrationFormRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	request := model.OwnerRegistrationRequest{}
	request.OwnerEmail = r.Form.Get("owner-email")
	request.OwnerName = r.Form.Get("owner-name")
	request.OwnerType = r.Form.Get("account-type")
	request.Password = r.Form.Get("owner-password")

	err = a.AuthService.OwnerRegistration(request)

	http.Redirect(w, r, "/owner/store", http.StatusSeeOther)
}
