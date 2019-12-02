package main

import (
	"fmt"
	//"log"
	"database/sql"
	"04-looping/controller"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"04-looping/middleware"
	"04-looping/model"
	_ "github.com/lib/pq"

	_ "net/http/pprof"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "test"
  )

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.Startup(templates)
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8000","cert.pem","key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
	host, port, user, password, dbname)
	fmt.Println(psqlInfo)

	db, err := sql.Open("postgres",psqlInfo)
	fmt.Println("sql.Open:", err)
	if err != nil {
		panic(err)
	  }
	 	
	  err = db.Ping()
	  fmt.Println("db.Ping:", err)
	  if err != nil {
		panic(err)
	  }
	
	  fmt.Println("Successfully connected!")
	
	model.SetDatabase(db)
	return db
		
}
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	//	fmt.Println(fi.Name())
	}
	return result
}
