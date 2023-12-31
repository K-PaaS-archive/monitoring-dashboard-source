package main

import (
	"crypto/tls"
	"fmt"
	"github.com/cloudfoundry-community/go-cfclient"
	commonModel "monitoring-portal/common/model"
	"monitoring-portal/zabbix-client/lib/go-zabbix"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cihub/seelog"
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/jinzhu/gorm"

	"monitoring-portal/common/config"
	"monitoring-portal/handlers"
	"monitoring-portal/iaas_new"
	"monitoring-portal/iaas_new/model"
	bm "monitoring-portal/paas/model"
	"monitoring-portal/utils"
)

type Config map[string]string

type MemberInfo struct {
	UserId        string    `gorm:"type:varchar(50);primary_key"`
	UserPw        string    `gorm:"type:varchar(500);null;"`
	UserEmail     string    `gorm:"type:varchar(100);null;"`
	UserNm        string    `gorm:"type:varchar(100);null;"`
	IaasUserId    string    `gorm:"type:varchar(100);null;"`
	IaasUserPw    string    `gorm:"type:varchar(100);null;"`
	CaasUserId    string    `gorm:"type:varchar(100);null;"`
	CaasUserPw    string    `gorm:"type:varchar(100);null;"`
	PaasUserId    string    `gorm:"type:varchar(100);null;"`
	PaasUserPw    string    `gorm:"type:varchar(100);null;"`
	IaasUserUseYn string    `gorm:"type:varchar(1);null;"`
	PaasUserUseYn string    `gorm:"type:varchar(1);null;"`
	CaasUserUseYn string    `gorm:"type:varchar(1);null;"`
	UpdatedAt     time.Time `gorm:"type:datetime;null;DEFAULT:null"`
	CreatedAt     time.Time `gorm:"type:datetime;null;DEFAULT:CURRENT_TIMESTAMP"`
}

var logger seelog.LoggerInterface

func main() {

	xmlFile, err := config.ConvertXmlToString("log_config.xml")
	if err != nil {
		os.Exit(-1)
	}

	logger, _ = seelog.LoggerFromConfigAsBytes([]byte(xmlFile))
	model.MonitLogger = logger
	utils.Logger = logger

	// 기본적인 프로퍼티 설정 정보 읽어오기
	configMap, err := config.ImportConfig("config.ini")
	if err != nil {
		logger.Error(err)
		os.Exit(-1)
	}

	timeGap, _ := strconv.Atoi(configMap["gmt.time.gap"])
	model.GmtTimeGap = timeGap
	bm.GmtTimeGap = timeGap

	apiPort, _ := strconv.Atoi(configMap["server.port"])

	sysType := configMap["system.monitoring.type"]

	// paas client
	var paaSInfluxServerClient client.Client
	var databases bm.Databases
	var boshClient *gogobosh.Client
	var cfClient *cfclient.Client

	// config.ini 파일에서 MySQL 접속정보를 추출
	dbConnInfo := config.InitDBConnectionConfig(configMap)

	paasConnectionString := utils.GetConnectionString(dbConnInfo.Host, dbConnInfo.Port, dbConnInfo.UserName, dbConnInfo.UserPassword, dbConnInfo.DbName)
	logger.Infof("DB Connection Info : %v\n", paasConnectionString)

	paasDbAccessObj, paasDbErr := gorm.Open(dbConnInfo.DbType, paasConnectionString+"?charset=utf8&parseTime=true")
	if paasDbErr != nil {
		logger.Errorf("%v\n", paasDbErr)
		return
	}

	// 2021.09.06 - 왜 있는건지??
	// 2021.11.11 - @Deprecated
	// memberInfo table (use common database table)
	//createTable(paasDbAccessObj)

	// Redis Client
	rdClient := redis.NewClient(&redis.Options{
		Addr:     configMap["redis.addr"],
		Password: configMap["redis.password"],
	})
	logger.Info(rdClient)

	cfConfig := bm.CFConfig{
		Host:           configMap["paas.monitoring.cf.host"],
		CaasBrokerHost: configMap["caas.monitoring.broker.host"],
	}

	// TODO IaaS Connection Info
	iaasClient := iaas_new.IaasClient{}

	var zabbixSession *zabbix.Session
	var iaasDbAccessObj *gorm.DB
	var iaaSInfluxServerClient client.Client
	var openstackProvider model.OpenstackProvider

	if strings.Contains(sysType, utils.SYS_TYPE_ALL) || strings.Contains(sysType, utils.SYS_TYPE_IAAS) {
		/*
			iaasClient, err = iaas_new.GetIaasClients(configMap)
			if err != nil {
				logger.Error(err)
				os.Exit(-1)
			}
		*/

		iaasDbAccessObj, iaaSInfluxServerClient, openstackProvider, zabbixSession, err = getIaasClent(configMap)

	}

	if strings.Contains(sysType, utils.SYS_TYPE_ALL) || strings.Contains(sysType, utils.SYS_TYPE_PAAS) {
		paaSInfluxServerClient, databases, cfClient, boshClient, err = getPaasClients(configMap)
		if err != nil {
			logger.Error(err)
			os.Exit(-1)
		}
	}

	// Route Path 정보와 처리 서비스 연결
	var handler http.Handler

	if strings.Contains(sysType, utils.SYS_TYPE_ALL) || strings.Contains(sysType, utils.SYS_TYPE_IAAS) {
		handler = handlers.NewHandler(openstackProvider, iaaSInfluxServerClient, paaSInfluxServerClient,
			iaasDbAccessObj, paasDbAccessObj, iaasClient.AuthOpts, databases,
			rdClient, sysType, boshClient, cfConfig, zabbixSession, cfClient)
		if err := http.ListenAndServe(fmt.Sprintf(":%v", apiPort), handler); err != nil {
			logger.Error(err)
		}
	} else {
		handler = handlers.NewHandler(openstackProvider, iaaSInfluxServerClient, paaSInfluxServerClient,
			iaasDbAccessObj, paasDbAccessObj, iaasClient.AuthOpts, databases,
			rdClient, sysType, boshClient, cfConfig, zabbixSession, cfClient)
		if err := http.ListenAndServe(fmt.Sprintf(":%v", apiPort), handler); err != nil {
			logger.Error(err)
		}
	}
}

// 2021.09.06 - 이거 왜 있는지??
// 2021.11.11 - @Deprecated
//func createTable(dbClient *gorm.DB) {
//	dbClient.Debug().AutoMigrate(&MemberInfo{})
//}

func getPaasClients(config map[string]string) (paaSInfluxServerClient client.Client, databases bm.Databases, cfClient *cfclient.Client, boshClient *gogobosh.Client, err error) {

	// InfluxDB
	paasUrl, _ := config["paas.metric.db.url"]
	paasuserName, _ := config["paas.metric.db.username"]
	paasPassword, _ := config["paas.metric.db.password"]

	paaSInfluxServerClient, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:               paasUrl,
		Username:           paasuserName,
		Password:           paasPassword,
		InsecureSkipVerify: true,
	})

	logger.Infof("paaSInfluxServerClient : %v\n", paaSInfluxServerClient)
	if err != nil {
		logger.Errorf("err : %v\n", err)
	}

	// PaaS Database
	boshDatabase, _ := config["paas.metric.db.name.bosh"]
	paastaDatabase, _ := config["paas.metric.db.name.paasta"]
	containerDatabase, _ := config["paas.metric.db.name.container"]
	loggingMeasurement, _ := config["paas.metric.db.name.logging"]

	databases.BoshDatabase = boshDatabase
	databases.PaastaDatabase = paastaDatabase
	databases.ContainerDatabase = containerDatabase
	databases.LoggingDatabase = loggingMeasurement

	// Cloud Foundry Client
	cfProvider := &cfclient.Config{
		ApiAddress: config["paas.cf.client.apiaddress"],
		Username:     "admin",
		Password:     "admin",
		SkipSslValidation: true,
	}
	cfClient, _ = cfclient.NewClient(cfProvider)



	// BOSH Client Config
	boshConfig := &gogobosh.Config{
		BOSHAddress:       config["bosh.client.api.address"],
		Username:          config["bosh.client.api.username"],
		Password:          config["bosh.client.api.password"],
		HttpClient:        http.DefaultClient,
		SkipSslValidation: true,
	}
	boshClient, err = gogobosh.NewClient(boshConfig)
	if err != nil {
		logger.Errorf("Failed to create connection to the bosh server. err=", err)
	}
	return
}

func getIaasClent(config map[string]string) (iaasDbAccessObj *gorm.DB, iaaSInfluxServerClient client.Client, openstackProvider model.OpenstackProvider, zabbixSession *zabbix.Session, err error) {
	// Mysql
	iaasConfigDbCon := new(commonModel.DBConfig)
	iaasConfigDbCon.DbType = config["iaas.monitoring.db.type"]
	iaasConfigDbCon.DbName = config["iaas.monitoring.db.dbname"]
	iaasConfigDbCon.UserName = config["iaas.monitoring.db.username"]
	iaasConfigDbCon.UserPassword = config["iaas.monitoring.db.password"]
	iaasConfigDbCon.Host = config["iaas.monitoring.db.host"]
	iaasConfigDbCon.Port = config["iaas.monitoring.db.port"]

	iaasConnectionString := utils.GetConnectionString(iaasConfigDbCon.Host, iaasConfigDbCon.Port, iaasConfigDbCon.UserName, iaasConfigDbCon.UserPassword, iaasConfigDbCon.DbName)
	//fmt.Println("String:", iaasConnectionString)
	iaasDbAccessObj, _ = gorm.Open(iaasConfigDbCon.DbType, iaasConnectionString+"?charset=utf8&parseTime=true")

	// 2021.09.06 - 이거 왜 있는지?
	//Alarm 처리 내역 정보 Table 생성
	//iaasDbAccessObj.Debug().AutoMigrate(&model.AlarmActionHistory{})

	// InfluxDB
	iaasUrl, _ := config["iaas.metric.db.url"]
	iaasUserName, _ := config["iaas.metric.db.username"]
	iaasPassword, _ := config["iaas.metric.db.password"]

	iaaSInfluxServerClient, _ = client.NewHTTPClient(client.HTTPConfig{
		Addr:               iaasUrl,
		Username:           iaasUserName,
		Password:           iaasPassword,
		InsecureSkipVerify: true,
	})

	// Openstack 정보
	openstackProvider.Region, _ = config["default.region"]
	openstackProvider.Username, _ = config["default.username"]
	openstackProvider.Password, _ = config["default.password"]
	openstackProvider.Domain, _ = config["default.domain"]
	openstackProvider.TenantName, _ = config["default.tenant_name"]
	openstackProvider.AdminTenantId, _ = config["default.tenant_id"]
	openstackProvider.KeystoneUrl, _ = config["keystone.url"]
	openstackProvider.IdentityEndpoint, _ = config["identity.endpoint"]
	openstackProvider.RabbitmqUser, _ = config["rabbitmq.user"]
	openstackProvider.RabbitmqPass, _ = config["rabbitmq.pass"]
	openstackProvider.RabbitmqTargetNode, _ = config["rabbitmq.target.node"]

	model.MetricDBName, _ = config["iaas.metric.db.name"]
	model.NovaUrl, _ = config["nova.target.url"]
	model.NovaVersion, _ = config["nova.target.version"]
	model.NeutronUrl, _ = config["neutron.target.url"]
	model.NeutronVersion, _ = config["neutron.target.version"]
	model.KeystoneUrl, _ = config["keystone.target.url"]
	model.KeystoneVersion, _ = config["keystone.target.version"]
	model.CinderUrl, _ = config["cinder.target.url"]
	model.CinderVersion, _ = config["cinder.target.version"]
	model.GlanceUrl, _ = config["glance.target.url"]
	model.GlanceVersion, _ = config["glance.target.version"]
	model.DefaultTenantId, _ = config["default.tenant_id"]
	model.RabbitMqIp, _ = config["rabbitmq.ip"]
	model.RabbitMqPort, _ = config["rabbitmq.port"]
	model.GMTTimeGap, _ = strconv.ParseInt(config["gmt.time.gap"], 10, 64)

	zabbixHost := config["zabbix.host"]
	zabbixAdminId := config["zabbix.admin.id"]
	zabbixAdminPw := config["zabbix.admin.pw"]
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	cache := zabbix.NewSessionFileCache().SetFilePath("./zabbix_session")
	zabbixSession, err = zabbix.CreateClient(zabbixHost).
		WithCache(cache).
		WithHTTPClient(client).
		WithCredentials(zabbixAdminId, zabbixAdminPw).Connect()
	if err != nil {
		utils.Logger.Error(err)
	}

	return
}
