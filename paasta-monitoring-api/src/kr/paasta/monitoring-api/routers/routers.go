package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"paasta-monitoring-api/connections"
	apiControllerV1 "paasta-monitoring-api/controllers/api/v1"
	AP "paasta-monitoring-api/controllers/api/v1/ap"
	commonModule "paasta-monitoring-api/controllers/api/v1/common"
	CP "paasta-monitoring-api/controllers/api/v1/cp"
	iaas "paasta-monitoring-api/controllers/api/v1/iaas"
	saas "paasta-monitoring-api/controllers/api/v1/saas"
	"paasta-monitoring-api/middlewares"
	"time"
)

//SetupRouter function will perform all route operations
func SetupRouter(conn connections.Connections) *echo.Echo {
	e := echo.New()

	// Logger 설정 (HTTP requests)
	/*
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}\n",
		}))
	*/

	// Recover 설정 (recovers panics, prints stack trace)
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		LogLevel:          log.ERROR,
		DisablePrintStack: true,
		DisableStackAll:   true,
	}))

	e.Use(middleware.RequestID())
	e.Use(middlewares.Logger(conn.Logger))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, values middleware.RequestLoggerValues) error {
			conn.Logger.WithFields(logrus.Fields{
				"URI":    values.URI,
				"status": values.Status,
			}).Info("request")
			return nil
		},
	}))

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

	// Common > Alarm
	alarm := commonModule.GetAlarmController(conn)
	alarmSns := commonModule.GetAlarmSnsController(conn)
	alarmPolicy := commonModule.GetAlarmPolicyController(conn)
	alarmStatistics := commonModule.GetAlarmStatisticsController(conn)
	alarmAction := commonModule.GetAlarmActionController(conn)
	logsearch := commonModule.GetLogSearchController(conn)

	// IaaS
	openstackModule := iaas.GetOpenstackController(conn.OpenstackProvider)
	zabbixModule := iaas.GetZabbixController(conn.ZabbixSession, conn.OpenstackProvider)

	// CP
	clusterModule := CP.GetClusterController(conn.CP)
	worklaodModule := CP.GetWorkloadController(conn.CP)
	podModule := CP.GetPodController(conn.CP)

	// SaaS
	pinpointModule := saas.GetPinpointController(conn.SaaS)
	saasModule := saas.GetSaasController(conn.SaaS)

	apBosh := AP.GetBoshController(conn)
	apPaasta := AP.GetPaastaController(conn)
	apContainer := AP.GetApContainerController(conn)

	// Router 설정
	//// Token은 항상 접근 가능하도록
	e.POST("/api/v1/token", apiToken.CreateToken) // 토큰 생성
	e.PUT("/api/v1/token", apiToken.RefreshToken) // 토큰 리프레시

	//// 그외에 다른 정보는 발급된 토큰을 기반으로 유효한 토큰을 가진 사용자만 접근하도록 middleware 설정
	//// 추가 설명 : middlewares.CheckToken 설정 (입력된 JWT 토큰 검증 및 검증된 요청자 API 접근 허용)
	//// Swagger에서는 CheckToken 프로세스에 의해 아래 function을 실행할 수 없음 (POSTMAN 이용)
	v1 := e.Group("/api/v1", middlewares.CheckToken(conn))
	v1.GET("/members", apiUser.GetMember)

	// Common > Alarm
	v1.GET("/alarm", alarm.GetAlarms)
	v1.POST("/alarm/sns", alarmSns.CreateAlarmSns)
	v1.GET("/alarm/sns", alarmSns.GetAlarmSns)
	v1.PUT("/alarm/sns", alarmSns.UpdateAlarmSns)
	v1.DELETE("/alarm/sns", alarmSns.DeleteAlarmSns)
	v1.POST("/alarm/policy", alarmPolicy.CreateAlarmPolicy)
	v1.GET("/alarm/policy", alarmPolicy.GetAlarmPolicy)
	v1.PUT("/alarm/policy", alarmPolicy.UpdateAlarmPolicy)
	v1.PUT("/alarm/target", alarmPolicy.UpdateAlarmTarget)
	v1.GET("/alarm/stats", alarmStatistics.GetAlarmStatistics)
	v1.GET("/alarm/stats/service", alarmStatistics.GetAlarmStatisticsService)
	v1.GET("/alarm/stats/resource", alarmStatistics.GetAlarmStatisticsResource)
	v1.POST("/alarm/action", alarmAction.CreateAlarmAction)
	v1.GET("/alarm/action", alarmAction.GetAlarmAction)
	v1.PATCH("/alarm/action", alarmAction.UpdateAlarmAction)
	v1.DELETE("/alarm/action", alarmAction.DeleteAlarmAction)
	v1.GET("/log/:uuid", logsearch.GetLogs)

	// AP > BOSH
	v1.GET("/ap/bosh", apBosh.GetBoshInfoList)
	v1.GET("/ap/bosh/overview", apBosh.GetBoshOverview)
	v1.GET("/ap/bosh/summary", apBosh.GetBoshSummary)
	v1.GET("/ap/bosh/process", apBosh.GetBoshProcessByMemory)
	v1.GET("/ap/bosh/chart/:uuid", apBosh.GetBoshChart)

	// AP > PaaS-TA
	v1.GET("/ap/paasta", apPaasta.GetPaastaInfoList)
	v1.GET("/ap/paasta/overview", apPaasta.GetPaastaOverview)
	v1.GET("/ap/paasta/summary", apPaasta.GetPaastaSummary)
	v1.GET("/ap/paasta/process", apPaasta.GetPaastaProcessByMemory)
	v1.GET("/ap/paasta/chart/:uuid", apPaasta.GetPaastaChart)

	// AP > Container
	v1.GET("/ap/container/cell", apContainer.GetCellInfo)
	v1.GET("/ap/container/zone", apContainer.GetZoneInfo)
	v1.GET("/ap/container/app", apContainer.GetAppInfo)
	v1.GET("/ap/container/container", apContainer.GetContainerInfo)
	v1.GET("/ap/container/overview", apContainer.GetContainerPageOverview)
	v1.GET("/ap/container/cell/status", apContainer.GetCellStatus)
	v1.GET("/ap/container/container/status", apContainer.GetContainerStatus)
	v1.GET("/ap/container/container/cpu/:id/usages", apContainer.GetContainerCpuUsages)
	v1.GET("/ap/container/container/cpu/:id/loads", apContainer.GetContainerCpuLoads)
	v1.GET("/ap/container/container/memory/:id/usages", apContainer.GetContainerMemoryUsages)
	v1.GET("/ap/container/container/disk/:id/usages", apContainer.GetContainerDiskUsages)
	v1.GET("/ap/container/container/network/:id/bytes", apContainer.GetContainerNetworkBytes)
	v1.GET("/ap/container/container/network/:id/drops", apContainer.GetContainerNetworkDrops)
	v1.GET("/ap/container/container/network/:id/errors", apContainer.GetContainerNetworkErrors)

	// IaaS
	v1.GET("/iaas/hypervisor/stats", openstackModule.GetHypervisorStatistics)
	v1.GET("/iaas/hypervisor/list", openstackModule.GetHypervisorList)
	v1.GET("/iaas/project/list", openstackModule.GetProjectList)
	v1.GET("/iaas/instance/usage/list", openstackModule.GetProjectUsage)
	v1.GET("/iaas/instance/cpu/usage", zabbixModule.GetCpuUsage)
	v1.GET("/iaas/instance/memory/usage", zabbixModule.GetMemoryUsage)
	v1.GET("/iaas/instance/disk/usage", zabbixModule.GetDiskUsage)
	v1.GET("/iaas/instance/cpu/load/average", zabbixModule.GetCpuLoadAverage)
	v1.GET("/iaas/instance/disk/io/rate", zabbixModule.GetDiskIORate)
	v1.GET("/iaas/instance/network/io/bytes", zabbixModule.GetNetworkIOBytes)

	// CP
	v1.GET("/cp/cluster/average/:type", clusterModule.GetClusterAverage)
	v1.GET("/cp/cluster/worknodes", clusterModule.GetWorkNodeList)
	v1.GET("/cp/cluster/worknode", clusterModule.GetWorkNode)
	v1.GET("/cp/workload/status", worklaodModule.GetWorkloadStatus)
	v1.GET("/cp/workload/list", worklaodModule.GetWorkloadList)
	v1.GET("/cp/workload/metrics", worklaodModule.GetWorkloadDetailMetrics)
	v1.GET("/cp/workload/container/list", worklaodModule.GetWorkloadContainerList)
	v1.GET("/cp/workload/container/metrics", worklaodModule.GetContainerMetrics)
	v1.GET("/cp/workload/container/log", worklaodModule.GetContainerLog)
	v1.GET("/cp/pod/status", podModule.GetPodStatus)
	v1.GET("/cp/pod/list", podModule.GetPodList)
	v1.GET("/cp/pod/metrics", podModule.GetPodDetailMetrics)
	v1.GET("/cp/pod/container/list", podModule.GetPodContainerList)

	// SaaS
	v1.GET("/saas/pinpoint/getAgentList", pinpointModule.GetAgentList)
	v1.GET("/saas/pinpoint/:chartType/getAgentStat", pinpointModule.GetAgentStat)

	v1.GET("/saas/app/status", saasModule.GetApplicationStatus)
	v1.GET("/saas/app/usage", saasModule.GetApplicationUsage)
	v1.GET("/saas/app/usage/list", saasModule.GetApplicationUsageList)

	return e
}

func ApiLogger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {

			now := time.Now()
			err := next(ctx)
			if err != nil {
				ctx.Error(err)
			}

			requestId := ctx.Request().Header.Get(echo.HeaderXRequestID)
			if requestId == "" {
				ctx.Response().Header().Get(echo.HeaderXRequestID)
			}
			fields := []zapcore.Field{
				zap.Int("status", ctx.Response().Status),
				zap.String("latency", time.Since(now).String()),
				zap.String("id", requestId),
				zap.String("method", ctx.Request().Method),
				zap.String("uri", ctx.Request().RequestURI),
				zap.String("host", ctx.Request().Host),
				zap.String("remote_ip", ctx.RealIP()),
			}

			n := ctx.Response().Status
			switch {
			case n >= 500:
				logger.Error("Server error", fields...)
			case n >= 400:
				logger.Warn("Client error", fields...)
			case n >= 300:
				logger.Info("Redirection", fields...)
			default:
				logger.Info("Success", fields...)
			}

			return nil

		}

	}
}
