package app

import "net/http"

func (a *App) loadRoutes() {
	
	a.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the home page!"))
	})

	a.router.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("About page"))
	})
}