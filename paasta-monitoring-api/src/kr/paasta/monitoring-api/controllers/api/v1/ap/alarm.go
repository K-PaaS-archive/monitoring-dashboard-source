package ap

import (
	"GoEchoProject/apiHelpers"
	"GoEchoProject/connections"
	"GoEchoProject/helpers"
	models "GoEchoProject/models/api/v1"
	AP "GoEchoProject/services/api/v1/ap"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ApAlarmController struct {
	DbInfo *gorm.DB
}

func GetApAlarmController(conn connections.Connections) *ApAlarmController {
	return &ApAlarmController{
		DbInfo: conn.DbInfo,
	}
}

// GetAlarmStatus
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      전체 알람 현황 가져오기
//  @Description  전체 알람 현황을 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=v1.Alarms}
//  @Router       /api/v1/ap/alarm/status [get]
func (ap *ApAlarmController) GetAlarmStatus(c echo.Context) error {
	results, err := AP.GetApAlarmService(ap.DbInfo).GetAlarmStatus()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get alarm status.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get alarms status.", results)
	return nil
}

// GetAlarmPolicy
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      전체 알람 정책 가져오기
//  @Description  전체 알람 정책을 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=v1.AlarmPolicies}
//  @Router       /api/v1/ap/alarm/policy [get]
func (ap *ApAlarmController) GetAlarmPolicy(c echo.Context) error {
	results, err := AP.GetApAlarmService(ap.DbInfo).GetAlarmPolicy()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get alarm policy.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get alarms policy.", results)
	return nil
}

// UpdateAlarmPolicy
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      전체 알람 정책 업데이트하기
//  @Description  전체 알람 정책을 업데이트 한다.
//  @Accept       json
//  @Produce      json
//  @Param        AlarmPolicyRequest  body      v1.AlarmPolicyRequest  true  "알람 정책을 변경하기 위한 정보를 주입한다."
//  @Success      200                {object}  apiHelpers.BasicResponseForm
//  @Router       /api/v1/ap/alarm/policy [put]
func (ap *ApAlarmController) UpdateAlarmPolicy(c echo.Context) error {
	var request []models.AlarmPolicyRequest
	err := helpers.BindRequestAndCheckValid(c, &request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Invalid JSON provided, please check the request JSON", err.Error())
		return err
	}

	results, err := AP.GetApAlarmService(ap.DbInfo).UpdateAlarmPolicy(request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to update alarm policy.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to update alarms policy.", results)
	return nil
}

// UpdateAlarmTarget
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      전체 알람 타겟 업데이트하기
//  @Description  전체 알람 타겟을 업데이트 한다.
//  @Accept       json
//  @Produce      json
//  @Param        AlarmTargetRequest  body      v1.AlarmPolicyRequest  true  "알람 타겟을 변경하기 위한 정보를 주입한다."
//  @Success      200                {object}  apiHelpers.BasicResponseForm
//  @Router       /api/v1/ap/alarm/target [put]
func (ap *ApAlarmController) UpdateAlarmTarget(c echo.Context) error {
	var request []models.AlarmTargetRequest
	err := helpers.BindRequestAndCheckValid(c, &request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Invalid JSON provided, please check the request JSON", err.Error())
		return err
	}

	results, err := AP.GetApAlarmService(ap.DbInfo).UpdateAlarmTarget(request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to update alarm target.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to update alarms target.", results)
	return nil
}

// RegisterSnsAccount
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      알람 받을 SNS 계정 등록하기
//  @Description  알람 받을 SNS 계정을 등록한다.
//  @Accept       json
//  @Produce      json
//  @Param        SnsAccountRequest  body      v1.SnsAccountRequest  true  "알람 받을 SNS 계정 정보를 주입한다."
//  @Success      200                {object}  apiHelpers.BasicResponseForm
//  @Router       /api/v1/ap/alarm/sns [post]
func (ap *ApAlarmController) RegisterSnsAccount(c echo.Context) error {
	var request models.SnsAccountRequest
	err := helpers.BindRequestAndCheckValid(c, &request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Invalid JSON provided, please check the request JSON", err.Error())
		return err
	}

	results, err := AP.GetApAlarmService(ap.DbInfo).RegisterSnsAccount(request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to register sns account.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to register sns account.", results)
	return nil
}

// GetSnsAccount
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      알람 받는 SNS 계정 가져오기
//  @Description  알람 받는 SNS 계정 정보를 가져온다.
//  @Accept       json
//  @Produce      json
//  @Success      200  {object}  apiHelpers.BasicResponseForm{responseInfo=v1.AlarmSns}
//  @Router       /api/v1/ap/alarm/sns [post]
func (ap *ApAlarmController) GetSnsAccount(c echo.Context) error {
	results, err := AP.GetApAlarmService(ap.DbInfo).GetSnsAccount()
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to get sns alarm list.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to get sns alarm list.", results)
	return nil
}

// DeleteSnsAccount
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      알람 받는 SNS 계정 삭제하기
//  @Description  알람 받는 SNS 계정을 삭제한다.
//  @Accept       json
//  @Produce      json
//  @Param        SnsAccountRequest  body      v1.SnsAccountRequest  true  "삭제할 SNS 계정을 정보(ChannelId)를  주입한다."
//  @Success      200                 {object}  apiHelpers.BasicResponseForm
//  @Router       /api/v1/ap/alarm/sns [post]
func (ap *ApAlarmController) DeleteSnsAccount(c echo.Context) error {
	var request models.SnsAccountRequest
	err := helpers.BindRequestAndCheckValid(c, &request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Invalid JSON provided, please check the request JSON", err.Error())
		return err
	}

	results, err := AP.GetApAlarmService(ap.DbInfo).DeleteSnsAccount(request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to delete sns account.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to delete sns account.", results)
	return nil
}

// UpdateSnsAccount
//  * Annotations for Swagger *
//  @tags         AP
//  @Summary      알람 받을 SNS 계정 수정하기
//  @Description  알람 받을 SNS 계정 정보를 수정한다.
//  @Accept       json
//  @Produce      json
//  @Param        SnsAccountRequest  body      v1.SnsAccountRequest  true  "수정할 SNS 계정 정보를 주입한다."
//  @Success      200                 {object}  apiHelpers.BasicResponseForm
//  @Router       /api/v1/ap/alarm/sns [post]
func (ap *ApAlarmController) UpdateSnsAccount(c echo.Context) error {
	var request models.SnsAccountRequest
	err := helpers.BindRequestAndCheckValid(c, &request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Invalid JSON provided, please check the request JSON", err.Error())
		return err
	}

	results, err := AP.GetApAlarmService(ap.DbInfo).UpdateSnsAccount(request)
	if err != nil {
		apiHelpers.Respond(c, http.StatusBadRequest, "Failed to update sns account.", err.Error())
		return err
	}

	apiHelpers.Respond(c, http.StatusOK, "Succeeded to update sns account.", results)
	return nil
}
