package main; routes

// Go lang has absolute paths
import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/deepanshumehtaa/bookstore-go/pkg/controller"
)

var RegisterBookStoreRoutes = func(router *mux.Router){
	router.HandleFunc("/book/", controller.CreateBook).Methods("POST")  // CreateBook is Func
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{book_id}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{book_id}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{book_id}", controller.DeleteBook).Methods("DELETE")
}

// fmt.Println("hI")

/*

https://perennialsky.medium.com/understand-handle-handler-and-handlefunc-in-go-e2c3c9ecef03

For any server two things are very important:
1. Routes 
2.

*/