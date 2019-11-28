package controller

import (
	"04-looping/viewmodel"
	"html/template"
	"net/http"
)

type home struct {
	homeTemplate *template.Template

}

func (h home) registerRoutes()  {
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/",h.handleHome)
	
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w,vm)
}