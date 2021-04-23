package rest

// handler: 클라이언트의 요청을 처리한다.

import (
	"log"
	"net/http"
	"strconv"
	models
	"example.com/m/src/dblayer"
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
	db, err := dblayer.NewORM(db, constring)
	if err != nil {
		return nil, err
	}
	return &Handler{
		db: db,
	}, nil
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
	customer, err = h.db.SignInUser(customer.Email, customer.Pass)
	if err != nil {
		//잘못된 패스워드 인 경우 forbidden http 에러 반환
		if err == dblayer.ErrINVALIDPASSWORD {
			c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		}
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "server database error"})
		return
	}
	// Go 구조체 정의 및 초기화.
	request := struct {
		models.Order
		Remember    bool   `json:"rememberCard"`
		UseExisting bool   `json:"useExisting"`
		Token       string `json:"token"`
	}{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, request)
		return
	}
	stripe.key = "sk_test_4eC39HqLyjWDarjtT1zdp7dc"
	chargeP := &stripe.ChargeParams{
		//요청에 명시된 판매 가격
		Amount: stripe.Int64(int64(request.Price)),
		//결제 통화
		Currency: stripe.String("usd"),
		// 설명
		Description: stripe.String("GoMusic.. Charge.."),
	}
	// 스트라이프 사용자 id 초기화
	stripeCustomerID := ""
	if request.UseExisting {
		//저장된 카드 사용
		log.Println("Getting credit card id...")
		// 스트라이프 사용자 id를 데이터베이스에서 조회하는 메서드
		stripeCustomerID, err = h.db.GetCreditCardID(request.CustomerID)
		if err != nil {
			log.Println(err)
			c.Json(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	} else {
		cp := &stripe.CustomerParams()
		cp.SetSource(request.Token)
		customer, err := customer.New(cp)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		stripeCustomerID = customer.ID
	}
	if request.Rememeber {
		//스트라이프 사용자 id를 저장하고 데이터베이스에 저장된 사용자  id와 연결
		err = h.db.SaveCreditCardForCustomer(request.CustomerID, stripeCustomerID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	//결제 요청
	// 동일 상품 주문 여부에 대한 확인 없이 새로운 주문으로 가정
	//*stripel.ChargeParams 타입 인스턴스에 스트라이프 사용자 id 설정
	chargeP.Customer = stripe.String(stripeCustomerID)
	// 신용 카드 결제 요청
	_, err = charge.New(chargeP)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 주문 내용 데이터 베이스에 저장
	err = h.db.AddOrder(request.Order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

}
