package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteLiteflowChainReq struct {
	g.Meta  `path:"/api/deleteLiteflowChain" tags:"Liteflow" method:"post" summary:"删除liteflow链"`
	Id      string `json:"id" v:"required#id不能为空" dc:"id"`
	ChainId string `json:"chainId" v:"required#链id不能为空" dc:"链id"`
}

type DeleteLiteflowChainData struct {
	Id      string `json:"id" dc:"主键ID"`
	ChainId string `json:"chainId" dc:"链id"`
}

type DeleteLiteflowChainRes struct {
	g.Meta `mime:"application/json"`
	Data   DeleteLiteflowChainData `json:"data" dc:"返回数据"`
}
