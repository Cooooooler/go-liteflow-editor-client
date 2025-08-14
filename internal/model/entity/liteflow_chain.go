// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowChain is the golang structure for table liteflow_chain.
type LiteflowChain struct {
	Id         string      `json:"id"         orm:"id"          ` // 主键ID(UUID)
	ChainId    string      `json:"chainId"    orm:"chain_id"    ` // 链路ID
	ChainName  string      `json:"chainName"  orm:"chain_name"  ` // 链路名称
	ChainDesc  string      `json:"chainDesc"  orm:"chain_desc"  ` // 链路描述
	ElData     string      `json:"elData"     orm:"el_data"     ` // EL表达式数据
	Enable     int         `json:"enable"     orm:"enable"      ` // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 更新时间
}
