package main

import (
	"log"
	"net/http"
	"web-crud/config"
	"web-crud/controller/categoriescontroller"
	"web-crud/controller/homecontroller"
)

func main() {
	config.ConnectDB()

	// 1. Homepage
	http.HandleFunc("/", homecontroller.Welcome)

	// 2. Categories
	http.HandleFunc("/categories/", categoriescontroller.Index)
	http.HandleFunc("/categories/create", categoriescontroller.Create)
	http.HandleFunc("/categories/edit", categoriescontroller.Edit)
	http.HandleFunc("/categories/delete", categoriescontroller.Delete)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
