package controller

import (
	"github.com/cloudfoundry-community/gogobosh"
	"github.com/influxdata/influxdb1-client/v2"
	"github.com/jinzhu/gorm"
	"monitoring-portal/paas/model"
	"monitoring-portal/paas/service"
	"monitoring-portal/paas/util"
	"net/http"
	"strconv"
)

type PaasController struct {
	txn          *gorm.DB
	influxClient client.Client
	databases    model.Databases
	boshClient   *gogobosh.Client
}

func GetPaasController(txn *gorm.DB, influxClient client.Client, databases model.Databases, boshClent *gogobosh.Client) *PaasController {
	return &PaasController{
		txn:          txn,
		influxClient: influxClient,
		databases:    databases,
		boshClient:   boshClent,
	}
}

func (p *PaasController) GetPaasOverview(w http.ResponseWriter, r *http.Request) {

	var apiRequest model.PaasRequest

	apiRequest.Origin = r.FormValue(":origin")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.ServiceName = r.URL.Query().Get("serviceName")
	apiRequest.Index = r.URL.Query().Get("index")
	apiRequest.Addr = r.URL.Query().Get("addr")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasOverview(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	}
	util.RenderJsonResponse(result, w)
}

func (p *PaasController) GetPaasOverviewStatus(w http.ResponseWriter, r *http.Request) {

	var apiRequest model.PaasRequest

	apiRequest.Status = r.FormValue(":status")

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasOverviewStatus(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	}
	util.RenderJsonResponse(result, w)
}

func (p *PaasController) GetPaasSummary(w http.ResponseWriter, r *http.Request) {

	var apiRequest model.PaasRequest

	apiRequest.Origin = r.FormValue(":origin")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.ServiceName = r.URL.Query().Get("serviceName")
	apiRequest.Index = r.URL.Query().Get("index")
	apiRequest.Addr = r.URL.Query().Get("addr")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")

	apiRequest.PagingReq.PageIndex, _ = strconv.Atoi(r.URL.Query().Get("pageIndex"))
	apiRequest.PagingReq.PageItem, _ = strconv.Atoi(r.URL.Query().Get("pageItems"))
	apiRequest.Name = r.URL.Query().Get("name")
	apiRequest.Ip = r.URL.Query().Get("ip")
	apiRequest.Status = r.URL.Query().Get("status")

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasSummary(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	}

	util.RenderJsonResponse(result, w)
}

func (p *PaasController) GetPaasTopProcessMemory(w http.ResponseWriter, r *http.Request) {
	var apiRequest model.PaasRequest
	apiRequest.Id = r.FormValue(":id")

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasTopProcessMemory(apiRequest)

	if err != nil {
		util.RenderJsonResponse(err, w)
	}
	util.RenderJsonResponse(result, w)
}

func (p *PaasController) GetPaasCpuUsage(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	var apiRequest model.PaasRequest

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")  // may be deprecated..
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")  // may be deprecated..
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{{model.METRIC_NAME_CPU_CORE_PREFIX, model.ALARM_TYPE_CPU}}
	apiRequest.IsLikeQuery = true

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	}
	util.RenderJsonResponse(result, w)
}

func (p *PaasController) GetPaasCpuLoad(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	var apiRequest model.PaasRequest

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NAME_CPU_LOAD_AVG_01_MIN, "1m"},
		{model.METRIC_NAME_CPU_LOAD_AVG_05_MIN, "5m"},
		{model.METRIC_NAME_CPU_LOAD_AVG_15_MIN, "15m"}}

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasMemoryUsage(w http.ResponseWriter, r *http.Request) {
	var result interface{}
	var apiRequest model.PaasRequest
	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	//apiRequest.Args = []model.MetricArg{{model.METRIC_NAME_MEMORY_USAGE, model.ALARM_TYPE_MEMORY}}
	apiRequest.Args = model.MemoryMetricArg{
		model.METRIC_NAME_TOTAL_MEMORY,
		model.METRIC_NAME_FREE_MEMORY,
		model.ALARM_TYPE_MEMORY}
	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMemoryUsage(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasDiskUsage(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NAME_DISK_ROOT_USAGE, "/"},
		{model.METRIC_NAME_DISK_VCAP_USAGE, "data"}}

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasDiskIO(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NAME_DISK_IO_ROOT_READ_BYTES, "/-Read"},
		{model.METRIC_NAME_DISK_IO_ROOT_WRITE_BYTES, "/-Write"},
		{model.METRIC_NAME_DISK_IO_VCAP_READ_BYTES, "data-Read"},
		{model.METRIC_NAME_DISK_IO_VCAP_WRITE_BYTES, "data-Write"}}
	apiRequest.IsLikeQuery = true
	apiRequest.IsRespondKb = true
	apiRequest.IsNonNegativeDerivative = true

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasNetworkByte(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NETWORK_IO_BYTES_SENT, "Sent"},
		{model.METRIC_NETWORK_IO_BYTES_RECV, "Recv"}}
	apiRequest.IsRespondKb = true
	apiRequest.IsNonNegativeDerivative = true

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasNetworkPacket(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NETWORK_IO_PACKET_SENT, "Sent"},
		{model.METRIC_NETWORK_IO_PACKET_RECV, "Recv"}}
	apiRequest.IsRespondKb = true
	apiRequest.IsNonNegativeDerivative = true

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasNetworkDrop(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NETWORK_IO_DROP_IN, "In"},
		{model.METRIC_NETWORK_IO_DROP_OUT, "Out"}}

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasNetworkError(w http.ResponseWriter, r *http.Request) {

	var result interface{}
	var apiRequest model.PaasRequest

	//fmt.Println("request=", r)

	apiRequest.Id = r.FormValue(":id")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")
	apiRequest.Args = []model.MetricArg{
		{model.METRIC_NETWORK_IO_ERR_IN, "In"},
		{model.METRIC_NETWORK_IO_ERR_OUT, "Out"}}

	//fmt.Println("apiRequest=", apiRequest)

	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasMetricStats(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetTopologicalView(w http.ResponseWriter, r *http.Request) {

	var apiRequest model.PaasRequest

	result, errMsg := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetTopologicalView(apiRequest)

	if errMsg != nil {
		util.RenderJsonResponse(errMsg, w)
	} else {
		util.RenderJsonResponse(result, w)
	}
}

func (p *PaasController) GetPaasAllOverview(w http.ResponseWriter, r *http.Request) {

	var apiRequest model.PaasRequest

	apiRequest.Origin = r.FormValue(":origin")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.ServiceName = r.URL.Query().Get("serviceName")
	apiRequest.Index = r.URL.Query().Get("index")
	apiRequest.Addr = r.URL.Query().Get("addr")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy = r.URL.Query().Get("groupBy")

	// PaaS-TA Overview
	result, err := service.GetPaasService(p.txn, p.influxClient, p.databases, p.boshClient).GetPaasOverview(apiRequest)
	if err != nil {
		util.RenderJsonResponse(err, w)
	}

	// Container Overview
	resList, err := service.GetContainerService(p.txn, p.influxClient, p.databases).GetContainerOverview(model.ContainerReq{})
	if err != nil {
		util.RenderJsonResponse(err, w)
	}

	// Bosh Overview
	boshOverview, err := service.GetBoshStatusService(p.txn, p.influxClient, p.databases).GetBoshStatusOverview(model.BoshSummaryReq{})
	if err != nil {
		util.RenderJsonResponse(err, w)
	}

	resultTotal, _ := strconv.Atoi(result.Total)
	resultRunning, _ := strconv.Atoi(result.Running)
	resultCritical, _ := strconv.Atoi(result.Critical)
	resultWarning, _ := strconv.Atoi(result.Warning)
	resultFailed, _ := strconv.Atoi(result.Failed)

	resListTotal, _ := strconv.Atoi(resList.Total)
	resListRunning, _ := strconv.Atoi(resList.Running)
	resListCritical, _ := strconv.Atoi(resList.Critical)
	resListWarning, _ := strconv.Atoi(resList.Warning)
	resListFailed, _ := strconv.Atoi(resList.Failed)

	boshOverviewTotal, _ := strconv.Atoi(boshOverview.Total)
	boshOverviewRunning, _ := strconv.Atoi(boshOverview.Running)
	boshOverviewCritical, _ := strconv.Atoi(boshOverview.Critical)
	boshOverviewWarning, _ := strconv.Atoi(boshOverview.Warning)
	boshOverviewFailed, _ := strconv.Atoi(boshOverview.Failed)

	resultTotal = resultTotal + resListTotal + boshOverviewTotal
	resultRunning = resultRunning + resListRunning + boshOverviewRunning
	resultCritical = resultCritical + resListCritical + boshOverviewCritical
	resultWarning = resultWarning + resListWarning + boshOverviewWarning
	resultFailed = resultFailed + resListFailed + boshOverviewFailed

	result.Total = strconv.Itoa(resultTotal)
	result.Running = strconv.Itoa(resultRunning)
	result.Critical = strconv.Itoa(resultCritical)
	result.Warning = strconv.Itoa(resultWarning)
	result.Failed = strconv.Itoa(resultFailed)

	util.RenderJsonResponse(result, w)
}
