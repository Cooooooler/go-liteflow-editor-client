// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowLog is the golang structure of table liteflow_log for DAO operations like Where/Data.
type LiteflowLog struct {
	g.Meta        `orm:"table:liteflow_log, do:true"`
	Id            interface{} // 主键ID(UUID)
	RequestId     interface{} // 请求ID
	ChainId       interface{} // 链路ID
	ChainName     interface{} // 链路名称
	NodeId        interface{} // 节点ID
	NodeName      interface{} // 节点名称
	ExecuteStatus interface{} // 执行状态(SUCCESS,FAILED,RUNNING)
	ExecuteTime   interface{} // 执行耗时(毫秒)
	ErrorMsg      interface{} // 错误信息
	InputData     interface{} // 输入数据
	OutputData    interface{} // 输出数据
	CreateTime    *gtime.Time // 创建时间
}
