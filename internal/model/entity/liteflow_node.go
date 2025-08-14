// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowNode is the golang structure for table liteflow_node.
type LiteflowNode struct {
	Id         string      `json:"id"         orm:"id"          ` // 主键ID(UUID)
	NodeId     string      `json:"nodeId"     orm:"node_id"     ` // 节点ID
	NodeName   string      `json:"nodeName"   orm:"node_name"   ` // 节点名称
	NodeType   string      `json:"nodeType"   orm:"node_type"   ` // 节点类型(common,switch,for,while等)
	ClassName  string      `json:"className"  orm:"class_name"  ` // 节点实现类名
	ScriptId   string      `json:"scriptId"   orm:"script_id"   ` // 关联脚本ID
	NodeDesc   string      `json:"nodeDesc"   orm:"node_desc"   ` // 节点描述
	Enable     int         `json:"enable"     orm:"enable"      ` // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 更新时间
}
