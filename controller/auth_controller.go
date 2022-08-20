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
	AuthService    service.AuthService
	ErrorHandler   middleware.ErrorHandler
	AuthMiddleware middleware.AuthMiddleware
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

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
		return
	}

	http.Redirect(w, r, "/owner/registration/submit/sucess", http.StatusSeeOther)

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

func (a *AuthController) ShowRegistrationSuccessPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/success_registration.html"))

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

func (a *AuthController) ShowOwnerProfilePage(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	owner, err := a.AuthService.GetOwnerProfileInformation(session.Id)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}

	templateData := make(map[string]interface{})

	templateData["owner"] = owner

	template, err := template.ParseFiles(path.Join("view", "owner/profile.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}
	template.Execute(w, templateData)
}

func (a *AuthController) HandleOwnerUpdateProfileRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}

	request := model.UpdateOwnerProfileRequest{}
	request.OwnerId = r.Form.Get("owner-id")
	request.OwnerName = r.Form.Get("update-owner-name")
	request.OwnerType = r.Form.Get("owner-type")

	err = a.AuthService.UpdateOwnerProfile(request)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}

	http.Redirect(w, r, "/owner/profile", http.StatusSeeOther)

}

func (a *AuthController) ShowUpdatePasswordPage(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/update_password.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
	}

	err = template.Execute(w, nil)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
	}
}

func (a *AuthController) HandleOwnerUpdatePasswordRequest(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := a.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		a.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	err = r.ParseForm()

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}

	request := model.UpdateOwnerPassword{}
	request.OwnerId = session.Id
	request.NewPassword = r.Form.Get("new-password")
	request.OldPassword = r.Form.Get("old-password")

	err = a.AuthService.UpdateOwnerPassword(request)

	if err != nil {
		a.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/login")
		return
	}

	http.Redirect(w, r, "/owner/profile", http.StatusSeeOther)

}
