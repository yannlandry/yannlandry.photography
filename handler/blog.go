package handler

import (
	"io"
	"net/http"
)

func Blog(response http.ResponseWriter, request *http.Request) {
	io.WriteString(response, "Blog")
}
