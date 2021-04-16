package rest

import (
	"github.com/gin-gonic/gin"
)


// RESTful API의 진입점 함수
// -> 내부에 RESTful API의 라우팅을 정의한다.
func RunAPI(address string) error{ // address = RESTful API 서버의 주소
	r := gin.Default() // Gin을 사용하기 위한 Gin 엔진 객체, URL을 정의하고 HTTP메서드를 지정할 때 사용한다.
	r.GET("/relativepath/to/url", func(c *gin.Context){ // /relativepath/to/url(상대경로)에 대한 GET 요청 처리  // *gin.Context: Gin 프레임워크가 제공, 해당 객체는 요청 확인과 처리 응답에 필요한 기능을 제공한다.
		// 로직의 구현 
	})
	r.GET("/products", func(c, *gin.Context) {
		// 클라이언트에게 상품 목록 반환
	})
	r.GET("/promos", func(c, *gin.Context) {
		// 클라이언트에게 프로모션 목록 반환
	})
	r.POST("/users/signin", func(c, *gin.Context) {
		// 사용자 로그인
	})
	r.POST("users", func(c, *gin.Context){
		// 사용자 추가 
	})
	r.POST("user/:id/signout", func(c, *gin.Context){ // 해당 경로는 사용자 ID를 포함한다 ( :id -> 변수 id 의미 ) // ID는 사용자마다 고유한 값이기 때문에 와일드카드(*) 사용
		// 해당 ID의 사용자 로그 아웃
	})
	r.GET("users/:id/orders", func(c, *gin.Context){
		// 해당 ID의 사용자 주문 내역 조회 
	})
	r.POST("user/charge", func(c, *gin.Context){
		// 신용카드 결제 처리 
	})
}