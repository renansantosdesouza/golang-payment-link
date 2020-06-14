package controllers

import (
	"html/template"
	"net/http"

	svc "github.com/payment-link/services"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	links := svc.GetAll()
	temp.ExecuteTemplate(w, "Index", links)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body := r.FormValue("BodyJson")
		svc.Create(body)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	linkId := r.URL.Query().Get("id")
	svc.Delete(linkId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	linkId := r.URL.Query().Get("id")
	link := svc.Read(linkId)
	temp.ExecuteTemplate(w, "Edit", link)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		linkId := r.FormValue("id")
		body := r.FormValue("BodyJson")
		svc.Update(linkId, body)
	}
	http.Redirect(w, r, "/", 301)
}
