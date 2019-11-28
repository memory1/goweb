package controller

import (
	"net/http"
	"html/template"
)

var (
	homeController home
	shopController shop
	loginController login
	standLocatorController standLocator 
)

func Startup(templates map[string]*template.Template) {
	homeController.homeTemplate = templates["home.html"]
	shopController.shopTemplate = templates["shop.html"]
	shopController.categoryTemplate = templates["shop_details.html"]
	standLocatorController.standLocatorTemplate = templates["stand-locator.html"]
	loginController.loginTemplate = templates["login.html"]
	homeController.registerRoutes()
	shopController.registerRoutes()
	loginController.registerRoutes()
	standLocatorController.registerRoutes()
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
}