package controller

import (
	"alingan/middleware"
	"html/template"
	"net/http"
	"path"
)

type PublicController struct {
	ErrorHandler middleware.ErrorHandler
}

func (p *PublicController) ShowIndexPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/index.html"), path.Join("view", "layout/public_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
	}

	err = template.Execute(w, nil)

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
	}
}

func (p *PublicController) ShowNotFoundPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/not_found.html"), path.Join("view", "layout/public_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
	}

	template.Execute(w, nil)

}

func (p *PublicController) ShowAboutPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/about.html"), path.Join("view", "layout/public_layout.html"))

	if err != nil {
		p.ErrorHandler.WebErrorHandlerForOwnerPublicRoute(&w, err.Error())
	}

	template.Execute(w, nil)

}
