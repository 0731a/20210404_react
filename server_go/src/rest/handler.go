package rest

// handler: 클라이언트의 요청을 처리한다.

import (
	"net/http"
	"strconv"

	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/dblayer"
	"github.com/PacktPublishing/Hands-On-Full-Stack-Development-with-Go/Chapter06/backend/src/models"
	"github.com/gin-gonic/gin"
)

// 핸들러의 모든 메서드를 포함하는 인터페이스, 확장성을 높이기 위함
type handerInterface interface {
	GetProducts(c *gin.Context)
	GetPromos(c *gin.Context)
	AddUser(c *gin.Context)
	SignIn(c *gin.Context)
	SignOut(c *gin.Context)
	GetOrders(c *gin.Context)
	Charge(c *gin.Context)
}

// 모든 메서드가 있는 Handler 구조체 정의
// Handler 타입은 데이터를 읽거나 수정하기 때문에 데이터베이스 레이어인터페이스에 접근 가능해야 함
type Handler struct {
	db dblayer.DBlayer
}

// Handler 생성자
func NewHandler() (*Handler, error) {
	// Handler 객체에 대한 포인터 생성
	return new(Handler), nil
}

func (h *Handler) GetProducts(c *gin.Context) {
	if h.db == nil {
		return
	}
	products, err := h.db.GetAllProducts()
	if err != nil {
		/* 첫번째 인자는 HTTP 상태코드, 두번째는 응답의 바디 */
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *Handler) SignIn(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer) // http 요청 body에서 JSON 문서를 추출 후 지정된 객체(*models.Customer)로 디코딩
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}
	customer, err = h.db.SignInUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) AddUser(c *gin.Context) {
	if h.db == nil {
		return
	}
	var customer models.Customer
	err := c.ShouldBindJSON(&customer) // http 요청 body에서 JSON 문서를 추출 후 지정된 객체(*models.Customer)로 디코딩
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}
	customer, err = h.db.AddUser(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customer)
}

func (h *Handler) SignOut(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p) // 문자형, 정수형으로 변환
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.db.SignOutUserById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}

func (h *Handler) GetOrders(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p) // 문자형, 정수형으로 변환
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errror": err.Error()})
		return
	}
	orders, err := h.db.GetCustomerOrdersByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orders)
}

func (h *Handler) Charge(c *gin.Context) {
	if h.db == nil {
		return
	}
}
