package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetNodeReq struct {
	g.Meta    `path:"/api/getNode" tags:"Node" method:"get" summary:"获取节点标签"`
	NodeType  string `json:"nodeType" dc:"节点类型"`
	SearchKey string `json:"searchKey" dc:"搜索关键字"`
	Page      int    `json:"page" dc:"页码"`
	PageSize  int    `json:"pageSize" dc:"每页条数"`
}

type NodeTag struct {
	Id         string `json:"id" dc:"id"`
	NodeId     string `json:"nodeId" dc:"节点id"`
	NodeName   string `json:"nodeName" dc:"节点名称"`
	NodeType   string `json:"nodeType" dc:"节点类型"`
	ClassName  string `json:"className" dc:"类名"`
	ScriptId   string `json:"scriptId" dc:"脚本id"`
	NodeDesc   string `json:"nodeDesc" dc:"节点描述"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	UpdateTime string `json:"updateTime" dc:"更新时间"`
}

type GetNodeRes struct {
	g.Meta   `mime:"application/json"`
	Data     []NodeTag `json:"data" dc:"返回数据"`
	PageInfo struct {
		CurrentPage int `json:"currentPage" dc:"当前页码"`
		PageSize    int `json:"pageSize" dc:"每页大小"`
		Total       int `json:"total" dc:"总记录数"`
		TotalPage   int `json:"totalPage" dc:"总页数"`
	} `json:"pageInfo" dc:"分页信息"`
}
