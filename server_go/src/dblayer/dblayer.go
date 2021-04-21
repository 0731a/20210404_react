package dblayer

import (
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
}

func (db *DBORM) GetAllProducts() (products []models.Products, err error) {
	return products, db.Find(&products).Error // select * from products 와 동일
}

func (db *DBORM) GetPromos() (products []models.Product, err error) {
	return products, db.Where("promotion IS NOT NULL").find(&products).Error  // == select * from products where promotion IS NOT NULL
}

func (db *DBORM) GetCustomerByName(firstname string, lastname string) (customer models.Customer, err error) {
	return customer,db.Where(&models.Customer{FirstName: firstname, LastName: lastname}).Find(&customer).Error // == select * from customers where firstname="~" and lastname="~"
}

func (db *DBORM) GetCustomerByID(id int) (customer models.Customer, err error) {
	return customer, db.First(&customer, id).Error // 쿼리의 조건을 만족하는 첫번째 결과만 반환
}

func (db *DBORM) GetProduct(id int) (product models.Product, err error) {
	return product, db.First(&product, id).Error
}

func (db *DBORM) AddUser(customer models.Customer) (models.Customer, error) {
	hassPassword(&customer.Pass)
	customer.LoggedIn = true
	err := db.Create(&customer).Error
	customer.Pass = "" // 객체 반환 전 보안을 위하여 비밀번호를 지운다.
	return customer, err  // 행 생성 
}

func (db *DBORM) SigninUser(email, pass string) (customer models.Customer, err error){
	if !checkPassword(pass){
		return customer, error.New("Invalid password")
	}
	// 사용자 행을 나타내는 *gorm.DB 타입
	result := db.Table("Customer").Where(&models.Customer{Email:email}) // 질의결과 구조체 반환
	err = result.Update("loggedin", 1).Error // 해당 행 업데이트 
	if err != nil {
		return customer, err
	}
	return customer, result.Find(&customer).Error
}

func (db *DBORM) SingOutUserById(id int) error {
	customer := models.Customer{ // id에 해당하는 사용자 구조체 생성 
		Model: gorm.Model{
			ID: uint(id),
		},
	}
	return db.Table("Customers").Where(&customer).Update("loggedin",0).Error // 해당 행 업데이트 
}

func (db *DBORM) GetCustomerOrdersByID(id int) (orders []models.Order, err error){
	return orders, db.Table("orders").Select("*")
	.Joins("join customers on customers.id = customer_id")
	.Joins("join products on products.id = product_id")
	.Where("customer_id=?",id).Scan(&orders).Error  // = SELECT * FROM 'orders' join customers on customers.id = customer_id join products on products.id = product_id WHERE (customer_id="1")
}
