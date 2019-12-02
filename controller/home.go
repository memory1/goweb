package controller

import (
	"log"
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

	if pusher, ok :=w.(http.Pusher); ok {
		log.Printf("pusher supported")
		pusher.Push("/css/app.css", &http.PushOptions{Header: http.Header{"Content-Type":[]string{"text/css"}},})
	}
	
	vm := viewmodel.NewHome()
	w.Header().Add("Content-Type", "text/html")
	h.homeTemplate.Execute(w,vm)
}