package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func BlogPost(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)
	slug := variables["slug"]

	io.WriteString(response, fmt.Sprintf("Blog Post: %s", slug))
}
