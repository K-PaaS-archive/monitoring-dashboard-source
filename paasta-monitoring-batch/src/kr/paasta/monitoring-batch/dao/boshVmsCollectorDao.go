package dao

import (
	"fmt"
	"sync"
	"github.com/jinzhu/gorm"
	"github.com/cloudfoundry-community/gogobosh"

	"monitoring-batch/util"
	mod "monitoring-batch/model"
	cb "monitoring-batch/model/base"
)

type boshStruct struct {
	boshClient 	*gogobosh.Client
}


func GetBoshVmsDao(boshClient  *gogobosh.Client) *boshStruct{

	return &boshStruct{
		boshClient: 	boshClient,
	}
}

func (b boshStruct) GetDeploymets() ([]mod.BoshDeployments, []error){

	var errs []error
	deployments, err := b.boshClient.GetDeployments()

	if err != nil {
		fmt.Println("##### bosh_vms.go - Deployments - Get Deployments error :", err)
		errs = append(errs, err)
		return nil, errs
	}

	var returnValue []mod.BoshDeployments
	var resultValue []mod.BoshDeployments

	var wg sync.WaitGroup

	wg.Add(len(deployments))

	for _, dep := range deployments{
		go func(wg *sync.WaitGroup, info gogobosh.Deployment){

			boshdeployment := mod.BoshDeployments{}

			vms, err := b.boshClient.GetDeploymentVMs(info.Name)
			if err != nil {
				errs = append(errs, err)
			}

			boshdeployment.Name = info.Name
			boshdeployment.VMS = vms

			returnValue = append(returnValue, boshdeployment)
			wg.Done()
		}(&wg, dep)
	}
	wg.Wait()

	if len(errs) > 0 {
		return nil, errs
	}

	for _, dep := range deployments{
		for _, deployment := range returnValue{

			if dep.Name ==  deployment.Name {
				resultValue = append(resultValue, deployment)
			}
		}
	}
	return resultValue, nil
}

func (f boshStruct) CreateZoneData(zoneName string, txn *gorm.DB) (cb.ErrMessage){

	zoneData := mod.Zone{Name: zoneName}

	status :=  txn.Debug().Create(&zoneData)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("CreateZoneData::", err )
	}
	return err
}


func (f boshStruct) GetZoneInfosByZoneNames(zoneNames []string, txn *gorm.DB) ([]mod.Zone, cb.ErrMessage){

	var zoneInfos []mod.Zone

	status := txn.Debug().Where("name in (?)", zoneNames).Find(&zoneInfos)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
	}
	return zoneInfos, err
}

func (f boshStruct) GetZoneInfosByZoneName(zoneName string, txn *gorm.DB) (mod.Zone, cb.ErrMessage){

	var zoneInfo mod.Zone

	status := txn.Debug().Where("name = ? ", zoneName).Find(&zoneInfo)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
	}
	return zoneInfo, err
}

func (f boshStruct) GetZoneInfosByZoneId(id int, txn *gorm.DB) (mod.Zone, cb.ErrMessage){

	var zoneInfo mod.Zone

	status := txn.Debug().Where("id = ? ", id).Find(&zoneInfo)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
	}
	return zoneInfo, err
}

func (f boshStruct) IsExistJobName(jobName string, txn *gorm.DB) (bool, cb.ErrMessage){

	var vmInfo mod.Vm
	var cnt int
	status := txn.Debug().Where("name = ?", jobName).Find(&vmInfo).Count(&cnt)

	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error====::", err )
	}
	if cnt > 0 {
		return true, err
	}else{
		return false, err
	}
}

func (f boshStruct) GetJobInfo(jobName string, txn *gorm.DB) (mod.Vm, cb.ErrMessage){

	var vmInfo mod.Vm
	status := txn.Debug().Where("name = ?", jobName).Find(&vmInfo)

	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error====::", err )
	}
	return vmInfo, err
}

func (f boshStruct) UpdateVmData(vmInfo mod.Vm, txn *gorm.DB) (cb.ErrMessage){

	status :=  txn.Debug().Model(&vmInfo).Update(map[string]interface{}{"ip" : vmInfo.Ip ,"modi_date": vmInfo.ModiDate,"modi_user": vmInfo.ModiUser})
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("UPdate Vm Data Error::", err )
	}
	return err
}

func (f boshStruct) UpdateVmZoneData(vmInfo mod.Vm, txn *gorm.DB) (cb.ErrMessage){

	status :=  txn.Debug().Table("vms").Where("id = ?", vmInfo.Id).Update(map[string]interface{}{"zone_id" : vmInfo.ZoneId ,"modi_date": vmInfo.ModiDate,"modi_user": vmInfo.ModiUser})
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("UPdate Vm Data Error::", err )
	}
	return err
}

func (f boshStruct) CreateVmData(vmInfo mod.Vm, txn *gorm.DB) (cb.ErrMessage){

	status :=  txn.Debug().Create(&vmInfo)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("CreateZoneData::", err )
	}
	return err
}

func (f boshStruct) GetJobInfoList(txn *gorm.DB) ([]mod.Vm, cb.ErrMessage){

	var vmInfos []mod.Vm
	status := txn.Debug().Find(&vmInfos)

	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error====::", err )
	}
	return vmInfos, err
}

func (f boshStruct) GetZoneInfosList(txn *gorm.DB) ([]mod.Zone, cb.ErrMessage){

	var zoneInfos []mod.Zone

	status := txn.Debug().Find(&zoneInfos)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
	}
	return zoneInfos, err
}

func (f boshStruct) DeleteZoneInfo(zone mod.Zone,txn *gorm.DB) (cb.ErrMessage){

	status := txn.Debug().Model(&zone).Where("id = ?", zone.Id).Delete(&zone)
	err := util.GetError().DbCheckError(status.Error)

	return  err
}

func (f boshStruct) DeleteVmInfo(vm mod.Vm,txn *gorm.DB) (cb.ErrMessage){

	status := txn.Debug().Model(&vm).Where("id = ?", vm.Id).Delete(&vm)
	err := util.GetError().DbCheckError(status.Error)

	return  err
}


/*
func (f boshStruct) GetZoneInfos( txn *gorm.DB)  cb.ErrMessage {

	var zoneInfos []mod.Zone

	status := txn.Debug().Find(&zoneInfos)
	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
	}
	return zoneInfos, err
}*/
