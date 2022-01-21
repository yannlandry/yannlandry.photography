package handler

import (
	"html/template"
	"log"
	"net/http"
)

func ExecuteTemplate(tmp *template.Template, response http.ResponseWriter, presenter interface{}) {
	if err := tmp.Execute(response, presenter); err != nil {
		log.Println(err)
	}
}
