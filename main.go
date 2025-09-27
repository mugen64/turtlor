package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mugen64/turtlor/components"
)

//go:embed static
var static embed.FS

func main() {
	homePage := components.Index()
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Println("server running on -> http://localhost:9090")
	log.Fatalln(http.ListenAndServe(":9090", pagesHandler))

}
