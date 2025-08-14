// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowLog is the golang structure for table liteflow_log.
type LiteflowLog struct {
	Id            string      `json:"id"            orm:"id"             ` // 主键ID(UUID)
	RequestId     string      `json:"requestId"     orm:"request_id"     ` // 请求ID
	ChainId       string      `json:"chainId"       orm:"chain_id"       ` // 链路ID
	ChainName     string      `json:"chainName"     orm:"chain_name"     ` // 链路名称
	NodeId        string      `json:"nodeId"        orm:"node_id"        ` // 节点ID
	NodeName      string      `json:"nodeName"      orm:"node_name"      ` // 节点名称
	ExecuteStatus string      `json:"executeStatus" orm:"execute_status" ` // 执行状态(SUCCESS,FAILED,RUNNING)
	ExecuteTime   int64       `json:"executeTime"   orm:"execute_time"   ` // 执行耗时(毫秒)
	ErrorMsg      string      `json:"errorMsg"      orm:"error_msg"      ` // 错误信息
	InputData     string      `json:"inputData"     orm:"input_data"     ` // 输入数据
	OutputData    string      `json:"outputData"    orm:"output_data"    ` // 输出数据
	CreateTime    *gtime.Time `json:"createTime"    orm:"create_time"    ` // 创建时间
}
