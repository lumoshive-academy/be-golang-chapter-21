package handler

import (
	"be-golang-chapter-21/impleme-http-serve/database"
	"be-golang-chapter-21/impleme-http-serve/library"
	"be-golang-chapter-21/impleme-http-serve/model"
	"be-golang-chapter-21/impleme-http-serve/repository"
	"be-golang-chapter-21/impleme-http-serve/service"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type CustomerHandler struct {
	serviceCustomer service.CustomerService
}

func NewCustomerHandler(cs service.CustomerService) CustomerHandler {
	return CustomerHandler{serviceCustomer: cs}
}

func (ch *CustomerHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	customer := model.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)

	if err != nil {
		library.BadResponse(w, err.Error())
		return
	}

	err = ch.serviceCustomer.LoginService(&customer)
	if err != nil {
		library.BadResponse(w, "Account Not Found")
		return
	}

	library.SuccessResponse(w, "Login success", customer)
}

func GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	id := query.Get("id")

	id_int, _ := strconv.Atoi(id)

	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	customer, err := serviceCustomer.CustomerByID(id_int)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    "Account Not Found",
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data:       customer,
	}
	json.NewEncoder(w).Encode(response)
}

func UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	id_int, _ := strconv.Atoi(id)

	payload, _ := io.ReadAll(r.Body)

	var customer model.Customer
	err := json.Unmarshal(payload, &customer)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	db, err := database.InitDB()
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}
	defer db.Close()

	repo := repository.NewCustomerRepository(db)
	serviceCustomer := service.NewCustomerService(repo)

	err = serviceCustomer.UpdateCustomer(id_int, customer)
	if err != nil {
		badResponse := model.Response{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		}
		json.NewEncoder(w).Encode(badResponse)
		return
	}

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "Updated",
		Data:       nil,
	}
	json.NewEncoder(w).Encode(response)

}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	response := model.Response{
		StatusCode: http.StatusOK,
		Message:    "deleted",
		Data:       id,
	}
	json.NewEncoder(w).Encode(response)

}
