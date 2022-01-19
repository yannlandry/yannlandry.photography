package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
	"github.com/yannlandry/yannlandry.photography/handler"
	"github.com/yannlandry/yannlandry.photography/util"
)

func main() {
	fmt.Fprintln(os.Stderr, "Starting yannlandry.photography...")

	// Command-line arguments
	var contentPath string
	flag.StringVar(&contentPath, "content-path", "", "Path to the repository defining the website's content")
	var baseURL string
	flag.StringVar(&baseURL, "base-url", "", "URL of the website to be prepended to links")
	var staticURL string
	flag.StringVar(&staticURL, "static-url", "", "URL of static assets to be prepended to images, stylesheets, etc.")
	flag.Parse()

	// Check command-line arguments
	if contentPath == "" {
		fmt.Fprintln(os.Stderr, "`--content-path` is a required argument.")
		os.Exit(2)
	}
	if baseURL == "" {
		fmt.Fprintln(os.Stderr, "`--base-url` is a required argument.")
		os.Exit(2)
	}
	if staticURL == "" {
		fmt.Fprintln(os.Stderr, "`--static-url` is a required argument.")
		os.Exit(2)
	}

	// Instantiate `URLBuilder`s
	var err error
	if util.BaseURL, err = util.NewURLBuilder(baseURL); err != nil {
		fmt.Fprintf(os.Stderr, "Failed parsing the base URL: %s\n", err)
		os.Exit(1)
	}
	if util.StaticURL, err = util.NewURLBuilder(staticURL); err != nil {
		fmt.Fprintf(os.Stderr, "Failed parsing the static URL: %s\n", err)
		os.Exit(1)
	}

	// Instantiate markdown renderer
	util.Markdown = util.NewMarkdownEngine(util.BaseURL, util.StaticURL)

	// Load website content
	if err := content.Content.Load(contentPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed loading the website content: %s\n", err)
		os.Exit(1)
	}

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Home)
	router.HandleFunc("/blog", handler.Blog)
	router.HandleFunc("/blog/{slug}", handler.BlogPost)
	router.HandleFunc("/blog/keyword/{keyword}", handler.BlogKeyword)
	router.HandleFunc("/{slug}", handler.Page)

	// Server
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	fmt.Fprintln(os.Stderr, server.ListenAndServe())
}
