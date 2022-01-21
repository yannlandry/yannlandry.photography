package handler

import (
	"net/http"

	"github.com/yannlandry/yannlandry.photography/content"
)

type ErrorPresenter struct {
	Code int
}

func Error(response http.ResponseWriter, request *http.Request, code int) {
	response.WriteHeader(code)
	ExecuteTemplate(content.Content.Error.Template, response, NewBasePresenter(&ErrorPresenter {
		Code: code,
	}))
}

func Error400(response http.ResponseWriter, request *http.Request) {
	Error(response, request, http.StatusBadRequest)
}

func Error403(response http.ResponseWriter, request *http.Request) {
	Error(response, request, http.StatusForbidden)
}

func Error404(response http.ResponseWriter, request *http.Request) {
	Error(response, request, http.StatusNotFound)
}

func Error500(response http.ResponseWriter, request *http.Request) {
	Error(response, request, http.StatusInternalServerError)
}
