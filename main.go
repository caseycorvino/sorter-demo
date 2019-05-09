package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = ParseTemplates() //maybe wrong
}

func main() {
	var port = "8080"
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	log.Println("Listening...")
	error := http.ListenAndServe(":"+port, nil)
	if error != nil {
		panic(error)
		return
	}
}

func index(w http.ResponseWriter, req *http.Request) {
	error := tpl.ExecuteTemplate(w, "landing.gohtml", nil)
	//tpl.ExecuteTemplate(w,"landing.gohtml",dataStruct) pass with data
	if error != nil {
		log.Println("LOGGED", error)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func apply(w http.ResponseWriter, req *http.Request) {
	var email string
	if req.Method == http.MethodPost {
		email = req.FormValue("email")
		log.Println(email)
	}
}

/*
	Helper method to get all templates in templates folder.
*/
func ParseTemplates() *template.Template {
	templ := template.New("")
	err := filepath.Walk("./templates", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".gohtml") {
			_, err = templ.ParseFiles(path)
			if err != nil {
				log.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}

	return templ
}
