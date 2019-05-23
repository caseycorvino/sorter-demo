package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/minio/minio-go"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var tpl *template.Template

type Configuration struct {
	Port           string
	SendGridApiKey string
	S3AccessKeyId  string
	S3SecretKeyId  string
	CSVBucket  	   string
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
	http.HandleFunc("/email-preview", emailPreview)
	http.HandleFunc("/upload", uploadToS3)
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

func emailPreview(w http.ResponseWriter, req *http.Request) {
	//tpl.ExecuteTemplate(w,"landing.gohtml",dataStruct) pass with data
	t, error := template.ParseFiles("templates/emails/email-apply.gohtml")
	if error != nil {
		panic(error)
		return
	}
	error = t.Execute(w, struct { Email string}{"caseycorvino@nyu.edu"})
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
		err := sendEmail(email, email)
		if err != nil {
			log.Println("LOGGED", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		err = sendEmail(email, "caseycorvino@nyu.edu")
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
func sendEmail(email string, reciever string) error{
	apikey := configuration.SendGridApiKey
	from := mail.NewEmail("Casey Corvino", "caseycorvino@nyu.edu")
	subject := "Apply Confirmation"
	to := mail.NewEmail("New Sorter User", reciever)
	plainTextContent := "Thanks for applying!"
	data := struct { Email string }{email}
	htmlContent := templateToString("templates/emails/email-apply.gohtml", data)
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



func uploadToS3(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		//get form data
		err := req.ParseMultipartForm(1000 << 20)
		if err != nil {
			fmt.Println(err)
			return
		}
		file, handler, err := req.FormFile("csv-file")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		s3Client, err := minio.New("s3.amazonaws.com", configuration.S3AccessKeyId, configuration.S3SecretKeyId, true)
		if err != nil {
			log.Fatalln(err)
		}

		t := time.Now()
		n, err := s3Client.PutObject(configuration.CSVBucket, t.Format("2006-01-02 15:04:05")+".csv", file, handler.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()
		log.Println("Uploaded", t.Format("2006-01-02 15:04:05"), " of size: ", n, "Successfully.")
		data := struct { Status string }{"Success"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(data)
		if err != nil{
			log.Fatalln(err)
		}
		return
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
func templateToString(filePath string, data interface{}) string{
	t, error := template.ParseFiles(filePath)
	if error != nil {
		panic(error)
		return "error"
	}
	buffer := new(bytes.Buffer)
	if error = t.Execute(buffer, data)
	error != nil {
		panic(error)
		return "error"
	}
	return buffer.String()
}
