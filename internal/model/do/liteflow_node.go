// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowNode is the golang structure of table liteflow_node for DAO operations like Where/Data.
type LiteflowNode struct {
	g.Meta     `orm:"table:liteflow_node, do:true"`
	Id         interface{} // 主键ID(UUID)
	NodeId     interface{} // 节点ID
	NodeName   interface{} // 节点名称
	NodeType   interface{} // 节点类型(common,switch,for,while等)
	ClassName  interface{} // 节点实现类名
	ScriptId   interface{} // 关联脚本ID
	NodeDesc   interface{} // 节点描述
	Enable     interface{} // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
