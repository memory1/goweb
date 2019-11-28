package controller

import (
	"04-looping/model"
	"log"
	"fmt"
	"html/template"
	"net/http"
	"04-looping/viewmodel"
)

type login struct {
	loginTemplate *template.Template

}

func (l login) registerRoutes()  {
	http.HandleFunc("/login", l.handleLogin)
	fmt.Println("login registerRoutes")
}

func (l login) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
	    if err != nil {
			fmt.Println("Error log in: ", err)
		}
		f := r.Form
		email := f.Get("email")
		pass := f.Get("password")
		if user, err := model.Login(email, pass); err !=nil {
			log.Printf("User has logged in: %v\n",user)
			http.Redirect(w,r,"/home", http.StatusTemporaryRedirect)
	    	return
		} else {
			log.Printf("Failed to log user in with email: %v\n", email)
			vm.Email = email
			vm.Password = pass 
		}
	}
	w.Header().Add("Content-Type", "text/html")
	l.loginTemplate.Execute(w,vm)
  }