package common

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"paasta-monitoring-api/apiHelpers"
	"paasta-monitoring-api/connections"
	models "paasta-monitoring-api/models/api/v1"
	service "paasta-monitoring-api/services/api/v1/common"
)

type LogSearchController struct {
	DbInfo         *gorm.DB
	InfluxDbClient models.InfluxDbClient
	BoshInfoList   []models.Bosh
}

func GetLogSearchController(conn connections.Connections) *LogSearchController {
	return &LogSearchController{
		DbInfo:         conn.DbInfo,
		InfluxDbClient: conn.InfluxDbClient,
		BoshInfoList:   conn.BoshInfoList,
	}
}

// GetLogs
//  @tags         Common
//  @Summary      로그 정보 가져오기
//  @Description  로그 정보를 가져온다.
//  @Description  특정기간(대상날짜, 시작시간, 종료시간 사용) 로그 조회와 최근기간 로그 조회 파라미터는 중복 사용이 불가능하다.
//  @Description  또한 특정기간 로그 조회 또는 최근기간 로그 조회를 위한 파라미터가 반드시 하나는 사용되어야 한다.
//  @Accept       json
//  @Produce      json
//  @Param        uuid        path      string  true   "로그 조회시 대상 VM의 UUID를 주입한다."  example(36dd3d08-5198-42b6-4130-d0c04479236f)
//  @Param        logType     query     string  true   "로그 정보를 조회하고자 하는 타입을 지정한다."  enums(bosh, cf)
//  @Param        keyword     query     string  false  "로그 조회시 특정 내용을 포함하는 키워드 검색이 필요할 경우 사용한다."
//  @Param        targetDate  query     string  false  "로그 정보를 조회하고자 하는 대상 날짜를 주입한다."            example(2022-07-28)
//  @Param        startTime   query     string  false  "로그 정보를 조회하고자 하는 시작 시간를 주입한다."            example(09:00:00)
//  @Param        endTime     query     string  false  "로그 정보를 조회하고자 하는 종료 시간를 주입한다."            example(18:00:00)
//  @Param        period      query     string  false  "로그 정보 조회시 현재를 기준으로 특정 기간 동안의 내용을 조회한다."  example(10s)
//  @Success      200         {object}  apiHelpers.BasicResponseForm{responseInfo=v1.Logs}
//  @Router       /api/v1/log/{uuid} [get]
func (l *LogSearchController) GetLogs(ctx echo.Context) error {
	results, err := service.GetLogSearchService(l.DbInfo, l.InfluxDbClient, l.BoshInfoList).GetLogs(ctx)
	if err != nil {
		apiHelpers.Respond(ctx, http.StatusBadRequest, "Failed to get logs.", err.Error())
		return err
	}

	apiHelpers.Respond(ctx, http.StatusOK, "Succeeded to get logs.", results)
	return nil
}
