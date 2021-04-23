package rest

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error{
	h, err := NewHandler()
	if err != nil {
		return error
	}
	return RunAPIWithHandler(address, h)
}


// RESTful API의 진입점 함수
// -> 내부에 RESTful API의 라우팅을 정의한다.
func RunAPIWithHandler(address string, h HandlerInterface) error{ // address = RESTful API 서버의 주소
	r := gin.Default() // Gin을 사용하기 위한 Gin 엔진 객체, URL을 정의하고 HTTP메서드를 지정할 때 사용한다.
	r.Use(MyCustomLogger())
	/*
		r.GET("/relativepath/to/url", func(c *gin.Context){ // /relativepath/to/url(상대경로)에 대한 GET 요청 처리  // *gin.Context: Gin 프레임워크가 제공, 해당 객체는 요청 확인과 처리 응답에 필요한 기능을 제공한다.
			// 로직의 구현 
		})
		
		- 핸들러 등록 ->

		r.GET("/relativepath/to/url", h.GetRelative)`
	*/
	r.GET("/products", h.GetProducts)
	r.GET("/promos", h.GetPromos)
	userGroup := r.Group("/user"){ // 그룹 라우팅: URL의 일부를 공유하는 경우 사용 
		userGroup.POST("/:id/signout", h.SignOut )
		userGroup.GET("/:id/orders", h.GetOrders)
	}
	usersGroup := r.Group("/users"){
		usersGroup.POST("/charge", h.Charge )
		usersGroup.POST("/signin", h.SignIn )
		usersGroup.POST("", h.AddUser )
	}
	
	
	
	r.GET("users/:id/orders", h.GetOrders )


	// Restful API 서버가 HTTP 클라이언트의 요청을 기다리도록 반드시 API 핸들러와 라우팅 정의 뒤에 호출 해야한다.
	return r.RunTLS(address, "cert.pem", "key.pem")
}

func MyCustomMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 요청 처리 전 실행
		c.Set("v","123") // c.Get("v")로 값 확인 가능
		c.Next() // 요청 처리
		// 아래 코드부터 핸들러 실행 종료시 실행
		status := c.Writer.Status() // 응답 코드 확인
		// status를 이용하는 코드 작성 가능 
	}
}

func MyCustomLogger() gin.HandlerFunc {
	return func(C *gin.Context){
		fmt.Println("**********************************")
		c.Next()
		fmt.Println("*****************************************")
	}
}