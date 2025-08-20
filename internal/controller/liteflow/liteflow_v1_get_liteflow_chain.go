package liteflow

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
	"go-liteflow-editor-client/utility"
)

func (c *ControllerV1) GetLiteflowChain(ctx context.Context, req *v1.GetLiteflowChainReq) (res *v1.GetLiteflowChainRes, err error) {
	res = &v1.GetLiteflowChainRes{}

	// 使用分页工具包验证和修复分页参数
	paginationParams := utility.ValidateAndFixPagination(
		&utility.PaginationParams{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		nil, // 使用默认选项
	)

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
		return res, nil
	}

	// 如果没有数据，直接返回空结果
	if total == 0 {
		g.Log().Infof(ctx, "GetLiteflowChain: 没有找到符合条件的链路数据")
		// 使用分页工具包计算分页信息
		paginationInfo := utility.CalculatePaginationInfo(0, paginationParams)
		res.PageInfo.CurrentPage = paginationInfo.CurrentPage
		res.PageInfo.PageSize = paginationInfo.PageSize
		res.PageInfo.Total = paginationInfo.Total
		res.PageInfo.TotalPage = paginationInfo.TotalPage
		res.Data = []v1.Chain{}
		return res, nil
	}

	// 使用分页工具包计算偏移量
	offset := utility.CalculateOffset(paginationParams.Page, paginationParams.PageSize)

	// 分页查询
	var chains []entity.LiteflowChain
	err = model.Page(offset, paginationParams.PageSize).Order("create_time DESC").Scan(&chains)
	if err != nil {
		g.Log().Errorf(ctx, "GetLiteflowChain: 查询数据失败: %v", err)
		return res, nil
	}

	// 转换数据格式
	var chainList []v1.Chain
	for _, chain := range chains {
		chainItem := v1.Chain{
			Id:         chain.Id,
			ChainId:    chain.ChainId,
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

	// 使用分页工具包计算分页信息
	paginationInfo := utility.CalculatePaginationInfo(total, paginationParams)
	res.PageInfo.CurrentPage = paginationInfo.CurrentPage
	res.PageInfo.PageSize = paginationInfo.PageSize
	res.PageInfo.Total = paginationInfo.Total
	res.PageInfo.TotalPage = paginationInfo.TotalPage

	// 记录成功日志
	g.Log().Infof(ctx, "GetLiteflowChain: 查询成功，总数: %d, 当前页: %d, 每页大小: %d, 返回数据条数: %d",
		total, paginationParams.Page, paginationParams.PageSize, len(chainList))

	return res, nil
}
