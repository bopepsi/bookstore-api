package routes

import (
	"github.com/bopepsi/p3-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/books/{id}", controllers.DeleteBookByID).Methods("DELETE")
	router.HandleFunc("/books/{id}", controllers.UpdateBookByID).Methods("PUT")
}
