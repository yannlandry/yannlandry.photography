package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
	"github.com/yannlandry/yannlandry.photography/handler"
)

func main() {
	fmt.Fprintln(os.Stderr, "Starting yannlandry.photography...")

	// Command-line arguments
	var contentPath string
	flag.StringVar(&contentPath, "content-path", "", "Path to the repository defining the website's content")
	flag.Parse()
	if contentPath == "" {
		fmt.Fprintln(os.Stderr, "`--content-path` is a required argument.")
		os.Exit(2)
	}

	// Load website content
	if err := content.Content.Load(contentPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading the website content: %s\n", err)
		os.Exit(1)
	}

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Home)
	router.HandleFunc("/blog", handler.Blog)
	router.HandleFunc("/blog/{slug}", handler.BlogPost)
	router.HandleFunc("/blog/keyword/{keyword}", handler.BlogKeyword)
	router.HandleFunc("/{page}", handler.Page)

	// Server
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	fmt.Fprintln(os.Stderr, server.ListenAndServe())
}
