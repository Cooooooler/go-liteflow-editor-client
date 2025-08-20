package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateNodeReq struct {
	g.Meta    `path:"/api/updateNode" tags:"Node" method:"post" summary:"更新节点"`
	Id        string `json:"id" v:"required#ID不能为空" dc:"主键ID"`
	NodeId    string `json:"nodeId" v:"required#节点ID不能为空" dc:"节点ID"`
	NodeName  string `json:"nodeName" dc:"节点名称"`
	NodeType  string `json:"nodeType" dc:"节点类型"`
	ClassName string `json:"className" dc:"类名"`
	ScriptId  string `json:"scriptId" dc:"脚本ID"`
	NodeDesc  string `json:"nodeDesc" dc:"节点描述"`
	Enable    int    `json:"enable" dc:"是否启用"`
}

type UpdateNodeData struct {
	Id         string `json:"id" dc:"主键ID"`
	NodeId     string `json:"nodeId" dc:"节点ID"`
	NodeName   string `json:"nodeName" dc:"节点名称"`
	NodeType   string `json:"nodeType" dc:"节点类型"`
	ClassName  string `json:"className" dc:"类名"`
	ScriptId   string `json:"scriptId" dc:"脚本ID"`
	NodeDesc   string `json:"nodeDesc" dc:"节点描述"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	UpdateTime string `json:"updateTime" dc:"更新时间"`
}

type UpdateNodeRes struct {
	g.Meta `mime:"application/json"`
	Data   UpdateNodeData `json:"data" dc:"返回数据"`
}
