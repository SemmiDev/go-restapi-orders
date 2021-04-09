package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func api(router *mux.Router) {
	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderId}", getOrder).Methods("GET")
	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderId}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderId}", deleteOrder).Methods("DELETE")
	initDB()
}

func main() {
	router := mux.NewRouter()
	api(router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("SERVER START AT PORT 8080")
	}else {
		log.Fatal("ups")
	}
}