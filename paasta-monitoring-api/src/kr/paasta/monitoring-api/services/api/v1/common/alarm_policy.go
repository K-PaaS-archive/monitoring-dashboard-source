package common

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"paasta-monitoring-api/dao/api/v1/common"
	"paasta-monitoring-api/helpers"
	models "paasta-monitoring-api/models/api/v1"
	"time"
)

type AlarmPolicyService struct {
	DbInfo *gorm.DB
}

func GetAlarmPolicyService(DbInfo *gorm.DB) *AlarmPolicyService {
	return &AlarmPolicyService{
		DbInfo: DbInfo,
	}
}

func (service *AlarmPolicyService) CreateAlarmPolicy(ctx echo.Context) (string, error) {
	logger := ctx.Request().Context().Value("LOG").(*logrus.Entry)

	var params []models.AlarmPolicies
	errValid := helpers.BindJsonAndCheckValid(ctx, &params)
	if errValid != nil {
		logger.Error(errValid)
		return "", errors.New("Invalid JSON provided, please check the request JSON")
	}

	for i, _ := range params {
		params[i].RegUser = ctx.Get("userId").(string)
		params[i].RegDate = time.Now()
	}

	err := common.GetAlarmPolicyDao(service.DbInfo).CreateAlarmPolicy(params)
	if err != nil {
		return "FAILED CREATE ALARM POLICY!", err
	}
	return "SUCCEEDED CREATE ALARM POLICY!", nil
}

func (service *AlarmPolicyService) GetAlarmPolicy(ctx echo.Context) ([]models.AlarmPolicies, error) {
	logger := ctx.Request().Context().Value("LOG").(*logrus.Entry)
	params := models.AlarmPolicies{
		OriginType: ctx.QueryParam("originType"),
		AlarmType:  ctx.QueryParam("alarmType"),
	}

	results, err := common.GetAlarmPolicyDao(service.DbInfo).GetAlarmPolicy(params)
	if err != nil {
		logger.Error(err)
		return results, err
	}
	return results, nil
}

func (service *AlarmPolicyService) UpdateAlarmPolicy(ctx echo.Context) (string, error) {
	logger := ctx.Request().Context().Value("LOG").(*logrus.Entry)

	var params []models.AlarmPolicyRequest
	err := helpers.BindJsonAndCheckValid(ctx, &params)
	if err != nil {
		logger.Error(err)
		return "", errors.New(models.ERR_PARAM_VALIDATION)
	}
	for _, policyParam := range params {
		param := models.AlarmPolicies{
			OriginType:        policyParam.OriginType,
			AlarmType:         policyParam.AlarmType,
			WarningThreshold:  policyParam.WarningThreshold,
			CriticalThreshold: policyParam.CriticalThreshold,
			RepeatTime:        policyParam.RepeatTime,
			MeasureTime:       policyParam.MeasureTime,
			ModiUser:          ctx.Get("userId").(string),
			ModiDate:          time.Now(),
		}
		err := common.GetAlarmPolicyDao(service.DbInfo).UpdateAlarmPolicy(param)
		if err != nil {
			logger.Error(err)
			return "FAILED UPDATE POLICY!", err
		}
	}
	return "SUCCEEDED UPDATE POLICY!", nil
}

func (service *AlarmPolicyService) UpdateAlarmTarget(ctx echo.Context) (string, error) {
	logger := ctx.Request().Context().Value("LOG").(*logrus.Entry)

	var params []models.AlarmTargetRequest
	err := helpers.BindJsonAndCheckValid(ctx, &params)
	if err != nil {
		logger.Error(err)
		return "", errors.New("Invalid JSON provided, please check the request JSON")
	}

	for _, param := range params {
		err := common.GetAlarmPolicyDao(service.DbInfo).UpdateAlarmTarget(param)
		if err != nil {
			logger.Error(err)
			return "FAILED UPDATE TARGET!", err
		}
	}
	return "SUCCEEDED UPDATE TARGET!", nil
}
