package common

import (
	"fmt"
	"gorm.io/gorm"
	models "paasta-monitoring-api/models/api/v1"
)

type AlarmActionDao struct {
	DbInfo *gorm.DB
}

func GetAlarmActionDao(DbInfo *gorm.DB) *AlarmActionDao {
	return &AlarmActionDao{
		DbInfo: DbInfo,
	}
}

func (dao *AlarmActionDao) CreateAlarmAction(params models.AlarmActions) error {
	results := dao.DbInfo.Debug().Create(&params)

	if results.Error != nil {
		fmt.Println(results.Error)
		return results.Error
	}

	return nil
}

func (dao *AlarmActionDao) GetAlarmAction(params models.AlarmActions) ([]models.AlarmActions, error) {
	var response []models.AlarmActions
	results := dao.DbInfo.Debug().Where(&params).Find(&response)

	if results.Error != nil {
		fmt.Println(results.Error)
		return response, results.Error
	}

	return response, nil
}

func (dao *AlarmActionDao) UpdateAlarmAction(params models.AlarmActions) *models.ApiError {
	results := dao.DbInfo.Debug().Model(&params).Where("id = ?", params.Id).Updates(&params)

	if results.Error != nil {
		return &models.ApiError{
			OriginError: results.Error,
			Message:     results.Error.Error(),
		}
	}

	return nil
}

func (dao *AlarmActionDao) DeleteAlarmAction(params models.AlarmActionRequest) error {
	results := dao.DbInfo.Debug().Table("alarm_actions").Where("id = ?", params.Id).Delete(&params)

	if results.Error != nil {
		fmt.Println(results.Error)
		return results.Error
	}

	return nil
}
