package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"paasta-monitoring-api/connections"
	apiControllerV1 "paasta-monitoring-api/controllers/api/v1"
	AP "paasta-monitoring-api/controllers/api/v1/ap"
	"paasta-monitoring-api/middlewares"
)

//SetupRouter function will perform all route operations
func SetupRouter(conn connections.Connections) *echo.Echo {
	e := echo.New()

	// Logger 설정 (HTTP requests)
	e.Use(middleware.Logger())
	// Recover 설정 (recovers panics, prints stack trace)
	e.Use(middleware.Recover())

	// CORS 설정 (control domain access)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		MaxAge:       86400,
		//AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))

	// swagger 2.0 설정
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Controller 설정
	apiToken := apiControllerV1.GetTokenController(conn)
	apiUser := apiControllerV1.GetUserController(conn)

	ApAlarm := AP.GetApAlarmController(conn)
	ApBosh := AP.GetBoshController(conn)

	// Router 설정
	//// Token은 항상 접근 가능하도록
	e.POST("/api/v1/token", apiToken.CreateToken) // 토큰 생성
	e.POST("/api/v1/token2", apiToken.CreateAccessToken)  // member_infos 기반 토큰 생성
	e.PUT("/api/v1/token", apiToken.RefreshToken) // 토큰 리프레시

	//// 그외에 다른 정보는 발급된 토큰을 기반으로 유효한 토큰을 가진 사용자만 접근하도록 middleware 설정
	//// 추가 설명 : middlewares.CheckToken 설정 (입력된 JWT 토큰 검증 및 검증된 요청자 API 접근 허용)
	//// Swagger에서는 CheckToken 프로세스에 의해 아래 function을 실행할 수 없음 (POSTMAN 이용)
	v1 := e.Group("/api/v1", middlewares.CheckToken(conn))
	v1.GET("/users", apiUser.GetUsers)

	// AP - Alarm
	v1.GET("/ap/alarm/status", ApAlarm.GetAlarmStatus)
	v1.GET("/ap/alarm/policy", ApAlarm.GetAlarmPolicy)
	v1.PUT("/ap/alarm/policy", ApAlarm.UpdateAlarmPolicy)
	v1.PUT("/ap/alarm/target", ApAlarm.UpdateAlarmTarget)
	v1.POST("/ap/alarm/sns", ApAlarm.RegisterSnsAccount)
	v1.GET("/ap/alarm/sns", ApAlarm.GetSnsAccount)
	v1.DELETE("/ap/alarm/sns", ApAlarm.DeleteSnsAccount)
	v1.PUT("/ap/alarm/sns", ApAlarm.UpdateSnsAccount)
	v1.POST("/ap/alarm/action", ApAlarm.CreateAlarmAction)
	v1.GET("/ap/alarm/action", ApAlarm.GetAlarmAction)
	v1.PATCH("/ap/alarm/action", ApAlarm.UpdateAlarmAction)
	v1.DELETE("/ap/alarm/action", ApAlarm.DeleteAlarmAction)

	// IaaS
	//IaaSModule := IAAS.GetOpenstackController()
	//v1.Get("/iaas/hyper/statistics", IaaSModule.GetHypervisorStatistics)
	e.GET("/api/v1/ap/alarm/statistics/total", ApAlarm.GetAlarmStatisticsTotal)
	e.GET("/api/v1/ap/alarm/statistics/service", ApAlarm.GetAlarmStatisticsService)
	e.GET("/api/v1/ap/alarm/statistics/resource", ApAlarm.GetAlarmStatisticsResource)

	e.GET("/api/v1/ap/bosh", ApBosh.GetBoshInfoList)
	e.GET("/api/v1/ap/bosh/overview", ApBosh.GetBoshOverview)
	e.GET("/api/v1/ap/bosh/summary", ApBosh.GetBoshSummary)
	e.GET("/api/v1/ap/bosh/process", ApBosh.GetBoshProcessByMemory)
	e.GET("/api/v1/ap/bosh/chart/:id", ApBosh.GetBoshChart)
	e.GET("/api/v1/ap/bosh/Log/:id", ApBosh.GetBoshLog)

	return e
}
