package repository

import (
	"be-golang-chapter-21/impleme-http-serve/model"
	"database/sql"
)

type CustomerRepository struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return CustomerRepository{DB: db}
}

func (cr *CustomerRepository) Login(customer *model.Customer) error {
	query := `SELECT username, password, email FROM customers WHERE username=$1 AND password=$2`
	err := cr.DB.QueryRow(query, customer.Username, customer.Password).Scan(&customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		return err
	}
	return nil
}

func (cr *CustomerRepository) CustomerByID(id int) (*model.Customer, error) {
	customer := model.Customer{}
	query := `SELECT username, password, email FROM customers WHERE id=$1`
	err := cr.DB.QueryRow(query, id).Scan(&customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (cr *CustomerRepository) Update(id int, data *model.Customer) (int, error) {
	query := `UPDATE SET username=$1 WHERE id=$1`
	result, err := cr.DB.Exec(query, data.Username, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return int(rowsAffected), nil
}
