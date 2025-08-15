package liteflow

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
)

// DeleteLiteflowChain 删除LiteFlow链路
//
// 该方法用于删除指定的LiteFlow执行链路，包含以下功能：
// 1. 输入参数验证（ID和ChainId不能为空）
// 2. 检查链路是否存在
// 3. 执行软删除（将enable字段设置为0）
// 4. 返回删除结果
//
// 参数:
//   - ctx: 上下文信息
//   - req: 删除链路请求参数，包含ID和ChainId
//
// 返回值:
//   - res: 删除链路的响应信息，包含被删除的链路数据
//   - err: 错误信息，如果成功则为nil
func (c *ControllerV1) DeleteLiteflowChain(ctx context.Context, req *v1.DeleteLiteflowChainReq) (res *v1.DeleteLiteflowChainRes, err error) {
	res = &v1.DeleteLiteflowChainRes{}

	// 输入验证
	if req.Id == "" {
		g.Log().Errorf(ctx, "DeleteLiteflowChain: ID不能为空")
		return res, fmt.Errorf("ID不能为空")
	}
	if req.ChainId == "" {
		g.Log().Errorf(ctx, "DeleteLiteflowChain: ChainId不能为空")
		return res, fmt.Errorf("ChainId不能为空")
	}

	// 检查链路是否存在
	var chain entity.LiteflowChain
	err = dao.LiteflowChain.Ctx(ctx).Where("id", req.Id).Where("chain_id", req.ChainId).Scan(&chain)
	if err != nil {
		g.Log().Errorf(ctx, "DeleteLiteflowChain: 查询链路失败: %v", err)
		return res, fmt.Errorf("查询链路失败: %v", err)
	}

	// 检查是否找到记录
	if chain.Id == "" {
		g.Log().Errorf(ctx, "DeleteLiteflowChain: 链路不存在，ID: %s, ChainId: %s", req.Id, req.ChainId)
		return res, fmt.Errorf("链路不存在")
	}

	// 检查链路是否已经被删除（软删除）
	if chain.Enable == 0 {
		g.Log().Infof(ctx, "DeleteLiteflowChain: 链路已经被删除，ID: %s, ChainId: %s", req.Id, req.ChainId)
		return res, fmt.Errorf("链路已经被删除")
	}

	// 执行软删除（将enable字段设置为0）
	_, err = dao.LiteflowChain.Ctx(ctx).Data("enable", 0).Where("id", req.Id).Where("chain_id", req.ChainId).Update()
	if err != nil {
		g.Log().Errorf(ctx, "DeleteLiteflowChain: 删除链路失败: %v", err)
		return res, fmt.Errorf("删除链路失败: %v", err)
	}

	// 设置返回数据
	res.Data.Id = chain.Id
	res.Data.ChainId = chain.ChainId

	g.Log().Infof(ctx, "DeleteLiteflowChain: 成功删除链路，ID: %s, ChainId: %s, ChainName: %s",
		chain.Id, chain.ChainId, chain.ChainName)

	return res, nil
}
