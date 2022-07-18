package controller

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

type PublicController struct{}

func (p *PublicController) ShowIndexPage(w http.ResponseWriter, r *http.Request) {

	template, err := template.ParseFiles(path.Join("view", "public/index.html"))

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Parse Files ", 500)
		return
	}

	err = template.Execute(w, nil)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Something Went Wrong - Execute Template ", 500)
		return
	}
}
