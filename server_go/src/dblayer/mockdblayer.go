package dblayer

import (
	"encoding/json"
	"fmt"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter07/backend/src/models"
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

func NewMockDBLayerWithData() *MockDBLayer {
	PRODUCTS := `[
		{
				"ID": 1,
				"CreatedAt": "2018-08-14T07:54:19Z",
				"UpdatedAt": "2019-01-11T00:28:40Z",
				"DeletedAt": null,
				"img": "img/strings.png",
				"small_img": "img/img-small/strings.png",
				"imgalt": "string",
				"price": 100,
				"promotion": 0,
				"productname": "Strings",
				"Description": ""
		},
		{
				"ID": 2,
				"CreatedAt": "2018-08-14T07:54:20Z",
				"UpdatedAt": "2019-01-11T00:29:11Z",
				"DeletedAt": null,
				"img": "img/redguitar.jpeg",
				"small_img": "img/img-small/redguitar.jpeg",
				"imgalt": "redg",
				"price": 299,
				"promotion": 240,
				"productname": "Red Guitar",
				"Description": ""
		},
		{
				"ID": 3,
				"CreatedAt": "2018-08-14T07:54:20Z",
				"UpdatedAt": "2019-01-11T22:05:42Z",
				"DeletedAt": null,
				"img": "img/drums.jpg",
				"small_img": "img/img-small/drums.jpg",
				"imgalt": "drums",
				"price": 17000,
				"promotion": 0,
				"productname": "Drums",
				"Description": ""
		},
		{
				"ID": 4,
				"CreatedAt": "2018-08-14T07:54:20Z",
				"UpdatedAt": "2019-01-11T00:29:53Z",
				"DeletedAt": null,
				"img": "img/flute.jpeg",
				"small_img": "img/img-small/flute.jpeg",
				"imgalt": "flute",
				"price": 210,
				"promotion": 190,
				"productname": "Flute",
				"Description": ""
		},
		{
				"ID": 5,
				"CreatedAt": "2018-08-14T07:54:20Z",
				"UpdatedAt": "2019-01-11T00:30:12Z",
				"DeletedAt": null,
				"img": "img/blackguitar.jpeg",
				"small_img": "img/img-small/blackguitar.jpeg",
				"imgalt": "Black guitar",
				"price": 200,
				"promotion": 0,
				"productname": "Black Guitar",
				"Description": ""
		},
		{
				"ID": 6,
				"CreatedAt": "2018-08-14T07:54:20Z",
				"UpdatedAt": "2019-01-11T00:30:35Z",
				"DeletedAt": null,
				"img": "img/saxophone.jpeg",
				"small_img": "img/img-small/saxophone.jpeg",
				"imgalt": "Saxophone",
				"price": 1000,
				"promotion": 980,
				"productname": "Saxophone",
				"Description": ""
		}
]
`

	ORDERS := `[
	{
			"ID": 1,
			"CreatedAt": "2018-12-29T23:35:36Z",
			"UpdatedAt": "2018-12-29T23:35:36Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 1,
			"product_id": 1,
			"sell_price": 90,
			"purchase_date": "2018-12-29T23:34:32Z"
	},
	{
			"ID": 2,
			"CreatedAt": "2018-12-29T23:35:48Z",
			"UpdatedAt": "2018-12-29T23:35:48Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 1,
			"product_id": 2,
			"sell_price": 299,
			"purchase_date": "2018-12-29T23:34:53Z"
	},
	{
			"ID": 3,
			"CreatedAt": "2018-12-29T23:35:57Z",
			"UpdatedAt": "2018-12-29T23:35:57Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 1,
			"product_id": 3,
			"sell_price": 16000,
			"purchase_date": "2018-12-29T23:35:05Z"
	},
	{
			"ID": 4,
			"CreatedAt": "2018-12-29T23:36:18Z",
			"UpdatedAt": "2018-12-29T23:36:18Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 2,
			"product_id": 1,
			"sell_price": 95,
			"purchase_date": "2018-12-29T23:36:18Z"
	},
	{
			"ID": 5,
			"CreatedAt": "2018-12-29T23:36:39Z",
			"UpdatedAt": "2018-12-29T23:36:39Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 2,
			"product_id": 2,
			"sell_price": 299,
			"purchase_date": "2018-12-29T23:36:39Z"
	},
	{
			"ID": 6,
			"CreatedAt": "2018-12-29T23:38:13Z",
			"UpdatedAt": "2018-12-29T23:38:13Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 2,
			"product_id": 4,
			"sell_price": 205,
			"purchase_date": "2018-12-29T23:37:01Z"
	},
	{
			"ID": 7,
			"CreatedAt": "2018-12-29T23:38:19Z",
			"UpdatedAt": "2018-12-29T23:38:19Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 3,
			"product_id": 4,
			"sell_price": 210,
			"purchase_date": "2018-12-29T23:37:28Z"
	},
	{
			"ID": 8,
			"CreatedAt": "2018-12-29T23:38:28Z",
			"UpdatedAt": "2018-12-29T23:38:28Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 3,
			"product_id": 5,
			"sell_price": 200,
			"purchase_date": "2018-12-29T23:37:41Z"
	},
	{
			"ID": 9,
			"CreatedAt": "2018-12-29T23:38:32Z",
			"UpdatedAt": "2018-12-29T23:38:32Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 3,
			"product_id": 6,
			"sell_price": 1000,
			"purchase_date": "2018-12-29T23:37:54Z"
	},
	{
			"ID": 10,
			"CreatedAt": "2019-01-13T00:44:55Z",
			"UpdatedAt": "2019-01-13T00:44:55Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 19,
			"product_id": 6,
			"sell_price": 1000,
			"purchase_date": "2018-12-29T23:37:54Z"
	},
	{
			"ID": 11,
			"CreatedAt": "2019-01-14T06:03:08Z",
			"UpdatedAt": "2019-01-14T06:03:08Z",
			"DeletedAt": null,
			"img": "",
			"small_img": "",
			"imgalt": "",
			"price": 0,
			"promotion": 0,
			"productname": "",
			"Description": "",
			"name": "",
			"firstname": "",
			"lastname": "",
			"email": "",
			"password": "",
			"loggedin": false,
			"orders": null,
			"customer_id": 1,
			"product_id": 3,
			"sell_price": 17000,
			"purchase_date": "0001-01-01T00:00:00Z"
	}
]
`
	CUSTOMERS := `[
	{
			"ID": 1,
			"CreatedAt": "2018-08-14T07:52:54Z",
			"UpdatedAt": "2019-01-13T22:00:45Z",
			"DeletedAt": null,
			"name": "",
			"firstname": "Mal",
			"lastname": "Zein",
			"email": "mal.zein@email.com",
			"password": "$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC",
			"loggedin": true,
			"orders": null
	},
	{
			"ID": 2,
			"CreatedAt": "2018-08-14T07:52:55Z",
			"UpdatedAt": "2019-01-12T22:39:01Z",
			"DeletedAt": null,
			"name": "",
			"firstname": "River",
			"lastname": "Sam",
			"email": "river.sam@email.com",
			"password": "$2a$10$mNbCLmfCAc0.4crDg3V3fe0iO1yr03aRfE7Rr3vdfKMGVnnzovCZq",
			"loggedin": false,
			"orders": null
	},
	{
			"ID": 3,
			"CreatedAt": "2018-08-14T07:52:55Z",
			"UpdatedAt": "2019-01-13T21:56:05Z",
			"DeletedAt": null,
			"name": "",
			"firstname": "Jayne",
			"lastname": "Ra",
			"email": "jayne.ra@email.com",
			"password": "$2a$10$ZeZI4pPPlQg89zfOOyQmiuKW9Z7pO9/KvG7OfdgjPAZF0Vz9D8fhC",
			"loggedin": false,
			"orders": null
	},
	{
			"ID": 19,
			"CreatedAt": "2019-01-13T08:43:44Z",
			"UpdatedAt": "2019-01-13T15:12:25Z",
			"DeletedAt": null,
			"name": "",
			"firstname": "John",
			"lastname": "Doe",
			"email": "john.doe@bla.com",
			"password": "$2a$10$T4c8rmpbgKrUA0sIqtHCaO0g2XGWWxFY4IGWkkpVQOD/iuBrwKrZu",
			"loggedin": false,
			"orders": null
	}]`

	var products []models.Product
	var customers []models.Customer
	var orders []models.Order
	json.Unmarshal([]byte(PRODUCTS), &products)
	json.Unmarshal([]byte(CUSTOMERS), &customers)
	json.Unmarshal([]byte(ORDERS), &orders)
	return NewMockDBLayer(products, customers, orders)
}

func (mock *MockDBLayer) GetProduct(id int) (models.Product, error) {
	result := models.Product{}
	if mock.err != nil {
		return result, mock.err
	}
	for _, product := range mock.products {
		if product.ID == uint(id) {
			return product, nil
		}
	}
	return result, fmt.Errorf("Could not find product with id %d", id)
}
