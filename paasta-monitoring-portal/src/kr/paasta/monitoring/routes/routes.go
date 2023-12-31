package routes

import "github.com/tedsuo/rata"

const (
    PING   = "PING"
    LOGIN  = "LOGIN API"
    LOGOUT = "LOGIN OUT"

    MEMBER_JOIN_INFO                      = "MEMBER_JOIN_INFO"
    MEMBER_JOIN_SAVE                      = "MEMBER_JOIN_SAVE"
    MEMBER_JOIN_CHECK_ID                  = "MEMBER_JOIN_CHECK_ID"
    MEMBER_JOIN_CHECK_EMAIL               = "MEMBER_JOIN_CHECK_EMAIL"
    MEMBER_JOIN_CHECK_DUPLICATION_IAAS_ID = "MEMBER_JOIN_CHECK_DUPLICATION_IAAS_ID"
    MEMBER_JOIN_CHECK_DUPLICATION_PAAS_ID = "MEMBER_JOIN_CHECK_DUPLICATION_PAAS_ID"
    MEMBER_JOIN_CHECK_DUPLICATION_CAAS_ID = "MEMBER_JOIN_CHECK_DUPLICATION_CAAS_ID"
    MEMBER_JOIN_CHECK_IAAS                = "MEMBER_JOIN_CHECK_IAAS"
    MEMBER_JOIN_CHECK_PAAS                = "MEMBER_JOIN_CHECK_PAAS"
    MEMBER_JOIN_CHECK_CAAS                = "MEMBER_JOIN_CHECK_CAAS"
    MEMBER_AUTH_CHECK                     = "MEMBER_AUTH_CHECK"
    MEMBER_INFO_VIEW                      = "MEMBER_INFO_VIEW"
    MEMBER_INFO_UPDATE                    = "MEMBER_INFO_UPDATE"
    MEMBER_INFO_DELETE                    = "MEMBER_INFO_DELETE"

    PAAS_ALARM_REALTIME_COUNT = "PAAS_ALARM_REALTIME_COUNT"
    PAAS_ALARM_REALTIME_LIST  = "PAAS_ALARM_REALTIME_LIST"

    PAAS_ALARM_POLICY_LIST   = "PAAS_ALARM_POLICY_LIST"
    PAAS_ALARM_POLICY_UPDATE = "PAAS_ALARM_POLICY_UPDATE"

    PAAS_ALARM_SNS_CHANNEL_LIST   = "PAAS_ALARM_SNS_CHANNEL_LIST"
    PAAS_ALARM_SNS_CHANNEL_CREATE = "PAAS_ALARM_SNS_CHANNEL_CREATE"
    PAAS_ALARM_SNS_CHANNEL_DELETE = "PAAS_ALARM_SNS_CHANNEL_DELETE"
    PAAS_ALARM_SNS_CHANNEL_UPDATE = "PAAS_ALARM_SNS_CHANNEL_UPDATE"

    PAAS_ALARM_STATUS_LIST    = "PAAS_ALARM_STATUS_LIST"
    PAAS_ALARM_STATUS_COUNT   = "PAAS_ALARM_STATUS_COUNT"
    PAAS_ALARM_STATUS_RESOLVE = "PAAS_ALARM_STATUS_RESOLVE"
    PAAS_ALARM_STATUS_DETAIL  = "PAAS_ALARM_DETAIL"
    PAAS_ALARM_STATUS_UPDATE  = "PAAS_ALARM_UPDATE"

    PAAS_ALARM_ACTION_CREATE = "PAAS_ALARM_ACTION_CREATE"
    PAAS_ALARM_ACTION_UPDATE = "PAAS_ALARM_ACTION_UPDATE"
    PAAS_ALARM_ACTION_DELETE = "PAAS_ALARM_ACTION_DELETE"

    PAAS_ALARM_STATISTICS               = "PAAS_ALARM_STATISTICS"
    PAAS_ALARM_STATISTICS_GRAPH_TOTAL   = "PAAS_ALARM_STATISTICS_GRAPH_TOTAL"
    PAAS_ALARM_STATISTICS_GRAPH_SERVICE = "PAAS_ALARM_STATISTICS_GRAPH_SERVICE"
    PAAS_ALARM_STATISTICS_GRAPH_MATRIX  = "PAAS_ALARM_STATISTICS_GRAPH_MATRIX"
    PAAS_ALARM_CONTAINER_DEPLOY         = "PAAS_ALARM_CONTAINER_DEPLOY"
    PAAS_ALARM_DISK_IO_LIST             = "PAAS_ALARM_DISK_IO_LIST"
    PAAS_ALARM_NETWORK_IO_LIST          = "PAAS_ALARM_NETWORK_IO_LIST"
    PAAS_ALARM_TOPPROCESS_LIST          = "PAAS_ALARM_TOPPROCESS_LIST"
    PAAS_ALARM_APP_RESOURCES            = "PAAS_ALARM_APP_RESOURCES"
    PAAS_ALARM_APP_RESOURCES_ALL        = "PAAS_ALARM_APP_RESOURCES_ALL"
    PAAS_ALARM_APP_USAGES               = "PAAS_ALARM_APP_USAGES"
    PAAS_ALARM_APP_MEMORY_USAGES        = "PAAS_ALARM_APP_MEMORY_USAGES"
    PAAS_ALARM_APP_DISK_USAGES          = "PAAS_ALARM_APP_DISK_USAGES"
    PAAS_ALARM_APP_NETWORK_USAGES       = "PAAS_ALARM_APP_NETWORK_USAGES"

    PAAS_BOSH_STATUS_OVERVIEW   = "PAAS_BOSH_STATUS_OVERVIEW"
    PAAS_BOSH_STATUS_SUMMARY    = "PAAS_BOSH_STATUS_SUMMARY"
    PAAS_BOSH_STATUS_TOPPROCESS = "PAAS_BOSH_STATUS_TOPPROCESS"

    PAAS_BOSH_CPU_USAGE_LIST      = "PAAS_BOSH_CPU_USAGE_LIST"
    PAAS_BOSH_CPU_LOAD_LIST       = "PAAS_BOSH_CPU_LOAD_LIST"
    PAAS_BOSH_MEMORY_USAGE_LIST   = "PAAS_BOSH_MEMORY_USAGE_LIST"
    PAAS_BOSH_DISK_USAGE_LIST     = "PAAS_BOSH_DISK_USAGE_LIST"
    PAAS_BOSH_DISK_IO_LIST        = "PAAS_BOSH_DISK_IO_LIST"
    PAAS_BOSH_NETWORK_BYTE_LIST   = "PAAS_BOSH_NETWORK_BYTE_LIST"
    PAAS_BOSH_NETWORK_PACKET_LIST = "PAAS_BOSH_NETWORK_PACKET_LIST"
    PAAS_BOSH_NETWORK_DROP_LIST   = "PAAS_BOSH_NETWORK_DROP_LIST"
    PAAS_BOSH_NETWORK_ERROR_LIST  = "PAAS_BOSH_NETWORK_ERROR_LIST"

    PAAS_PAASTA_OVERVIEW          = "PAAS_PAASTA_OVERVIEW"
    PAAS_PAASTA_SUMMARY           = "PAAS_PAASTA_SUMMARY"
    PAAS_PAASTA_TOPPROCESS_MEMORY = "PAAS_PAASTA_TOPPROCESS_MEMORY"
    PAAS_PAASTA_OVERVIEW_STATUS   = "PAAS_PAASTA_OVERVIEW_STATUS"

    PAAS_PAASTA_CPU_USAGE      = "PAAS_PAASTA_CPU_USAGE"
    PAAS_PAASTA_CPU_LOAD       = "PAAS_PAASTA_CPU_LOAD"
    PAAS_PAASTA_MEMORY_USAGE   = "PAAS_PAASTA_MEMORY_USAGE"
    PAAS_PAASTA_DISK_USAGE     = "PAAS_PAASTA_DISK_USAGE"
    PAAS_PAASTA_DISK_IO        = "PAAS_PAASTA_DISK_IO"
    PAAS_PAASTA_NETWORK_BYTE   = "PAAS_PAASTA_NETWORK_BYTE"
    PAAS_PAASTA_NETWORK_PACKET = "PAAS_PAASTA_NETWORK_PACKET"
    PAAS_PAASTA_NETWORK_DROP   = "PAAS_PAASTA_NETWORK_DROP"
    PAAS_PAASTA_NETWORK_ERROR  = "PAAS_PAASTA_NETWORK_ERROR"

    PAAS_TOPOLOGICAL_VIEW = "PAAS_TOPOLOGICAL_VIEW"

    PAAS_CELL_OVERVIEW          = "PAAS_CELL_OVERVIEW"
    PAAS_CONTAINER_OVERVIEW     = "PAAS_CONTAINER_OVERVIEW"
    PAAS_CONTAINER_SUMMARY      = "PAAS_CONTAINER_SUMMARY"
    PAAS_CONTAINER_RELATIONSHIP = "PAAS_CONTAINER_RELATIONSHIP"

    PAAS_CELL_OVERVIEW_STATE_LIST      = "PAAS_CELL_OVERVIEW_STATE_LIST"
    PAAS_CONTAINER_OVERVIEW_STATE_LIST = "PAAS_CONTAINER_OVERVIEW_STATE_LIST"

    PAAS_CONTAINER_OVERVIEW_MAIN = "PAAS_CONTAINER_OVERVIEW_MAIN"

    PAAS_CONTAINER_CPU_USAGE_LIST     = "PAAS_CONTAINER_CPU_USAGE_LIST"
    PAAS_CONTAINER_CPU_LOADS_LIST     = "PAAS_CONTAINER_CPU_LOADS_LIST"
    PAAS_CONTAINER_MEMORY_USAGE_LIST  = "PAAS_CONTAINER_MEMORY_USAGE_LIST"
    PAAS_CONTAINER_DISK_USAGE_LIST    = "PAAS_CONTAINER_DISK_USAGE_LIST"
    PAAS_CONTAINER_NETWORK_BYTE_LIST  = "PAAS_CONTAINER_NETWORK_BYTE_LIST"
    PAAS_CONTAINER_NETWORK_DROP_LIST  = "PAAS_CONTAINER_NETWORK_DROP_LIST"
    PAAS_CONTAINER_NETWORK_ERROR_LIST = "PAAS_CONTAINER_NETWORK_ERROR_LIST"

    PAAS_LOG_SEARCH = "PAAS_LOG_SEARCH"

    // potal - paas api
    PAAS_APP_CPU_USAGES                = "PAAS_APP_CPU_USAGES"
    PAAS_APP_MEMORY_USAGES             = "PAAS_APP_MEMORY_USAGES"
    PAAS_APP_NETWORK_USAGES            = "PAAS_APP_NETWORK_USAGES"
    PAAS_APP_AUTOSCALING_POLICY_UPDATE = "PAAS_APP_AUTOSCALING_POLICY_UPDATE"
    PAAS_APP_AUTOSCALING_POLICY_INFO   = "PAAS_APP_AUTOSCALING_POLICY_INFO"
    PAAS_APP_POLICY_UPDATE             = "PAAS_APP_POLICY_UPDATE"
    PAAS_APP_POLICY_INFO               = "PAAS_APP_POLICY_INFO"
    PAAS_APP_ALARM_LIST                = "PAAS_APP_ALARM_LIST"
    PAAS_APP_POLICY_DELETE             = "PAAS_APP_POLICY_DELETE"
    PAAS_PAAS_ALL_OVERVIEW             = "PAAS_PAAS_ALL_OVERVIEW"

    // SAAS
    SAAS_API_APPLICATION_LIST   = "SAAS_API_APPLICATION_LIST"
    SAAS_API_APPLICATION_STATUS = "SAAS_API_APPLICATION_STATUS"
    SAAS_API_APPLICATION_GAUGE  = "SAAS_API_APPLICATION_GAUGE"
    SAAS_API_APPLICATION_REMOVE = "SAAS_API_APPLICATION_REMOVE"

    SAAS_ALARM_INFO               = "SAAS_ALARM_INFO"
    SAAS_ALARM_SNS_INFO           = "SAAS_ALARM_SNS_INFO"
    SAAS_ALARM_UPDATE             = "SAAS_ALARM_UPDATE"
    SAAS_ALARM_LOG                = "SAAS_ALARM_LOG"
    SAAS_ALARM_COUNT              = "SAAS_ALARM_COUNT"
    SAAS_ALARM_SNS_SAVE           = "SAAS_ALARM_SNS_SAVE"
    SAAS_ALARM_STATUS_UPDATE      = "SAAS_ALARM_STATUS_UPDATE"
    SAAS_ALARM_ACTION             = "SAAS_ALARM_ACTION"
    SAAS_ALARM_ACTION_DELETE      = "SAAS_ALARM_ACTION_DELETE"
    SAAS_ALARM_ACTION_UPDATE      = "SAAS_ALARM_ACTION_UPDATE"
    SAAS_ALARM_SNS_CHANNEL_LIST   = "SAAS_ALARM_SNS_CHANNEL_LIST"
    SAAS_ALARM_SNS_CHANNEL_DELETE = "SAAS_ALARM_SNS_CHANNEL_DELETE"
    SAAS_ALARM_ACTION_LIST        = "SAAS_ALARM_ACTION_LIST"

    // CAAS_API
    CAAS_K8S_CLUSTER_AVG = "CAAS_K8S_CLUSTER_AVG"
    CAAS_WORK_NODE_LIST  = "CAAS_WORK_NODE_LIST"
    CAAS_WORK_NODE_INFO  = "CAAS_WORK_NODE_INFO"
    CAAS_CONTIANER_LIST  = "CAAS_CONTIANER_LIST"
    CAAS_CONTIANER_INFO  = "CAAS_CONTIANER_INFO"
    CAAS_CONTIANER_LOG   = "CAAS_CONTIANER_LOG"

    CAAS_CLUSTER_OVERVIEW  = "CAAS_CLUSTER_OVERVIEW"
    CAAS_WORKLOADS_STATUS  = "CAAS_WORKLOADS_STATUS"
    CAAS_MASTER_NODE_USAGE = "CAAS_MASTER_NODE_USAGE"

    CAAS_WORKLOADS_CONTI_SUMMARY = "CAAS_WORKLOADS_CONTI_SUMMARY"
    CAAS_WORKLOADS_USAGE         = "CAAS_WORKLOADS_USAGE"
    CAAS_POD_STAT                = "CAAS_POD_STAT"
    CAAS_POD_LIST                = "CAAS_POD_LIST"
    CAAS_POD_INFO                = "CAAS_POD_INFO"

    CAAS_WORK_NODE_GRAPH     = "CAAS_WORK_NODE_GRAPH"
    CAAS_WORKLOADS_GRAPH     = "CAAS_WORKLOADS_GRAPH"
    CAAS_POD_GRAPH           = "CAAS_POD_GRAPH"
    CAAS_CONTIANER_GRAPH     = "CAAS_CONTIANER_GRAPH"
    CAAS_WORK_NODE_AVG       = "CAAS_WORK_NODE_AVG"
    CAAS_WORK_NODE_GRAPHLIST = "CAAS_WORK_NODE_GRAPHLIST"

    CAAS_ALARM_INFO               = "CAAS_ALARM_INFO"
    CAAS_ALARM_SNS_INFO           = "CAAS_ALARM_SNS_INFO"
    CAAS_ALARM_UPDATE             = "CAAS_ALARM_UPDATE"
    CAAS_ALARM_LOG                = "CAAS_ALARM_LOG"
    CAAS_ALARM_COUNT              = "CAAS_ALARM_COUNT"
    CAAS_ALARM_SNS_SAVE           = "CAAS_ALARM_SNS_SAVE"
    CAAS_ALARM_STATUS_UPDATE      = "CAAS_ALARM_STATUS_UPDATE"
    CAAS_ALARM_ACTION             = "CAAS_ALARM_ACTION"
    CAAS_ALARM_ACTION_DELETE      = "CAAS_ALARM_ACTION_DELETE"
    CAAS_ALARM_ACTION_UPDATE      = "CAAS_ALARM_ACTION_UPDATE"
    CAAS_ALARM_SNS_CHANNEL_LIST   = "CAAS_ALARM_SNS_CHANNEL_LIST"
    CAAS_ALARM_SNS_CHANNEL_DELETE = "CAAS_ALARM_SNS_CHANNEL_DELETE"
    CAAS_ALARM_ACTION_LIST        = "CAAS_ALARM_ACTION_LIST"

    // Web Resource
    Main   = "Main"
    Static = "Static"

    // 2021.11.02 - IAAS 모니터링
    IAAS_ALARM_STATUS = "IAAS_ALARM_STATUS"

    IAAS_ALARM_POLICY_LIST   = "IAAS_ALARM_POLICY_LIST"
    IAAS_ALARM_POLICY_UPDATE = "IAAS_ALARM_POLICY_UPDATE"

    IAAS_ALARM_SNS_CHANNEL_LIST   = "IAAS_ALARM_SNS_CHANNEL_LIST"
    IAAS_ALARM_SNS_CHANNEL_CREATE = "IAAS_ALARM_SNS_CHANNEL_CREATE"
    IAAS_ALARM_SNS_CHANNEL_DELETE = "IAAS_ALARM_SNS_CHANNEL_DELETE"
    IAAS_ALARM_SNS_CHANNEL_UPDATE = "IAAS_ALARM_SNS_CHANNEL_UPDATE"

    IAAS_ALARM_STATUS_LIST    = "IAAS_ALARM_STATUS_LIST"
    IAAS_ALARM_STATUS_COUNT   = "IAAS_ALARM_STATUS_COUNT"
    IAAS_ALARM_STATUS_RESOLVE = "IAAS_ALARM_STATUS_RESOLVE"
    IAAS_ALARM_STATUS_DETAIL  = "IAAS_ALARM_DETAIL"
    IAAS_ALARM_STATUS_UPDATE  = "IAAS_ALARM_UPDATE"

    IAAS_ALARM_ACTION_CREATE = "IAAS_ALARM_ACTION_CREATE"
    IAAS_ALARM_ACTION_UPDATE = "IAAS_ALARM_ACTION_UPDATE"
    IAAS_ALARM_ACTION_DELETE = "IAAS_ALARM_ACTION_DELETE"

    IAAS_ALARM_STATISTICS               = "IAAS_ALARM_STATISTICS"
    IAAS_ALARM_STATISTICS_GRAPH_TOTAL   = "IAAS_ALARM_STATISTICS_GRAPH_TOTAL"
    IAAS_ALARM_STATISTICS_GRAPH_SERVICE = "IAAS_ALARM_STATISTICS_GRAPH_SERVICE"
    IAAS_ALARM_STATISTICS_GRAPH_MATRIX  = "IAAS_ALARM_STATISTICS_GRAPH_MATRIX"

    IAAS_GET_HYPERVISOR_LIST     = "IAAS_GET_HYPERVISOR_LIST"
    IAAS_GET_HYPER_STATISTICS    = "IAAS_GET_HYPER_STATISTICS"
    IAAS_GET_SERVER_LIST         = "IAAS_GET_SERVER_LIST"
    IAAS_GET_PROJECT_LIST        = "IAAS_GET_PROJECT_LIST"
    IAAS_GET_INSTANCE_USAGE_LIST = "IAAS_GET_INSTANCE_USAGE_LIST"

    IAAS_GET_CPU_USAGE        = "IAAS_GET_CPU_USAGE"
    IAAS_GET_MEMORY_USAGE     = "IAAS_GET_MEMORY_USAGE"
    IAAS_GET_DISK_USAGE       = "IAAS_GET_DISK_USAGE"
    IAAS_GET_CPU_LOAD_AVERAGE = "IAAS_GET_CPU_LOAD_AVERAGE"
    IAAS_GET_DISK_IO_RATE     = "IAAS_GET_DISK_IO_RATE"
    IAAS_GET_NETWORK_IO_BTYES = "IAAS_GET_NETWORK_IO_BTYES"

    PAAS_DIAGRAM = "PAAS_DIAGRAM"
)

var Routes = rata.Routes{

    {Path: "/v2/ping", Method: "GET", Name: PING},
    {Path: "/v2/login", Method: "POST", Name: LOGIN},
    {Path: "/v2/logout", Method: "POST", Name: LOGOUT},

    {Path: "/v2/member/join", Method: "GET", Name: MEMBER_JOIN_INFO},
    {Path: "/v2/member/join", Method: "POST", Name: MEMBER_JOIN_SAVE},
    {Path: "/v2/member/join/check/id/:id", Method: "GET", Name: MEMBER_JOIN_CHECK_ID},
    {Path: "/v2/member/join/check/email/:email", Method: "GET", Name: MEMBER_JOIN_CHECK_EMAIL},

    {Path: "/v2/member/auth/check/:id", Method: "GET", Name: MEMBER_AUTH_CHECK},
    {Path: "/v2/member/info/view", Method: "POST", Name: MEMBER_INFO_VIEW},
    {Path: "/v2/member/info", Method: "PATCH", Name: MEMBER_INFO_UPDATE},
    {Path: "/v2/member/info", Method: "DELETE", Name: MEMBER_INFO_DELETE},

    // Web Resource
    {Path: "/", Method: "GET", Name: Main},
    {Path: "/public/", Method: "GET", Name: Static},
}

var PaasRoutes = rata.Routes{
    {Path: "/v2/member/join/check/duplication/paas/:id", Method: "GET", Name: MEMBER_JOIN_CHECK_DUPLICATION_PAAS_ID},
    {Path: "/v2/member/join/check/paas", Method: "POST", Name: MEMBER_JOIN_CHECK_PAAS},

    {Path: "/v2/paas/alarm/realtime/count", Method: "GET", Name: PAAS_ALARM_REALTIME_COUNT},
    {Path: "/v2/paas/alarm/realtime/list", Method: "GET", Name: PAAS_ALARM_REALTIME_LIST},

    {Path: "/v2/paas/alarm/policies", Method: "GET", Name: PAAS_ALARM_POLICY_LIST},
    {Path: "/v2/paas/alarm/policy", Method: "PUT", Name: PAAS_ALARM_POLICY_UPDATE},

    {Path: "/v2/paas/alarm/sns/channel", Method: "POST", Name: PAAS_ALARM_SNS_CHANNEL_CREATE},
    {Path: "/v2/paas/alarm/sns/channel/list", Method: "GET", Name: PAAS_ALARM_SNS_CHANNEL_LIST},
    {Path: "/v2/paas/alarm/sns/channel/:id", Method: "DELETE", Name: PAAS_ALARM_SNS_CHANNEL_DELETE},
    {Path: "/v2/paas/alarm/sns/channel", Method: "PUT", Name: PAAS_ALARM_SNS_CHANNEL_UPDATE}, // 2021.05.18 - PaaS 채널 SNS 수정 기능 추가

    {Path: "/v2/paas/alarm/statuses", Method: "GET", Name: PAAS_ALARM_STATUS_LIST},
    {Path: "/v2/paas/alarm/status/count", Method: "GET", Name: PAAS_ALARM_STATUS_COUNT},
    {Path: "/v2/paas/alarm/status/:id", Method: "GET", Name: PAAS_ALARM_STATUS_DETAIL},
    {Path: "/v2/paas/alarm/status/:id", Method: "PUT", Name: PAAS_ALARM_STATUS_UPDATE},
    {Path: "/v2/paas/alarm/status/:resolveStatus", Method: "GET", Name: PAAS_ALARM_STATUS_RESOLVE},

    {Path: "/v2/paas/alarm/action", Method: "POST", Name: PAAS_ALARM_ACTION_CREATE},
    {Path: "/v2/paas/alarm/action/:actionId", Method: "PATCH", Name: PAAS_ALARM_ACTION_UPDATE},
    {Path: "/v2/paas/alarm/action/:actionId", Method: "DELETE", Name: PAAS_ALARM_ACTION_DELETE},

    {Path: "/v2/paas/alarm/statistics", Method: "GET", Name: PAAS_ALARM_STATISTICS},
    {Path: "/v2/paas/alarm/statistics/graph/total", Method: "GET", Name: PAAS_ALARM_STATISTICS_GRAPH_TOTAL},
    {Path: "/v2/paas/alarm/statistics/graph/service", Method: "GET", Name: PAAS_ALARM_STATISTICS_GRAPH_SERVICE},
    {Path: "/v2/paas/alarm/statistics/graph/matrix", Method: "GET", Name: PAAS_ALARM_STATISTICS_GRAPH_MATRIX},

    {Path: "/v2/paas/alarm/container/deploy", Method: "GET", Name: PAAS_ALARM_CONTAINER_DEPLOY},
    {Path: "/v2/paas/alarm/disk/io/:origin", Method: "GET", Name: PAAS_ALARM_DISK_IO_LIST},
    {Path: "/v2/paas/alarm/network/io/:origin", Method: "GET", Name: PAAS_ALARM_NETWORK_IO_LIST},
    {Path: "/v2/paas/alarm/topprocess/:origin", Method: "GET", Name: PAAS_ALARM_TOPPROCESS_LIST},
    {Path: "/v2/paas/alarm/app/resources", Method: "GET", Name: PAAS_ALARM_APP_RESOURCES},
    {Path: "/v2/paas/alarm/app/resources/all", Method: "GET", Name: PAAS_ALARM_APP_RESOURCES_ALL},
    {Path: "/v2/paas/alarm/app/cpu/:guid/:idx/usages", Method: "GET", Name: PAAS_ALARM_APP_USAGES},
    {Path: "/v2/paas/alarm/app/memory/:guid/:idx/usages", Method: "GET", Name: PAAS_ALARM_APP_MEMORY_USAGES},
    {Path: "/v2/paas/alarm/app/disk/:guid/:idx/usages", Method: "GET", Name: PAAS_ALARM_APP_DISK_USAGES},
    {Path: "/v2/paas/alarm/app/network/:guid/:idx/usages", Method: "GET", Name: PAAS_ALARM_APP_NETWORK_USAGES},

    {Path: "/v2/paas/bosh/overview", Method: "GET", Name: PAAS_BOSH_STATUS_OVERVIEW},
    {Path: "/v2/paas/bosh/summary", Method: "GET", Name: PAAS_BOSH_STATUS_SUMMARY},
    {Path: "/v2/paas/bosh/topprocess/:id/memory", Method: "GET", Name: PAAS_BOSH_STATUS_TOPPROCESS},

    {Path: "/v2/paas/bosh/cpu/:id/usages", Method: "GET", Name: PAAS_BOSH_CPU_USAGE_LIST},
    {Path: "/v2/paas/bosh/cpu/:id/loads", Method: "GET", Name: PAAS_BOSH_CPU_LOAD_LIST},
    {Path: "/v2/paas/bosh/memory/:id/usages", Method: "GET", Name: PAAS_BOSH_MEMORY_USAGE_LIST},
    {Path: "/v2/paas/bosh/disk/:id/usages", Method: "GET", Name: PAAS_BOSH_DISK_USAGE_LIST},
    {Path: "/v2/paas/bosh/disk/:id/ios", Method: "GET", Name: PAAS_BOSH_DISK_IO_LIST},
    {Path: "/v2/paas/bosh/network/:id/bytes", Method: "GET", Name: PAAS_BOSH_NETWORK_BYTE_LIST},
    {Path: "/v2/paas/bosh/network/:id/packets", Method: "GET", Name: PAAS_BOSH_NETWORK_PACKET_LIST},
    {Path: "/v2/paas/bosh/network/:id/drops", Method: "GET", Name: PAAS_BOSH_NETWORK_DROP_LIST},
    {Path: "/v2/paas/bosh/network/:id/errors", Method: "GET", Name: PAAS_BOSH_NETWORK_ERROR_LIST},

    // PaaS Overview
    {Path: "/v2/paas/paasta/overview", Method: "GET", Name: PAAS_PAASTA_OVERVIEW},
    {Path: "/v2/paas/paasta/summary", Method: "GET", Name: PAAS_PAASTA_SUMMARY},
    {Path: "/v2/paas/paasta/topprocess/:id/memory", Method: "GET", Name: PAAS_PAASTA_TOPPROCESS_MEMORY},
    {Path: "/v2/paas/paasta/overview/:status", Method: "GET", Name: PAAS_PAASTA_OVERVIEW_STATUS},

    // PaaS Detail
    {Path: "/v2/paas/paasta/cpu/:id/usages", Method: "GET", Name: PAAS_PAASTA_CPU_USAGE},
    {Path: "/v2/paas/paasta/cpu/:id/loads", Method: "GET", Name: PAAS_PAASTA_CPU_LOAD},
    {Path: "/v2/paas/paasta/memory/:id/usages", Method: "GET", Name: PAAS_PAASTA_MEMORY_USAGE},
    {Path: "/v2/paas/paasta/disk/:id/usages", Method: "GET", Name: PAAS_PAASTA_DISK_USAGE},
    {Path: "/v2/paas/paasta/disk/:id/ios", Method: "GET", Name: PAAS_PAASTA_DISK_IO},
    {Path: "/v2/paas/paasta/network/:id/bytes", Method: "GET", Name: PAAS_PAASTA_NETWORK_BYTE},
    {Path: "/v2/paas/paasta/network/:id/packets", Method: "GET", Name: PAAS_PAASTA_NETWORK_PACKET},
    {Path: "/v2/paas/paasta/network/:id/drops", Method: "GET", Name: PAAS_PAASTA_NETWORK_DROP},
    {Path: "/v2/paas/paasta/network/:id/errors", Method: "GET", Name: PAAS_PAASTA_NETWORK_ERROR},

    // PaaS Dashboard
    {Path: "/v2/paas/main/topological", Method: "GET", Name: PAAS_TOPOLOGICAL_VIEW},

    // container overview
    {Path: "/v2/paas/cell/overview", Method: "GET", Name: PAAS_CELL_OVERVIEW},
    {Path: "/v2/paas/container/overview", Method: "GET", Name: PAAS_CONTAINER_OVERVIEW},
    {Path: "/v2/paas/container/summary", Method: "GET", Name: PAAS_CONTAINER_SUMMARY},
    {Path: "/v2/paas/container/relationship/:name", Method: "GET", Name: PAAS_CONTAINER_RELATIONSHIP},

    {Path: "/v2/paas/cell/overview/:status", Method: "GET", Name: PAAS_CELL_OVERVIEW_STATE_LIST},
    {Path: "/v2/paas/container/overview/:status", Method: "GET", Name: PAAS_CONTAINER_OVERVIEW_STATE_LIST},

    {Path: "/v2/paas/container/relationship", Method: "GET", Name: PAAS_CONTAINER_OVERVIEW_MAIN},

    {Path: "/v2/paas/container/cpu/:id/usages", Method: "GET", Name: PAAS_CONTAINER_CPU_USAGE_LIST},
    {Path: "/v2/paas/container/cpu/:id/loads", Method: "GET", Name: PAAS_CONTAINER_CPU_LOADS_LIST},
    {Path: "/v2/paas/container/memory/:id/usages", Method: "GET", Name: PAAS_CONTAINER_MEMORY_USAGE_LIST},
    {Path: "/v2/paas/container/disk/:id/usages", Method: "GET", Name: PAAS_CONTAINER_DISK_USAGE_LIST},
    {Path: "/v2/paas/container/network/:id/bytes", Method: "GET", Name: PAAS_CONTAINER_NETWORK_BYTE_LIST},
    {Path: "/v2/paas/container/network/:id/drops", Method: "GET", Name: PAAS_CONTAINER_NETWORK_DROP_LIST},
    {Path: "/v2/paas/container/network/:id/errors", Method: "GET", Name: PAAS_CONTAINER_NETWORK_ERROR_LIST},

    {Path: "/v2/paas/log", Method: "GET", Name: PAAS_LOG_SEARCH},

    // potal - paas api
    {Path: "/v2/paas/app/instance/:guid/:idx/cpu/usages", Method: "GET", Name: PAAS_APP_CPU_USAGES},
    {Path: "/v2/paas/app/instance/:guid/:idx/memory/usages", Method: "GET", Name: PAAS_APP_MEMORY_USAGES},
    {Path: "/v2/paas/app/instance/:guid/:idx/network/bytes", Method: "GET", Name: PAAS_APP_NETWORK_USAGES},

    {Path: "/v2/paas/app/autoscaling/policy", Method: "POST", Name: PAAS_APP_AUTOSCALING_POLICY_UPDATE},
    {Path: "/v2/paas/app/autoscaling/policy", Method: "GET", Name: PAAS_APP_AUTOSCALING_POLICY_INFO},
    {Path: "/v2/paas/app/alarm/policy", Method: "POST", Name: PAAS_APP_POLICY_UPDATE},
    {Path: "/v2/paas/app/alarm/policy", Method: "GET", Name: PAAS_APP_POLICY_INFO},
    {Path: "/v2/paas/app/alarm/list", Method: "GET", Name: PAAS_APP_ALARM_LIST},
    {Path: "/v2/paas/app/policy/:guid", Method: "DELETE", Name: PAAS_APP_POLICY_DELETE},
    {Path: "/v2/paas/all/overview", Method: "GET", Name: PAAS_PAAS_ALL_OVERVIEW},

    {Path: "/v2/paas/diagram", Method: "GET", Name: PAAS_DIAGRAM},
}

var CaasRoutes = rata.Routes{
    {Path: "/v2/member/join/check/duplication/caas/:id", Method: "GET", Name: MEMBER_JOIN_CHECK_DUPLICATION_CAAS_ID},
    {Path: "/v2/member/join/check/caas", Method: "POST", Name: MEMBER_JOIN_CHECK_CAAS},
    {Path: "/v2/caas/monitoring/clusterAvg", Method: "GET", Name: CAAS_K8S_CLUSTER_AVG},
    {Path: "/v2/caas/monitoring/workerNodeList", Method: "GET", Name: CAAS_WORK_NODE_LIST},
    {Path: "/v2/caas/monitoring/workerNodeInfo", Method: "GET", Name: CAAS_WORK_NODE_INFO},
    {Path: "/v2/caas/monitoring/contiList", Method: "GET", Name: CAAS_CONTIANER_LIST},
    {Path: "/v2/caas/monitoring/contiInfo", Method: "GET", Name: CAAS_CONTIANER_INFO},
    {Path: "/v2/caas/monitoring/contiInfoLog", Method: "GET", Name: CAAS_CONTIANER_LOG},
    {Path: "/v2/caas/monitoring/clusterOverview", Method: "GET", Name: CAAS_CLUSTER_OVERVIEW},
    {Path: "/v2/caas/monitoring/workloadsStatus", Method: "GET", Name: CAAS_WORKLOADS_STATUS},
    {Path: "/v2/caas/monitoring/masterNodeUsage", Method: "GET", Name: CAAS_MASTER_NODE_USAGE},
    {Path: "/v2/caas/monitoring/workNodeAvg", Method: "GET", Name: CAAS_WORK_NODE_AVG},
    {Path: "/v2/caas/monitoring/workerloadsConainerSummary", Method: "GET", Name: CAAS_WORKLOADS_CONTI_SUMMARY},
    {Path: "/v2/caas/monitoring/workloadsUsage", Method: "GET", Name: CAAS_WORKLOADS_USAGE},
    {Path: "/v2/caas/monitoring/podStat", Method: "GET", Name: CAAS_POD_STAT},
    {Path: "/v2/caas/monitoring/podList", Method: "GET", Name: CAAS_POD_LIST},
    {Path: "/v2/caas/monitoring/podInfo", Method: "GET", Name: CAAS_POD_INFO},
    {Path: "/v2/caas/monitoring/workerNodeGraph", Method: "GET", Name: CAAS_WORK_NODE_GRAPH},
    {Path: "/v2/caas/monitoring/workerNodeGraphList", Method: "GET", Name: CAAS_WORK_NODE_GRAPHLIST},
    {Path: "/v2/caas/monitoring/workloadsGraph", Method: "GET", Name: CAAS_WORKLOADS_GRAPH},
    {Path: "/v2/caas/monitoring/podGraph", Method: "GET", Name: CAAS_POD_GRAPH},
    {Path: "/v2/caas/monitoring/containerGraph", Method: "GET", Name: CAAS_CONTIANER_GRAPH},

    {Path: "/v2/caas/monitoring/alarmInfo", Method: "GET", Name: CAAS_ALARM_INFO},         // 완료
    {Path: "/v2/caas/monitoring/alarmUpdate", Method: "PUT", Name: CAAS_ALARM_UPDATE},     // 완료
    {Path: "/v2/caas/monitoring/alarmLog", Method: "GET", Name: CAAS_ALARM_LOG},           // 완료
    {Path: "/v2/caas/monitoring/alarmSnsInfo", Method: "GET", Name: CAAS_ALARM_SNS_INFO},  // 완료
    {Path: "/v2/caas/monitoring/alarmCount", Method: "GET", Name: CAAS_ALARM_COUNT},       // 완료
    {Path: "/v2/caas/monitoring/alarmSnsSave", Method: "POST", Name: CAAS_ALARM_SNS_SAVE}, // 완료

    {Path: "/v2/caas/monitoring/alarmStatus/:id", Method: "PUT", Name: CAAS_ALARM_STATUS_UPDATE},
    {Path: "/v2/caas/monitoring/alarmAction", Method: "POST", Name: CAAS_ALARM_ACTION},
    {Path: "/v2/caas/monitoring/alarmAction/:id", Method: "DELETE", Name: CAAS_ALARM_ACTION_DELETE},
    {Path: "/v2/caas/monitoring/alarmAction/:id", Method: "PATCH", Name: CAAS_ALARM_ACTION_UPDATE},
    {Path: "/v2/caas/monitoring/snsChannel/list", Method: "GET", Name: CAAS_ALARM_SNS_CHANNEL_LIST},
    {Path: "/v2/caas/monitoring/snsChannel/:id", Method: "DELETE", Name: CAAS_ALARM_SNS_CHANNEL_DELETE},
    {Path: "/v2/caas/monitoring/alarmAction/:id", Method: "GET", Name: CAAS_ALARM_ACTION_LIST},
}

var SaasRoutes = rata.Routes{
    {Path: "/v2/saas/app/application/list", Method: "GET", Name: SAAS_API_APPLICATION_LIST},
    {Path: "/v2/saas/app/application/status", Method: "GET", Name: SAAS_API_APPLICATION_STATUS},
    {Path: "/v2/saas/app/application/gauge", Method: "GET", Name: SAAS_API_APPLICATION_GAUGE},

    {Path: "/v2/saas/app/application/alarmInfo", Method: "GET", Name: SAAS_ALARM_INFO},         // 완료
    {Path: "/v2/saas/app/application/alarmSnsInfo", Method: "GET", Name: SAAS_ALARM_SNS_INFO},  // 완료
    {Path: "/v2/saas/app/application/alarmUpdate", Method: "PUT", Name: SAAS_ALARM_UPDATE},     // 완료
    {Path: "/v2/saas/app/application/alarmLog", Method: "GET", Name: SAAS_ALARM_LOG},           // 완료
    {Path: "/v2/saas/app/application/alarmCount", Method: "GET", Name: SAAS_ALARM_COUNT},       // 완료
    {Path: "/v2/saas/app/application/alarmSnsSave", Method: "POST", Name: SAAS_ALARM_SNS_SAVE}, // 완료

    {Path: "/v2/saas/app/application/alarmStatus/:id", Method: "PUT", Name: SAAS_ALARM_STATUS_UPDATE},
    {Path: "/v2/saas/app/application/alarmAction", Method: "POST", Name: SAAS_ALARM_ACTION},
    {Path: "/v2/saas/app/application/alarmAction/:id", Method: "DELETE", Name: SAAS_ALARM_ACTION_DELETE},
    {Path: "/v2/saas/app/application/alarmAction/:id", Method: "PATCH", Name: SAAS_ALARM_ACTION_UPDATE},
    {Path: "/v2/saas/app/application/snsChannel/list", Method: "GET", Name: SAAS_ALARM_SNS_CHANNEL_LIST},
    {Path: "/v2/saas/app/application/snsChannel/:id", Method: "DELETE", Name: SAAS_ALARM_SNS_CHANNEL_DELETE},
    {Path: "/v2/saas/app/application/alarmAction/:id", Method: "GET", Name: SAAS_ALARM_ACTION_LIST},
    {Path: "/v2/saas/app/application/removeAgentId", Method: "DELETE", Name: SAAS_API_APPLICATION_REMOVE},
}

// TODO 2021.11.01 - IAAS 모니터링
var IaasRoutes = rata.Routes{
    {Path: "/v2/member/join/check/duplication/iaas/:id", Method: "GET", Name: MEMBER_JOIN_CHECK_DUPLICATION_IAAS_ID},
    {Path: "/v2/member/join/check/iaas", Method: "POST", Name: MEMBER_JOIN_CHECK_IAAS},

    {Path: "/v2/iaas/alarm/policies", Method: "GET", Name: IAAS_ALARM_POLICY_LIST},
    {Path: "/v2/iaas/alarm/policy", Method: "PUT", Name: IAAS_ALARM_POLICY_UPDATE},

    {Path: "/v2/iaas/alarm/sns/channel", Method: "POST", Name: IAAS_ALARM_SNS_CHANNEL_CREATE},
    {Path: "/v2/iaas/alarm/sns/channel/list", Method: "GET", Name: IAAS_ALARM_SNS_CHANNEL_LIST},
    {Path: "/v2/iaas/alarm/sns/channel/:id", Method: "DELETE", Name: IAAS_ALARM_SNS_CHANNEL_DELETE},
    {Path: "/v2/iaas/alarm/sns/channel", Method: "PUT", Name: IAAS_ALARM_SNS_CHANNEL_UPDATE}, // 2021.05.18 - PaaS 채널 SNS 수정 기능 추가

    {Path: "/v2/iaas/alarm/statuses", Method: "GET", Name: IAAS_ALARM_STATUS_LIST},
    {Path: "/v2/iaas/alarm/status/count", Method: "GET", Name: IAAS_ALARM_STATUS_COUNT},
    {Path: "/v2/iaas/alarm/status/:id", Method: "GET", Name: IAAS_ALARM_STATUS_DETAIL},
    {Path: "/v2/iaas/alarm/status/:id", Method: "PUT", Name: IAAS_ALARM_STATUS_UPDATE},
    {Path: "/v2/iaas/alarm/status/:resolveStatus", Method: "GET", Name: IAAS_ALARM_STATUS_RESOLVE},

    {Path: "/v2/iaas/alarm/action", Method: "POST", Name: IAAS_ALARM_ACTION_CREATE},
    {Path: "/v2/iaas/alarm/action/:actionId", Method: "PATCH", Name: IAAS_ALARM_ACTION_UPDATE},
    {Path: "/v2/iaas/alarm/action/:actionId", Method: "DELETE", Name: IAAS_ALARM_ACTION_DELETE},

    {Path: "/v2/iaas/alarm/statistics", Method: "GET", Name: IAAS_ALARM_STATISTICS},
    {Path: "/v2/iaas/alarm/statistics/graph/total", Method: "GET", Name: IAAS_ALARM_STATISTICS_GRAPH_TOTAL},
    {Path: "/v2/iaas/alarm/statistics/graph/service", Method: "GET", Name: IAAS_ALARM_STATISTICS_GRAPH_SERVICE},
    {Path: "/v2/iaas/alarm/statistics/graph/matrix", Method: "GET", Name: IAAS_ALARM_STATISTICS_GRAPH_MATRIX},

    {Path: "/v2/iaas/hypervisor/list", Method: "GET", Name: IAAS_GET_HYPERVISOR_LIST},
    {Path: "/v2/iaas/hyper/statistics", Method: "GET", Name: IAAS_GET_HYPER_STATISTICS},
    {Path: "/v2/iaas/server/list", Method: "GET", Name: IAAS_GET_SERVER_LIST},
    {Path: "/v2/iaas/project/list", Method: "GET", Name: IAAS_GET_PROJECT_LIST},
    {Path: "/v2/iaas/instance/usage/list", Method: "GET", Name: IAAS_GET_INSTANCE_USAGE_LIST},
    // IAAS 모니터링 차트 데이터
    {Path: "/v2/iaas/instance/cpu/usage/", Method: "GET", Name: IAAS_GET_CPU_USAGE},
    {Path: "/v2/iaas/instance/memory/usage/", Method: "GET", Name: IAAS_GET_MEMORY_USAGE},
    {Path: "/v2/iaas/instance/disk/usage/", Method: "GET", Name: IAAS_GET_DISK_USAGE},
    {Path: "/v2/iaas/instance/cpu/load/average", Method: "GET", Name: IAAS_GET_CPU_LOAD_AVERAGE},
    {Path: "/v2/iaas/instance/disk/io/rate", Method: "GET", Name: IAAS_GET_DISK_IO_RATE},
    {Path: "/v2/iaas/instance/network/io/bytes", Method: "GET", Name: IAAS_GET_NETWORK_IO_BTYES},
}
