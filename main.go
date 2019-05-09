package main

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	http.HandleFunc("/apply", apply)
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
		err := sendEmail(email)
		if err != nil {
			log.Println("LOGGED", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
	return
}

/*
	Send waitlist confirmation method.
 */
func sendEmail(email string) error{
	apikey := "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"
	from := mail.NewEmail("Casey Corvino", "caseycorvino@nyu.edu")
	subject := "Apply Confirmation"
	to := mail.NewEmail("New Sorter User", email)
	plainTextContent := "Thanks for applying!"
	htmlContent := "<p style='color: gray; border-bottom: 1px solid black'>Thanks for applying!</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apikey)
	_, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
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
