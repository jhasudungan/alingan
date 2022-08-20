package controller

import (
	"alingan/middleware"
	"html/template"
	"log"
	"net/http"
	"path"
)

type PublicController struct {
	ErrorHandler middleware.ErrorHandler
}

func (p *PublicController) ShowIndexPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/index.html"), path.Join("view", "layout/public_layout.html"))

	if err != nil {
		p.ErrorHandler.FinalUnhandlerError(&w, err)
	}

	err = template.Execute(w, nil)

	if err != nil {
		p.ErrorHandler.FinalUnhandlerError(&w, err)
	}
}

func (p *PublicController) ShowNotFoundPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/not_found.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Parse Files ", 500)
		return
	}

	template.Execute(w, nil)

}
