package dblayer

import (
	"errors"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter/backend/src/models"
)

var ErrINVALIDPASSWORD = errors.New("invalid password")

type DBLayer interface {
	GetAllProduct() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProdict(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SingOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Product, error)
	AddOrder(models.Order) error
	GetCreditCardID(int) (string, error)
	SaveCreditCardForCustomer(int, string) error
}
