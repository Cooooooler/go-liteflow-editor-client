package node

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "go-liteflow-editor-client/api/node/v1"
	"go-liteflow-editor-client/internal/constant"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
)

/*
* DeleteNode 删除节点
* 该方法用于删除LiteFlow节点，包含以下功能：
* 1. 输入参数验证（ID和节点ID不能为空）
* 2. 检查节点是否存在
* 3. 执行软删除（设置enable为0）
* 4. 使用事务确保数据一致性
* 5. 返回删除结果信息
* 参数:
*   - ctx: 上下文信息
*   - req: 删除节点请求参数，包含要删除的节点ID信息
* 返回值:
*   - res: 删除节点的响应信息，包含删除结果数据
*   - err: 错误信息，如果成功则为nil
 */

func (c *ControllerV1) DeleteNode(ctx context.Context, req *v1.DeleteNodeReq) (res *v1.DeleteNodeRes, err error) {
	res = &v1.DeleteNodeRes{}

	// 输入验证
	if req.Id == "" {
		g.Log().Errorf(ctx, "DeleteNode: ID不能为空")
		return res, fmt.Errorf("ID不能为空")
	}
	if req.NodeId == "" {
		g.Log().Errorf(ctx, "DeleteNode: 节点ID不能为空")
		return res, fmt.Errorf("节点ID不能为空")
	}

	// 使用事务确保数据一致性
	err = dao.LiteflowNode.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 检查节点是否存在
		existingNode := &entity.LiteflowNode{}
		err := dao.LiteflowNode.Ctx(ctx).TX(tx).Where("id", req.Id).Scan(existingNode)
		if err != nil {
			return fmt.Errorf("查询节点失败: %v", err)
		}
		if existingNode.Id == "" {
			return fmt.Errorf("节点不存在")
		}

		// 验证节点ID是否匹配
		if existingNode.NodeId != req.NodeId {
			return fmt.Errorf("节点ID不匹配")
		}

		// 检查节点是否已经被删除
		if existingNode.Enable == constant.StatusDisabled {
			return fmt.Errorf("节点已经被删除")
		}

		// 执行软删除（设置enable为0）
		_, err = dao.LiteflowNode.Ctx(ctx).TX(tx).Data(g.Map{
			"enable":      constant.StatusDisabled,
			"update_time": gtime.Now(),
		}).Where("id", req.Id).Update()
		if err != nil {
			return fmt.Errorf("删除节点失败: %v", err)
		}

		// 更新existingNode的Enable字段，用于后续返回
		existingNode.Enable = constant.StatusDisabled
		existingNode.UpdateTime = gtime.Now()

		// 设置返回数据
		res.Data.Id = existingNode.Id
		res.Data.NodeId = existingNode.NodeId

		return nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "DeleteNode: %v", err)
		return res, err
	}

	g.Log().Infof(ctx, "DeleteNode: 成功删除节点，ID: %s, 节点ID: %s", res.Data.Id, res.Data.NodeId)
	return res, nil
}
