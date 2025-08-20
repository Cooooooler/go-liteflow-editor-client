package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteNodeReq struct {
	g.Meta `path:"/api/deleteNode" tags:"Node" method:"post" summary:"删除节点"`
	Id     string `json:"id" v:"required#ID不能为空" dc:"主键ID"`
	NodeId string `json:"nodeId" v:"required#节点ID不能为空" dc:"节点ID"`
}

type DeleteNodeData struct {
	Id     string `json:"id" dc:"主键ID"`
	NodeId string `json:"nodeId" dc:"节点ID"`
}

type DeleteNodeRes struct {
	g.Meta `mime:"application/json"`
	Data   DeleteNodeData `json:"data" dc:"返回数据"`
}
