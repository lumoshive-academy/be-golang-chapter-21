package main

import (
	"be-golang-chapter-21/impleme-http-serve/handler"
	"be-golang-chapter-21/impleme-http-serve/middleware"
	"fmt"
	"net/http"
)

func main() {
	serverMux := http.NewServeMux()

	serverMux.HandleFunc("POST /login", handler.LoginHandler)

	muxwithmiddleware := http.NewServeMux()
	muxwithmiddleware.HandleFunc("GET /customer_detail", handler.GetCustomerByID)

	role := middleware.Role(muxwithmiddleware)
	middleware := middleware.Middleware(role)

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", middleware)
}
