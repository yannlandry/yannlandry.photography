package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yannlandry/yannlandry.photography/content"
	"github.com/yannlandry/yannlandry.photography/handler"
	"github.com/yannlandry/yannlandry.photography/util"
)

func main() {
	log.Println("Starting yannlandry.photography...")

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
		log.Fatalln("`--content-path` is a required argument.")
	}
	if baseURL == "" {
		log.Fatalln("`--base-url` is a required argument.")
	}
	if staticURL == "" {
		log.Fatalln("`--static-url` is a required argument.")
	}

	// Instantiate `URLBuilder`s
	var err error
	if util.BaseURL, err = util.NewURLBuilder(baseURL); err != nil {
		log.Fatalf("Failed parsing the base URL: %s\n", err)
	}
	if util.StaticURL, err = util.NewURLBuilder(staticURL); err != nil {
		log.Fatalf("Failed parsing the static URL: %s\n", err)
	}

	// Instantiate markdown renderer
	util.Markdown = util.NewMarkdownEngine(util.BaseURL, util.StaticURL)

	// Load website content
	if err := content.Content.Load(contentPath); err != nil {
		log.Fatalf("Failed loading the website content: %s\n", err)
	}
	log.Println("Done loading configuration.")

	// Router
	router := mux.NewRouter()
	router.HandleFunc("/", handler.Home)
	router.HandleFunc("/blog", handler.Blog)
	router.HandleFunc("/blog/", handler.Blog)
	router.HandleFunc("/blog/{slug}", handler.BlogPost)
	router.HandleFunc("/blog/keyword/{keyword}", handler.BlogKeyword)
	router.HandleFunc("/{slug}", handler.Page)
	router.HandleFunc("/{slug}/", handler.Page)

	// Server
	server := &http.Server{
		Handler: router,
		Addr:    ":8080",
	}
	log.Println(server.ListenAndServe())
}
