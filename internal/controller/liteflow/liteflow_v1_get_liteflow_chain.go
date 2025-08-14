package liteflow

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gpage"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
)

func (c *ControllerV1) GetLiteflowChain(ctx context.Context, req *v1.GetLiteflowChainReq) (res *v1.GetLiteflowChainRes, err error) {
	// 初始化响应
	res = &v1.GetLiteflowChainRes{
		Code:    200,
		Message: "success",
		Data:    []v1.Chain{},
	}

	// 设置默认分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	// 限制分页大小，防止查询过多数据
	if pageSize > 100 {
		pageSize = 100
	}

	// 构建查询模型
	model := dao.LiteflowChain.Ctx(ctx).Where("enable", 1) // 只查询启用的链路

	// 添加搜索条件
	if req.SearchKey != "" {
		model = model.Where("chain_name LIKE ? OR chain_desc LIKE ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		g.Log().Errorf(ctx, "GetLiteflowChain: 查询总数失败: %v", err)
		res.Code = 500
		res.Message = "查询总数失败"
		return res, nil
	}

	// 如果没有数据，直接返回空结果
	if total == 0 {
		g.Log().Infof(ctx, "GetLiteflowChain: 没有找到符合条件的链路数据")
		// 设置分页信息
		res.PageInfo.CurrentPage = page
		res.PageInfo.PageSize = pageSize
		res.PageInfo.Total = 0
		res.PageInfo.TotalPage = 0
		return res, nil
	}

	// 使用 gpage 创建分页对象
	pager := gpage.New(total, pageSize, page, "")

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 分页查询
	var chains []entity.LiteflowChain
	err = model.Page(offset, pageSize).Order("create_time DESC").Scan(&chains)
	if err != nil {
		g.Log().Errorf(ctx, "GetLiteflowChain: 查询数据失败: %v", err)
		res.Code = 500
		res.Message = "查询数据失败"
		return res, nil
	}

	// 转换数据格式
	var chainList []v1.Chain
	for _, chain := range chains {
		chainItem := v1.Chain{
			Id:         chain.Id,
			ChainName:  chain.ChainName,
			ChainDesc:  chain.ChainDesc,
			ChainDsl:   chain.ChainDsl,
			ElData:     chain.ElData,
			Enable:     chain.Enable,
			CreateTime: "",
			UpdateTime: "",
		}

		// 格式化时间
		if chain.CreateTime != nil {
			chainItem.CreateTime = chain.CreateTime.Format("2006-01-02 15:04:05")
		}
		if chain.UpdateTime != nil {
			chainItem.UpdateTime = chain.UpdateTime.Format("2006-01-02 15:04:05")
		}

		chainList = append(chainList, chainItem)
	}

	res.Data = chainList

	// 设置分页信息
	res.PageInfo.CurrentPage = pager.CurrentPage
	res.PageInfo.PageSize = pageSize
	res.PageInfo.Total = pager.TotalSize
	res.PageInfo.TotalPage = pager.TotalPage

	// 记录成功日志
	g.Log().Infof(ctx, "GetLiteflowChain: 查询成功，总数: %d, 当前页: %d, 每页大小: %d, 返回数据条数: %d",
		total, page, pageSize, len(chainList))

	return res, nil
}
