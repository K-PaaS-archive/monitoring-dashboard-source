package model

import "time"

type (
	BaseField struct {
		RegDate  time.Time `gorm:"type:datetime;DEFAULT:current_timestamp;not null;"`
		RegUser  string    `gorm:"type:varchar(36);DEFAULT:'system';not null;"`
		ModiDate time.Time `gorm:"type:datetime;DEFAULT:current_timestamp;null;"`
		ModiUser string    `gorm:"type:varchar(36);DEFAULT:'system';null;"`
	}

	//서버정보
	Bosh struct {
		Id          uint   `gorm:"primary_key"`
		ServerName  string `gorm:"type:varchar(200);not null;"`
		Type        string `gorm:"type:varchar(3);not null;"`
		ServiceName string `gorm:"type:varchar(200);not null;unique_index;"`
		Ip          string `gorm:"type:varchar(15);not null;unique_index;"`
		ExternalIp  string `gorm:"type:varchar(15);null;"`
		Port        string `gorm:"type:varchar(5);not null;"`
		Comment     string `gorm:"type:varchar(5000);null;"`
		BaseField
	}

	Zone struct {
		Id   uint   `gorm:"primary_key"`
		Name string `gorm:"type:varchar(200);not null;"`
		Vm   []Vm
		BaseField
	}

	Vm struct {
		Id     uint   `gorm:"primary_key"`
		ZoneId uint   `gorm:"primary_key;type:integer;not null;"`
		Name   string `gorm:"type:varchar(200);not null;"`
		Ip     string `gorm:"type:varchar(15);not null;unique_index;"`
		VmType string `gorm:"type:varchar(3);not null;"`
		BaseField
	}

	AlarmPolicy struct {
		Id                uint   `gorm:"primary_key"`
		OriginType        string `gorm:"type:varchar(3);not null;"`
		AlarmType         string `gorm:"type:varchar(6);not null;"`
		WarningThreshold  int    `gorm:"type:integer;not null;"`
		CriticalThreshold int    `gorm:"type:integer;not null;"`
		RepeatTime        int    `gorm:"type:integer;not null;"`
		MeasureTime       int    `gorm:"type:integer;not null;"`
		Comment           string `gorm:"type:varchar(5000);"`
		BaseField
	}

	AlarmTarget struct {
		Id          uint   `gorm:"primary_key"`
		OriginType  string `gorm:"type:varchar(3);not null;"`
		MailAddress string `gorm:"type:varchar(100);not null;"`
		MailSendYn  string `gorm:"type:varchar(1);not null;DEFAULT:'Y';"`
		BaseField
	}

	//Batch에서 Data를 생성한다.
	//정해진 시간 주기로 Batch가 실행되며, 임계치를 초과한 시스템 정보를 저장한다.
	Event struct {
		Id            uint      `gorm:"primary_key"`
		OriginType    string    `gorm:"type:varchar(3);not null;"`
		OriginId      uint      `gorm:"type:int;not null;"`
		AlarmType     string    `gorm:"type:varchar(6);not null;"`
		AlarmLevel    string    `gorm:"type:varchar(8);not null;"`
		Ip            string    `gorm:"type:varchar(15);null;"`
		AppYn         string    `gorm:"type:varchar(1);not null;"`
		AppName       string    `gorm:"type:varchar(1000);"`
		AppIndex      uint      `gorm:"type:int;"`
		ContainerName string    `gorm:"type:varchar(1000);"`
		CellName      string    `gorm:"type:varchar(500);null;"`
		Message       string    `gorm:"type:text;"`
		CpuUsage      uint      `gorm:"type:integer;not null;"`
		MemoryUsage   uint      `gorm:"type:integer;not null;"`
		DiskUsage     uint      `gorm:"type:integer;not null;"`
		RegDate       time.Time `gorm:"type:datetime;DEFAULT:current_timestamp;not null;"`
		RegUser       string    `gorm:"type:varchar(36);DEFAULT:'system';not null;"`
	}

	//알람정보
	//Batch에서 Data를 생성한다.
	//Event정보를 정해진 시간 주기로 읽어, 임계치가 정해진 시간을 초과 된 경우 저장된다.
	//저장 후 Sms, Email등으로 Alarm이 관리자에게 전송된다.
	Alarm struct {
		Id            uint   `gorm:"primary_key"`
		OriginType    string `gorm:"type:varchar(3);not null;"`
		OriginId      uint   `gorm:"type:int;not null;"`
		AlarmType     string `gorm:"type:varchar(6);not null;"`
		Level         string `gorm:"type:varchar(8);not null;"`
		Ip            string `gorm:"type:varchar(15);null;"`
		AppYn         string `gorm:"type:varchar(1);not null;"`
		AppName       string `gorm:"type:varchar(500);null;"`
		AppIndex      uint   `gorm:"type:int;null;"`
		ContainerName string `gorm:"type:varchar(40);null;"`
		AlarmTitle    string `gorm:"type:varchar(5000);not null;"`
		AlarmMessage  string `gorm:"type:text;not null;"`
		ResolveStatus string `gorm:"type:varchar(1);not null;"`   //처리 여부 1: Alarm 발생, 2: 처리중, 3: 처리 완료
		AlarmCnt      int    `gorm:"type:int;not null;DEFAULT:1"` //Alarm 발생 횟수
		BaseField
		AlarmSendDate time.Time `gorm:"type:datetime;null;DEFAULT:null"`    //Alarm 전송 시간
		CompleteDate  time.Time `gorm:"type:datetime;null;DEFAULT:null"`    //Alarm 처리 완료 시간
		CompleteUser  string    `gorm:"type:varchar(36);null;DEFAULT:null"` //Alarm 처리 완료자
	}

	AlarmAction struct {
		Id              uint   `gorm:"primary_key"`
		AlarmId         uint   `gorm:"type:integer;not null;"`
		AlarmActionDesc string `gorm:"type:text;"`
		BaseField
	}

	AutoScaleConfig struct {
		No                 uint    `gorm:"primary_key"`
		Guid               string  `gorm:"type:varchar(255);not null;"`
		Org                string  `gorm:"type:varchar(255);null;"`
		Space              string  `gorm:"type:varchar(255);null;"`
		App                string  `gorm:"type:varchar(255);null;"`
		InstanceMinCnt     int     `gorm:"type:integer;null;"`
		InstanceMaxCnt     int     `gorm:"type:integer;null;"`
		CpuThresholdMinPer float32 //`gorm:"type:integer;null;"`
		CpuThresholdMaxPer float32 //`gorm:"type:integer;null;"`
		MemoryMinSize      int     `gorm:"type:integer;null;"`
		MemoryMaxSize      int     `gorm:"type:integer;null;"`
		CheckTimeSec       int     `gorm:"type:integer;not null;"`
		AutoIncreaseYn     string  `gorm:"type:char(1);not null;"`
		AutoDecreaseYn     string  `gorm:"type:char(1);not null;"`
	}

	MemberInfo struct {
		UserId        string    `gorm:"type:varchar(50);primary_key"`
		UserPw        string    `gorm:"type:varchar(500);null;"`
		UserEmail     string    `gorm:"type:varchar(100);null;"`
		UserNm        string    `gorm:"type:varchar(100);null;"`
		IaasUserId    string    `gorm:"type:varchar(100);null;"`
		IaasUserPw    string    `gorm:"type:varchar(100);null;"`
		PaasUserId    string    `gorm:"type:varchar(100);null;"`
		PaasUserPw    string    `gorm:"type:varchar(100);null;"`
		IaasUserUseYn string    `gorm:"type:varchar(1);null;"`
		PaasUserUseYn string    `gorm:"type:varchar(1);null;"`
		UpdatedAt     time.Time `gorm:"type:datetime;null;DEFAULT:null"`
		CreatedAt     time.Time `gorm:"type:datetime;null;DEFAULT:CURRENT_TIMESTAMP"`
	}

	AlarmSns struct {
		ChannelId  uint      `gorm:"type:int unsigned auto_increment;not null;primary_key"`
		OriginType string    `gorm:"type:varchar(3);not null;DEFAULT:'all';"`
		SnsType    string    `gorm:"type:varchar(50);not null;"`
		SnsId      string    `gorm:"type:varchar(50);not null;"`
		Token      string    `gorm:"type:varchar(300);null;"`
		Expl       string    `gorm:"type:varchar(100);not null;"`
		SnsSendYn  string    `gorm:"type:varchar(1);not null;DEFAULT:'Y';"`
		RegDate    time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		RegUser    string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
		ModiDate   time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		ModiUser   string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
	}

	AlarmSnsTarget struct {
		ChannelId uint      `gorm:"type:int unsigned;not null;unique_index:uix_alarm_sns_targets;"`
		TargetId  int64     `gorm:"type:bigint(20) unsigned;not null;unique_index:uix_alarm_sns_targets;"`
		RegDate   time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		RegUser   string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
		ModiDate  time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		ModiUser  string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
	}

	AppAutoScalingPolicy struct {
		AppGuid             string    `gorm:"type:varchar(50);not null;primary_key;"`
		InstanceMinCnt      uint      `gorm:"type:int unsigned;not null;"`
		InstanceMaxCnt      uint      `gorm:"type:int unsigned;not null;"`
		CpuMinThreshold     uint      `gorm:"type:int unsigned;not null;"`
		CpuMaxThreshold     uint      `gorm:"type:int unsigned;not null;"`
		MemoryMinThreshold  uint      `gorm:"type:int unsigned;not null;"`
		MemoryMaxThreshold  uint      `gorm:"type:int unsigned;not null;"`
		InstanceScalingUnit uint      `gorm:"type:int unsigned;not null;"`
		MeasureTimeSec      uint      `gorm:"type:int unsigned;not null;"`
		AutoScalingCpuYn    string    `gorm:"type:varchar(1);not null;"`
		AutoScalingMemoryYn string    `gorm:"type:varchar(1);not null;"`
		AutoScalingOutYn    string    `gorm:"type:varchar(1);not null;"`
		AutoScalingInYn     string    `gorm:"type:varchar(1);not null;"`
		RegDate             time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		RegUser             string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
		ModiDate            time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		ModiUser            string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
	}

	AppAlarmPolicy struct {
		AppGuid                 string    `gorm:"type:varchar(50);not null;primary_key;"`
		CpuWarningThreshold     uint      `gorm:"type:int unsigned;not null;"`
		CpuCriticalThreshold    uint      `gorm:"type:int unsigned;not null;"`
		MemoryWarningThreshold  uint      `gorm:"type:int unsigned;not null;"`
		MemoryCriticalThreshold uint      `gorm:"type:int unsigned;not null;"`
		MeasureTimeSec          uint      `gorm:"type:int unsigned;not null;"`
		Email                   string    `gorm:"type:varchar(100);null;DEFAULT:null;"`
		EmailSendYn             string    `gorm:"type:varchar(1);not null;"`
		AlarmUseYn              string    `gorm:"type:varchar(1);not null;"`
		RegDate                 time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		RegUser                 string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
		ModiDate                time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		ModiUser                string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
	}

	AppAlarmHistory struct {
		AlarmId            uint      `gorm:"type:int unsigned;not null;primary_key;AUTO_INCREMENT;"`
		AppGuid            string    `gorm:"type:varchar(50);not null;"`
		AppIdx             uint      `gorm:"type:int unsigned;not null;"`
		ResourceType       string    `gorm:"type:varchar(10);not null;"`
		AlarmLevel         string    `gorm:"type:varchar(8);not null;"`
		AppName            string    `gorm:"type:varchar(50);not null;"`
		CellIp             string    `gorm:"type:varchar(15);not null;"`
		ContainerId        string    `gorm:"type:varchar(50);not null;"`
		ContainerInterface string    `gorm:"type:varchar(20);not null;"`
		AlarmTitle         string    `gorm:"type:varchar(1000);not null;"`
		AlarmMessage       string    `gorm:"type:text;not null;"`
		AlarmSendDate      time.Time `gorm:"type:datetime;null;DEFAULT:null;"`
		TerminateYn        string    `gorm:"type:varchar(1);not null;"`
		RegDate            time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		RegUser            string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
		ModiDate           time.Time `gorm:"type:datetime;not null;DEFAULT:CURRENT_TIMESTAMP;"`
		ModiUser           string    `gorm:"type:varchar(36);not null;DEFAULT:'system';"`
	}

	SendTargetAppAlarmHistory struct {
		AppAlarmHistory
		Email       string
		EmailSendYn string
	}
)
