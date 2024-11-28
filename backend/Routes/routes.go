package routes

import (
	"log"
	"net/http"

	Controllers "proyecto/Controllers"

	"github.com/gorilla/mux"
)

func wrapHandler(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		h(w, r)
	}
}

func ConductoRoutes(l *log.Logger, router *mux.Router) {
	c := Controllers.NewConductorController(l)
	router.HandleFunc("/conductor", wrapHandler(c.GetAll)).Methods("GET")
	router.HandleFunc("/conductor/{id}", wrapHandler(c.GetAllById)).Methods("GET")
	router.HandleFunc("/conductor", wrapHandler(c.Create)).Methods("POST")
	router.HandleFunc("/conductor/{id}", wrapHandler(c.Update)).Methods("PUT")
	router.HandleFunc("/conductor/{id}", wrapHandler(c.Delete)).Methods("DELETE")
}

func BusRoutes(l *log.Logger, router *mux.Router) {
	b := Controllers.NewBusController(l)
	router.HandleFunc("/bus", wrapHandler(b.GetAll)).Methods("GET")
	router.HandleFunc("/bus/{id}", wrapHandler(b.GetAllById)).Methods("GET")
	router.HandleFunc("/bus", wrapHandler(b.Create)).Methods("POST")
	router.HandleFunc("/bus/{id}", wrapHandler(b.Update)).Methods("PUT")
	router.HandleFunc("/bus/{id}", wrapHandler(b.Delete)).Methods("DELETE")
}

func RutaRoutes(l *log.Logger, router *mux.Router) {
	r := Controllers.NewRutaController(l)
	router.HandleFunc("/ruta", wrapHandler(r.GetAll)).Methods("GET")
	router.HandleFunc("/ruta/{id}", wrapHandler(r.GetAllById)).Methods("GET")
	router.HandleFunc("/ruta", wrapHandler(r.Create)).Methods("POST")
	router.HandleFunc("/ruta/{id}", wrapHandler(r.Update)).Methods("PUT")
	router.HandleFunc("/ruta/{id}", wrapHandler(r.Delete)).Methods("DELETE")
}

func AsignacionRoutes(l *log.Logger, router *mux.Router) {
	a := Controllers.NewAsignacionController(l)
	router.HandleFunc("/asignacion", wrapHandler(a.GetAll)).Methods("GET")
	router.HandleFunc("/asignacion/{id}", wrapHandler(a.GetAllById)).Methods("GET")
	router.HandleFunc("/asignacion", wrapHandler(a.Create)).Methods("POST")
	router.HandleFunc("/asignacion/{id}", wrapHandler(a.Update)).Methods("PUT")
	router.HandleFunc("/asignacion/{id}", wrapHandler(a.Delete)).Methods("DELETE")
}
