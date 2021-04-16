package dblayer

import(
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter/backend/src/models"
)

type DBLayer interface{
	GetAllProducts() ([]models.Product, error)
	GetPromos() ([]models.Product, error)
	GetCustomerByName(string, string) (models.Customer, error)
	GetCustomerByID(int) (models.Customer, error)
	GetProdict(uint) (models.Product, error)
	AddUser(models.Customer) (models.Customer, error)
	SignInUser(username, password string) (models.Customer, error)
	SingOutUserById(int) error
	GetCustomerOrdersByID(int) ([]models.Product, error)
}