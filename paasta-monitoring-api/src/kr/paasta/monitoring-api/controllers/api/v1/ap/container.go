package ap

import (
	"github.com/cloudfoundry-community/go-cfclient"
	influxDb "github.com/influxdata/influxdb1-client/v2"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"paasta-monitoring-api/apiHelpers"
	"paasta-monitoring-api/connections"
	models "paasta-monitoring-api/models/api/v1"
	AP "paasta-monitoring-api/services/api/v1/ap"
)

type ApContainerController struct {
	DbInfo             *gorm.DB
	InfluxDbInfo       influxDb.Client
	Databases          models.Databases
	CloudFoundryClient *cfclient.Client
}

func GetApContainerController(conn connections.Connections) *ApContainerController {
	return &ApContainerController{
		DbInfo:             conn.DbInfo,
		InfluxDbInfo:       conn.InfluxDBClient,
		Databases:          conn.Databases,
		CloudFoundryClient: conn.CloudFoundryClient,
	}
}

// GetCellInfo
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      Cell 정보 가져오기
//  @Description  Cell 정보를 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=v1.CellInfo}
//  @Router       /api/v1/ap/container/cell [get]
func (ap *ApContainerController) GetCellInfo(c echo.Context) error {
	results, err := AP.GetApContainerService(ap.DbInfo, ap.InfluxDbInfo, ap.Databases, ap.CloudFoundryClient).GetCellInfo()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get cell info.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get cell info.", results)
	return nil
}

// GetZoneInfo
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      Zone 정보 가져오기
//  @Description  Zone 정보를 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=v1.ZoneInfo}
//  @Router       /api/v1/ap/container/zone [get]
func (ap *ApContainerController) GetZoneInfo(c echo.Context) error {
	results, err := AP.GetApContainerService(ap.DbInfo, ap.InfluxDbInfo, ap.Databases, ap.CloudFoundryClient).GetZoneInfo()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get zone info.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get zone info.", results)
	return nil
}

// GetAppInfo
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      App 정보 가져오기
//  @Description  App 정보를 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=cfclient.App}
//  @Router       /api/v1/ap/container/zone [get]
func (ap *ApContainerController) GetAppInfo(c echo.Context) error {
	results, err := AP.GetApContainerService(ap.DbInfo, ap.InfluxDbInfo, ap.Databases, ap.CloudFoundryClient).GetAppInfo()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get apps info.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get apps info.", results)
	return nil
}
