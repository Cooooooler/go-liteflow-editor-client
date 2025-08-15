package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetLiteflowChainReq struct {
	g.Meta    `path:"/api/getLiteflowChain" tags:"Liteflow" method:"get" summary:"获取liteflow链"`
	SearchKey string `json:"searchKey" dc:"搜索关键字"`
	Page      int    `json:"page" dc:"页码"`
	PageSize  int    `json:"pageSize" dc:"每页条数"`
}

type Chain struct {
	Id         string `json:"id" dc:"id"`
	ChainName  string `json:"chainName" dc:"链名称"`
	ChainDesc  string `json:"chainDesc" dc:"链描述"`
	ChainDsl   string `json:"chainDsl" dc:"链dsl"`
	ElData     string `json:"elData" dc:"el数据"`
	Enable     int    `json:"enable" dc:"是否启用"`
	CreateTime string `json:"createTime" dc:"创建时间"`
	UpdateTime string `json:"updateTime" dc:"更新时间"`
}

type GetLiteflowChainRes struct {
	g.Meta   `mime:"application/json"`
	Data     []Chain `json:"data" dc:"返回数据"`
	PageInfo struct {
		CurrentPage int `json:"currentPage" dc:"当前页码"`
		PageSize    int `json:"pageSize" dc:"每页大小"`
		Total       int `json:"total" dc:"总记录数"`
		TotalPage   int `json:"totalPage" dc:"总页数"`
	} `json:"pageInfo" dc:"分页信息"`
}
