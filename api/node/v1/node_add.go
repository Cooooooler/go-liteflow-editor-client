package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddNodeReq struct {
	g.Meta    `path:"/api/addNode" tags:"Node" method:"post" summary:"新增节点"`
	NodeId    string `json:"nodeId" v:"required#节点ID不能为空" dc:"节点ID"`
	NodeName  string `json:"nodeName" v:"required#节点名称不能为空" dc:"节点名称"`
	NodeType  string `json:"nodeType" v:"required#节点类型不能为空" dc:"节点类型"`
	ClassName string `json:"className" dc:"类名"`
	ScriptId  string `json:"scriptId" dc:"脚本ID"`
	NodeDesc  string `json:"nodeDesc" dc:"节点描述"`
	Enable    int    `json:"enable" dc:"是否启用"`
}

type AddNodeData struct {
	Id         string `json:"id" dc:"主键ID"`
	NodeId     string `json:"nodeId" dc:"节点ID"`
	NodeName   string `json:"nodeName" dc:"节点名称"`
	NodeType   string `json:"nodeType" dc:"节点类型"`
	ClassName  string `json:"className" dc:"类名"`
	ScriptId   string `json:"scriptId" dc:"脚本ID"`
	NodeDesc   string `json:"nodeDesc" dc:"节点描述"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

type AddNodeRes struct {
	g.Meta `mime:"application/json"`
	Data   AddNodeData `json:"data" dc:"返回数据"`
}
