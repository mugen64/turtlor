package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/mugen64/turtlor/configs"
	"github.com/mugen64/turtlor/internal/app/ui/home"
)

//go:embed static
var static embed.FS

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	homePage := home.Index(config)
	pagesHandler := http.NewServeMux()
	pagesHandler.Handle("/", templ.Handler(homePage))
	pagesHandler.Handle("/static/", http.FileServer(http.FS(static)))

	log.Println("server running on -> http://localhost:9090")
	log.Fatalln(http.ListenAndServe(":9090", pagesHandler))

}
