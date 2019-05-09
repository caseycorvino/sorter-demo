package main

import (
	"encoding/json"
	"fmt"
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

type Configuration struct {
	Port           string
	SendGridApiKey string
}
var configuration = Configuration{}

func init() {
	tpl = ParseTemplates() //maybe wrong
}

func main() {
	// Read parameters
	file, error := os.Open("config/parameters.json")
	if error != nil {
		panic(error)
		return
	}
	decoder := json.NewDecoder(file)
	error = decoder.Decode(&configuration)
	if error != nil {
		panic(error)
		return
	}

	var port = configuration.Port
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/apply", apply)
	log.Println("Listening  on port: "  + port)
	error = http.ListenAndServe(":"+port, nil)
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
	apikey := configuration.SendGridApiKey
	from := mail.NewEmail("Casey Corvino", "caseycorvino@nyu.edu")
	subject := "Apply Confirmation"
	to := mail.NewEmail("New Sorter User", email)
	plainTextContent := "Thanks for applying!"
	htmlContent := "<strong>Thanks for applying!</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apikey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}else{
		log.Println("Email Status: ",response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
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
