package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type AddLiteflowChainReq struct {
	g.Meta    `path:"/api/addLiteflowChain" tags:"Liteflow" method:"post" summary:"新增liteflow链"`
	ChainName string `json:"chainName" v:"required#链名称不能为空" dc:"链名称"`
	ChainDesc string `json:"chainDesc" v:"required#链描述不能为空" dc:"链描述"`
}

type AddLiteflowChainData struct {
	Id         string `json:"id" dc:"主键ID"`
	ChainId    string `json:"chainId" dc:"链路ID"`
	ChainName  string `json:"chainName" dc:"链名称"`
	ChainDesc  string `json:"chainDesc" dc:"链描述"`
	ChainDsl   string `json:"chainDsl" dc:"链路dsl"`
	ElData     string `json:"elData" dc:"EL表达式数据"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
}

type AddLiteflowChainRes struct {
	g.Meta `mime:"application/json"`
	Data   AddLiteflowChainData `json:"data" dc:"返回数据"`
}
