package main

import (
	"be-golang-chapter-21/impleme-http-serve/database"
	"be-golang-chapter-21/impleme-http-serve/handler"
	"be-golang-chapter-21/impleme-http-serve/middleware"
	"be-golang-chapter-21/impleme-http-serve/repository"
	"be-golang-chapter-21/impleme-http-serve/service"
	"fmt"
	"net/http"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repoCustomer := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repoCustomer)
	handlerCustomer := handler.NewCustomerHandler(serviceCustomer)

	serverMux := http.NewServeMux()

	authMux := http.NewServeMux()
	authMux.HandleFunc("POST /login", handlerCustomer.LoginHandler)

	resourceMux := http.NewServeMux()
	resourceMux.HandleFunc("GET /customer_detail", handler.GetCustomerByID)
	resourceMux.HandleFunc("PUT /update_customer/{id}", handler.GetCustomerByID)
	resourceMux.HandleFunc("DELETE /delete_customer/{id}", handler.DeleteCustomer)

	role := middleware.Role(resourceMux)
	middleware := middleware.Middleware(role)

	serverMux.Handle("/", authMux)
	serverMux.Handle("/customer/", http.StripPrefix("/customer", middleware))

	fmt.Println("server started on port 8080")
	http.ListenAndServe(":8080", serverMux)
}
