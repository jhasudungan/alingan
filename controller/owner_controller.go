package controller

import (
	"alingan/middleware"
	"alingan/service"
	"html/template"
	"net/http"
	"path"
)

type OwnerController struct {
	ReportService  service.ReportService
	AuthMiddleware middleware.AuthMiddleware
	ErrorHandler   middleware.ErrorHandler
}

func (o *OwnerController) ShowDashboard(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, session := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	ownerId := session.Id

	dashboardReportData, err := o.ReportService.WebDashboardReport(ownerId)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}

	templateData := make(map[string]interface{})
	templateData["dashboardReports"] = dashboardReportData

	template, err := template.ParseFiles(path.Join("view", "owner/dashboard.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}

	err = template.Execute(w, templateData)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}

}

func (o *OwnerController) ShowGuidePage(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/guide.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}
}

func (o *OwnerController) ShowMobileAppPage(w http.ResponseWriter, r *http.Request) {

	isAuthenticated, err, _ := o.AuthMiddleware.AuthenticateOwner(r)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	if isAuthenticated == false {
		o.ErrorHandler.WebErrorHandlerForOwnerAuthMiddleware(&w, err.Error())
		return
	}

	template, err := template.ParseFiles(path.Join("view", "owner/mobile.html"), path.Join("view", "layout/owner_layout.html"))

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		o.ErrorHandler.WebErrorHandlerForOwnerPrivateRoute(&w, err.Error(), "/owner/dashboard")
		return
	}
}
