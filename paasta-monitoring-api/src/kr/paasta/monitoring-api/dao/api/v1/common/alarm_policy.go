package common

import (
	"fmt"
	"gorm.io/gorm"
	models "paasta-monitoring-api/models/api/v1"
	"time"
)

type AlarmPolicyDao struct {
	DbInfo *gorm.DB
}

func GetAlarmPolicyDao(DbInfo *gorm.DB) *AlarmPolicyDao {
	return &AlarmPolicyDao{
		DbInfo: DbInfo,
	}
}

func (dao *AlarmPolicyDao) GetAlarmStatus() ([]models.Alarms, error) {
	var response []models.Alarms
	results := dao.DbInfo.Debug().Table("alarms").
		Select("*").
		Find(&response)

	if results.Error != nil {
		fmt.Println(results.Error)
		return response, results.Error
	}

	return response, nil
}

func (dao *AlarmPolicyDao) CreateAlarmPolicy(params []models.AlarmPolicies) error {
	results := dao.DbInfo.Debug().CreateInBatches(&params, 100)
	if results.Error != nil {
		return results.Error
	}

	return nil
}

func (dao *AlarmPolicyDao) GetAlarmPolicy(params models.AlarmPolicies) ([]models.AlarmPolicies, error) {
	var response []models.AlarmPolicies
	results := dao.DbInfo.Debug().Where(params).Find(&response)
	if results.Error != nil {
		fmt.Println(results.Error)
		return response, results.Error
	}

	return response, nil
}

func (dao *AlarmPolicyDao) UpdateAlarmPolicy(param models.AlarmPolicies) error {
	results := dao.DbInfo.Debug().Model(&param).
		Where("origin_type = ? AND alarm_type = ?", param.OriginType, param.AlarmType).
		Updates(&param)

	if results.Error != nil {
		fmt.Println(results.Error)
		return results.Error
	}

	return nil
}

func (dao *AlarmPolicyDao) UpdateAlarmTarget(request models.AlarmTargetRequest) error {
	results := dao.DbInfo.Debug().Table("alarm_targets").
		Where("origin_type = ?", request.OriginType).
		Updates(map[string]interface{}{
			"mail_address": request.MailAddress,
			"mail_send_yn": request.MailSendYN,
			"modi_date":    time.Now().UTC().Add(time.Hour * 9),
			"modi_user":    "admin"})

	if results.Error != nil {
		fmt.Println(results.Error)
		return results.Error
	}

	return nil
}
