package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UpdateLiteflowChainReq struct {
	g.Meta    `path:"/api/updateLiteflowChain" tags:"Liteflow" method:"post" summary:"更新liteflow链"`
	Id        string `json:"id" v:"required#id不能为空" dc:"id"`
	ChainId   string `json:"chainId" v:"required#链路ID不能为空" dc:"链路ID"`
	ChainName string `json:"chainName" dc:"链名称"`
	ChainDesc string `json:"chainDesc" dc:"链描述"`
	ChainDsl  string `json:"chainDsl" dc:"链路dsl"`
	ElData    string `json:"elData" dc:"EL表达式数据"`
	Enable    int    `json:"enable" dc:"是否启用"`
}

type UpdateLiteflowChainData struct {
	Id         string `json:"id" dc:"主键ID"`
	ChainId    string `json:"chainId" dc:"链路ID"`
	ChainName  string `json:"chainName" dc:"链名称"`
	ChainDesc  string `json:"chainDesc" dc:"链描述"`
	ChainDsl   string `json:"chainDsl" dc:"链路dsl"`
	ElData     string `json:"elData" dc:"EL表达式数据"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	UpdateTime string `json:"updateTime" dc:"更新时间"`
}

type UpdateLiteflowChainRes struct {
	g.Meta `mime:"application/json"`
	Data   UpdateLiteflowChainData `json:"data" dc:"返回数据"`
}
