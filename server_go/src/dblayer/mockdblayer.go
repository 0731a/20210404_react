package dblayer

import (
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/tree/master/Chapter08/backend/src/models"
)

type MockDBLayer struct {
	err       error // 에러를 발생시키는 시나리오 일때 설정
	products  []models.Product
	customers []models.Customer
	orders    []models.Order
}

// 모의 객체 생성자
func NewMockDBLayer(products []models.Product, customers []models.Customer, orders []models.Order) *MockDBLayer {
	return &MockDBLayer{
		products:  products,
		customers: customers,
		orders:    orders,
	}
}
