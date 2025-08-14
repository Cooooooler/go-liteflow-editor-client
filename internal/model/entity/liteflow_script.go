// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// LiteflowScript is the golang structure for table liteflow_script.
type LiteflowScript struct {
	Id         string      `json:"id"         orm:"id"          ` // 主键ID(UUID)
	ScriptId   string      `json:"scriptId"   orm:"script_id"   ` // 脚本ID
	ScriptName string      `json:"scriptName" orm:"script_name" ` // 脚本名称
	ScriptType string      `json:"scriptType" orm:"script_type" ` // 脚本类型(groovy,js,python等)
	ScriptData string      `json:"scriptData" orm:"script_data" ` // 脚本内容
	ScriptDesc string      `json:"scriptDesc" orm:"script_desc" ` // 脚本描述
	Enable     int         `json:"enable"     orm:"enable"      ` // 是否启用 1:启用 0:禁用
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" ` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" ` // 更新时间
}
