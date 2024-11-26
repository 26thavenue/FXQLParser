package app

import (
	"net/http"

	_ "github.com/26thavenue/FXQLParser/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func (a *App) loadRoutes() {
	
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})

	a.router.HandleFunc("/swagger/",httpSwagger.WrapHandler)

	a.router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	})
}