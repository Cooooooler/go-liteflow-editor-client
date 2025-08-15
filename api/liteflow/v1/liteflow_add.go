package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddLiteflowChainReq struct {
	g.Meta    `path:"/api/addLiteflowChain" tags:"Liteflow" method:"post" summary:"新增liteflow链"`
	ChainName string `json:"chainName" v:"required#链名称不能为空" dc:"链名称"`
	ChainDesc string `json:"chainDesc" v:"required#链描述不能为空" dc:"链描述"`
}

type AddLiteflowChainRes struct {
	g.Meta `mime:"application/json"`
	Data   interface{} `json:"data" dc:"返回数据"`
}
