package dao

import (
	client "github.com/influxdata/influxdb/client/v2"
	cb "kr/paasta/monitoring-batch/model/base"
	mod "kr/paasta/monitoring-batch/model"
	"kr/paasta/monitoring-batch/util"
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
)

type boshAlarmStruct struct {
	influxClient 	client.Client
}


func GetBoshAlarmDao(influxClient client.Client) *boshAlarmStruct{

	return &boshAlarmStruct{
		influxClient: 	influxClient,
	}
}


func (b boshAlarmStruct) GetBoshAlarmPolicy(txn *gorm.DB) ([]mod.AlarmPolicy, cb.ErrMessage) {

	var alarmPolicy []mod.AlarmPolicy

	status := txn.Debug().Model(&alarmPolicy).Where("origin_type = ? ", cb.ORIGIN_TYPE_BOSH).Find(&alarmPolicy)

	err := util.GetError().DbCheckError(status.Error)
	if err != nil{
		fmt.Println("Error::", err )
		return   nil, err
	}

	return alarmPolicy, nil
}

func (b boshAlarmStruct) GetBoshCpuUsage(request mod.BoshReq) (_ client.Response, errMsg cb.ErrMessage)  {

	var errLogMsg string
	defer func() {
		if r := recover(); r != nil {
			errMsg = cb.ErrMessage{
				"Message": errLogMsg ,
			}
		}
	}()
	// alarm measure time default : 120s
	measureTime := "120s"

	for _,value := range request.MeasureTimeList{
		if value.Item == cb.ALARM_TYPE_CPU{
			measureTime = strconv.Itoa(value.MeasureTime) + "s"
		}
	}

	//deployment = 'bosh'
	//cpuUsageSql := "select mean(value) as usage  from bosh_metrics where origin = '%s' and metricname =~ /cpuStats.core*/ and time > now() - 2m group by time(1m) order by time desc limit 1"
	//cpuUsageSql := "select mean(value) as usage  from bosh_metrics where deployment = '%s' and metricname =~ /cpuStats.core*/ and time > now() - %s"
	cpuUsageSql := "select mean(value) as usage  from bosh_metrics where metricname =~ /cpuStats.core*/ and time > now() - %s"
	var q client.Query

	q = client.Query{
		//Command:  fmt.Sprintf( cpuUsageSql , request.ServiceName, measureTime),
		Command:  fmt.Sprintf( cpuUsageSql , measureTime),
		Database: request.MetricDatabase,
	}
	fmt.Println("CPU Sql======>", q)

	resp, err := b.influxClient.Query(q)

	if err != nil{
		errLogMsg = err.Error()
	}
	return util.GetError().CheckError(*resp, err)
}

func (b boshAlarmStruct) GetBoshMemoryUsage(request mod.BoshReq) (_ client.Response, errMsg cb.ErrMessage)  {

	var errLogMsg string
	defer func() {
		if r := recover(); r != nil {
			errMsg = cb.ErrMessage{
				"Message": errLogMsg ,
			}
		}
	}()
	// alarm measure time default : 120s
	measureTime := "120s"

	for _,value := range request.MeasureTimeList{
		if value.Item == cb.ALARM_TYPE_MEMORY{
			measureTime = strconv.Itoa(value.MeasureTime) + "s"
		}
	}
        
// 	memoryTotalSql := "select mean(value) as usage from bosh_metrics where deployment = '%s' and metricname = 'memoryStats.TotalMemory' and time > now() - %s; "
// 	memoryFreeSql := "select mean(value) as memUsage from bosh_metrics where deployment = '%s' and metricname = 'memoryStats.FreeMemory' and time > now() - %s "
	memoryTotalSql := "select mean(value) as usage from bosh_metrics where  metricname = 'memoryStats.TotalMemory' and time > now() - %s ;"
	memoryFreeSql := "select mean(value) as memUsage from bosh_metrics where  metricname = 'memoryStats.FreeMemory' and time > now() - %s ;"
	var q client.Query

	q = client.Query{
//		Command:  fmt.Sprintf( memoryTotalSql + memoryFreeSql , "bosh" , measureTime, "bosh" , measureTime),
		Command:  fmt.Sprintf( memoryTotalSql + memoryFreeSql , measureTime,  measureTime),
		Database: request.MetricDatabase,
	}
        fmt.Println("333333333333333333333", q )
	resp, err := b.influxClient.Query(q)
        fmt.Println("4444444444444444444:", err)
	fmt.Println("Memory Sql==>", q)
        fmt.Println("usage===>", resp)
	if err != nil{
		errLogMsg = err.Error()
	}
	return util.GetError().CheckError(*resp, err)
}

func (b boshAlarmStruct) GetBoshDiskUsage(request mod.BoshReq) (_ client.Response, errMsg cb.ErrMessage)  {

	var errLogMsg string
	defer func() {
		if r := recover(); r != nil {
			errMsg = cb.ErrMessage{
				"Message": errLogMsg ,
			}
		}
	}()
	// alarm measure time default : 120s
	measureTime := "120s"

	for _,value := range request.MeasureTimeList{
		if value.Item == cb.ALARM_TYPE_DISK{
			measureTime = strconv.Itoa(value.MeasureTime) + "s"
		}
	}

	//memoryUsageSql := "select mean(value) as usage from bosh_metrics where deployment = '%s' and metricname = 'diskStats./var/vcap/data.Usage' and time > now() - %s"
	memoryUsageSql := "select mean(value) as usage from bosh_metrics where metricname = 'diskStats./var/vcap/data.Usage' and time > now() - %s"
	var q client.Query

	q = client.Query{
		//Command:  fmt.Sprintf( memoryUsageSql , request.ServiceName, measureTime),
		Command:  fmt.Sprintf( memoryUsageSql , measureTime),
		Database: request.MetricDatabase,
	}

	resp, err := b.influxClient.Query(q)
	fmt.Println("Disk Sql==>%s", q)
	if err != nil{
		errLogMsg = err.Error()
	}
	return util.GetError().CheckError(*resp, err)
}

func (b boshAlarmStruct) GetBoshRootDiskUsage(request mod.BoshReq) (_ client.Response, errMsg cb.ErrMessage)  {

	var errLogMsg string
	defer func() {
		if r := recover(); r != nil {
			errMsg = cb.ErrMessage{
				"Message": errLogMsg ,
			}
		}
	}()
	// alarm measure time default : 120s
	measureTime := "120s"

	for _,value := range request.MeasureTimeList{
		if value.Item == cb.ALARM_TYPE_DISK{
			measureTime = strconv.Itoa(value.MeasureTime) + "s"
		}
	}

	//memoryUsageSql := "select mean(value) as usage from bosh_metrics where deployment = '%s' and metricname = 'diskStats./.Used' and time > now() - %s"
	memoryUsageSql := "select mean(value) as usage from bosh_metrics where metricname = 'diskStats./.Used' and time > now() - %s"

	var q client.Query

	q = client.Query{
		//Command:  fmt.Sprintf( memoryUsageSql , request.ServiceName, measureTime ),
		Command:  fmt.Sprintf( memoryUsageSql ,  measureTime ),
		Database: request.MetricDatabase,
	}

	resp, err := b.influxClient.Query(q)
	fmt.Println("Disk Sql==>%s", q)
	if err != nil{
		errLogMsg = err.Error()
	}
	return util.GetError().CheckError(*resp, err)
}
