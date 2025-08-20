package node

import (
	"context"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
	"go-liteflow-editor-client/utility"

	"github.com/gogf/gf/v2/frame/g"

	v1 "go-liteflow-editor-client/api/node/v1"
)

func (c *ControllerV1) GetNode(ctx context.Context, req *v1.GetNodeReq) (res *v1.GetNodeRes, err error) {
	res = &v1.GetNodeRes{}

	// 使用分页工具包验证和修复分页参数
	paginationParams := utility.ValidateAndFixPagination(
		&utility.PaginationParams{
			Page:     req.Page,
			PageSize: req.PageSize,
		},
		nil, // 使用默认选项
	)

	// 构建查询模型
	model := dao.LiteflowNode.Ctx(ctx).Where("enable", 1) // 只查询启用的节点

	// 添加搜索条件
	if req.NodeType != "" {
		model = model.Where("node_type = ?", req.NodeType)
	}
	if req.SearchKey != "" {
		model = model.Where("node_name LIKE ? OR node_desc LIKE ?", "%"+req.SearchKey+"%", "%"+req.SearchKey+"%")
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		g.Log().Errorf(ctx, "GetNode: 查询总数失败: %v", err)
		return res, nil
	}

	// 如果没有数据，直接返回空结果
	if total == 0 {
		g.Log().Infof(ctx, "GetNode: 没有找到符合条件的链路数据")
		// 使用分页工具包计算分页信息
		paginationInfo := utility.CalculatePaginationInfo(0, paginationParams)
		res.PageInfo.CurrentPage = paginationInfo.CurrentPage
		res.PageInfo.PageSize = paginationInfo.PageSize
		res.PageInfo.Total = paginationInfo.Total
		res.PageInfo.TotalPage = paginationInfo.TotalPage
		res.Data = []v1.NodeTag{} // 确保返回空数组而不是null
		return res, nil
	}

	// 使用分页工具包计算偏移量
	offset := utility.CalculateOffset(paginationParams.Page, paginationParams.PageSize)

	// 分页查询
	var nodes []entity.LiteflowNode
	err = model.Page(offset, paginationParams.PageSize).Order("create_time DESC").Scan(&nodes)
	if err != nil {
		g.Log().Errorf(ctx, "GetNode: 查询数据失败: %v", err)
		return res, nil
	}

	// 转换数据格式
	var nodeList []v1.NodeTag
	for _, node := range nodes {
		nodeItem := v1.NodeTag{
			Id:         node.Id,
			NodeId:     node.NodeId,
			NodeName:   node.NodeName,
			NodeType:   node.NodeType,
			ClassName:  node.ClassName,
			ScriptId:   node.ScriptId,
			NodeDesc:   node.NodeDesc,
			Enable:     node.Enable,
			CreateTime: "",
			UpdateTime: "",
		}

		// 格式化时间
		if node.CreateTime != nil {
			nodeItem.CreateTime = node.CreateTime.Format("2006-01-02 15:04:05")
		}
		if node.UpdateTime != nil {
			nodeItem.UpdateTime = node.UpdateTime.Format("2006-01-02 15:04:05")
		}

		nodeList = append(nodeList, nodeItem)
	}

	res.Data = nodeList

	// 使用分页工具包计算分页信息
	paginationInfo := utility.CalculatePaginationInfo(total, paginationParams)
	res.PageInfo.CurrentPage = paginationInfo.CurrentPage
	res.PageInfo.PageSize = paginationInfo.PageSize
	res.PageInfo.Total = paginationInfo.Total
	res.PageInfo.TotalPage = paginationInfo.TotalPage

	// 记录成功日志
	g.Log().Infof(ctx, "GetNode: 查询成功，总数: %d, 当前页: %d, 每页大小: %d, 返回数据条数: %d",
		total, paginationParams.Page, paginationParams.PageSize, len(nodeList))

	return res, nil
}
