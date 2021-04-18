package models

import "time"

type Product struct {
	gorm.Model
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt" gorm:"column:imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	ProductName string  `json:"productname" gorm:"column:productname"`
	Description string  `json:"desc"`
}

type Customer struct {
	gorm.Model
	Name      string   `json:"name"`
	FirstName string  `gorm:"column:firstname"` `json:"firstname"`
	LastName  string  `gorm:"column:lastname"` `json:"lastname"`
	Email     string  `gorm:"column:email"` `json:"email"`
	LoggedIn  bool    `gorm:"column:loggedin"` `json:"loggedin"`
	Pass      string  `json:"password"`
	Orders    []Order `json:"orders"`
}

type Order struct {
	gorm.Model
	Product
	Customer
	CustomerID   int       `gorm:"column:customer_id" json:"customer_id"`
	ProductID    int       `gorm:"column:product_id" json:"product_id`
	Price        float64   `gorm:"column:price" json:"sell_price`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date`
}
