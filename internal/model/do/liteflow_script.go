// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowScript is the golang structure of table liteflow_script for DAO operations like Where/Data.
type LiteflowScript struct {
	g.Meta     `orm:"table:liteflow_script, do:true"`
	Id         interface{} // 主键ID(UUID)
	ScriptId   interface{} // 脚本ID
	ScriptName interface{} // 脚本名称
	ScriptType interface{} // 脚本类型(groovy,js,python等)
	ScriptData interface{} // 脚本内容
	ScriptDesc interface{} // 脚本描述
	Enable     interface{} // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
}
