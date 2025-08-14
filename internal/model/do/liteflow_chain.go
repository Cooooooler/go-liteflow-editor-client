// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowChain is the golang structure of table liteflow_chain for DAO operations like Where/Data.
type LiteflowChain struct {
	g.Meta     `orm:"table:liteflow_chain, do:true"`
	Id         interface{} // 主键ID(UUID)
	ChainId    interface{} // 链路ID
	ChainName  interface{} // 链路名称
	ChainDesc  interface{} // 链路描述
	ElData     interface{} // EL表达式数据
	Enable     interface{} // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
