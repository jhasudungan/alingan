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

func (a *AuthController) ShowAgentLoginForm(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "agent/login.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPublicRoute(&w, err.Error())
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPublicRoute(&w, err.Error())
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

func (a *AuthController) HandleAgentLoginFormRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPublicRoute(&w, err.Error())
		return
	}

	request := model.AgentLoginRequest{}
	request.AgentEmail = r.Form.Get("agent-email-login")
	request.AgentPassword = r.Form.Get("agent-email-password")

	session, err := a.AuthService.AgentLogin(request)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForAgentPublicRoute(&w, err.Error())
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "alingan-session",
		Value:   session.Token,
		Expires: session.Expiry,
		Path:    "/",
	})

	http.Redirect(w, r, "/agent/new/transaction", http.StatusSeeOther)
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

func (a *AuthController) HandleOwnerLogOutRequest(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("alingan-session")

	ownerToken := c.Value

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	c = &http.Cookie{
		Name:   "deleted-cookie",
		MaxAge: 0,
		Value:  "",
	}

	http.SetCookie(w, c)

	a.AuthService.OwnerLogout(ownerToken)

	http.Redirect(w, r, "/owner/login", http.StatusSeeOther)
}

func (a *AuthController) HandleAgentLogOutRequest(w http.ResponseWriter, r *http.Request) {

	c, err := r.Cookie("alingan-session")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	agentToken := c.Value

	c = &http.Cookie{
		Name:   "deleted-cookie",
		MaxAge: 0,
		Value:  "",
	}

	http.SetCookie(w, c)

	a.AuthService.AgentLogout(agentToken)

	http.Redirect(w, r, "/agent/login", http.StatusSeeOther)
}
