package routes

import (
	"books/pkg/controllers"

	"github.com/gorilla/mux"
)

var BookstoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookbyId).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
