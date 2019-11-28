package controller

import (
	"strconv"
	"regexp"
	"fmt"
	"04-looping/viewmodel"
	"html/template"
	"net/http"
	"04-looping/model"
)

type shop struct {
	shopTemplate *template.Template
	categoryTemplate *template.Template
}

func (s shop) registerRoutes()  {
	http.HandleFunc("/shop",s.handleShop)
	http.HandleFunc("/shop/",s.handleShop)
}

func (s shop) handleShop(w http.ResponseWriter, r *http.Request) {
	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches :=categoryPattern.FindStringSubmatch(r.URL.Path)
	fmt.Println("URL.Path: ", r.URL.Path)
	fmt.Println("matches: ", matches)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		s.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		w.Header().Add("Content-Type", "text/html")
		s.shopTemplate.Execute(w,vm)
	}
}

func (s shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	
	vm := viewmodel.NewShopDetail(products)
	w.Header().Add("Content-Type", "text/html")
	s.categoryTemplate.Execute(w, vm)
}