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
* UpdateNode 更新节点
* 该方法用于更新现有的LiteFlow节点，包含以下功能：
* 1. 输入参数验证（ID和节点ID不能为空）
* 2. 检查节点是否存在
* 3. 验证节点ID唯一性（如果修改了节点ID）
* 4. 使用事务确保数据一致性
* 5. 更新节点信息并返回完整的节点数据
* 参数:
*   - ctx: 上下文信息
*   - req: 更新节点请求参数，包含要更新的节点信息
* 返回值:
*   - res: 更新后节点的响应信息，包含完整的节点数据
*   - err: 错误信息，如果成功则为nil
 */

func (c *ControllerV1) UpdateNode(ctx context.Context, req *v1.UpdateNodeReq) (res *v1.UpdateNodeRes, err error) {
	res = &v1.UpdateNodeRes{}

	// 输入验证
	if req.Id == "" {
		g.Log().Errorf(ctx, "UpdateNode: ID不能为空")
		return res, fmt.Errorf("ID不能为空")
	}
	if req.NodeId == "" {
		g.Log().Errorf(ctx, "UpdateNode: 节点ID不能为空")
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

		// 如果修改了节点ID，检查新节点ID是否与其他节点重复
		if req.NodeId != "" && req.NodeId != existingNode.NodeId {
			count, err := dao.LiteflowNode.Ctx(ctx).TX(tx).Where("node_id", req.NodeId).WhereNot("id", req.Id).Count()
			if err != nil {
				return fmt.Errorf("检查节点ID唯一性失败: %v", err)
			}
			if count > 0 {
				return fmt.Errorf("节点ID已存在")
			}
		}

		// 构建更新数据
		updateData := g.Map{}

		// 只更新非空字段
		if req.NodeId != "" {
			updateData["node_id"] = req.NodeId
		}
		if req.NodeName != "" {
			updateData["node_name"] = req.NodeName
		}
		if req.NodeType != "" {
			updateData["node_type"] = req.NodeType
		}
		if req.ClassName != "" {
			updateData["class_name"] = req.ClassName
		}
		if req.ScriptId != "" {
			updateData["script_id"] = req.ScriptId
		}
		if req.NodeDesc != "" {
			updateData["node_desc"] = req.NodeDesc
		}

		// 启用状态验证和设置
		if req.Enable == constant.StatusEnabled || req.Enable == constant.StatusDisabled {
			updateData["enable"] = req.Enable
		}

		// 设置更新时间
		updateData["update_time"] = gtime.Now()

		// 执行更新
		_, err = dao.LiteflowNode.Ctx(ctx).TX(tx).Data(updateData).Where("id", req.Id).Update()
		if err != nil {
			return fmt.Errorf("更新节点失败: %v", err)
		}

		// 查询更新后的完整节点信息
		updatedNode := &entity.LiteflowNode{}
		err = dao.LiteflowNode.Ctx(ctx).TX(tx).Where("id", req.Id).Scan(updatedNode)
		if err != nil {
			return fmt.Errorf("查询更新后的节点信息失败: %v", err)
		}

		// 设置返回数据
		res.Data.Id = updatedNode.Id
		res.Data.NodeId = updatedNode.NodeId
		res.Data.NodeName = updatedNode.NodeName
		res.Data.NodeType = updatedNode.NodeType
		res.Data.ClassName = updatedNode.ClassName
		res.Data.ScriptId = updatedNode.ScriptId
		res.Data.NodeDesc = updatedNode.NodeDesc
		res.Data.Enable = updatedNode.Enable
		res.Data.CreateTime = updatedNode.CreateTime.String()
		res.Data.UpdateTime = updatedNode.UpdateTime.String()

		return nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "UpdateNode: %v", err)
		return res, err
	}

	g.Log().Infof(ctx, "UpdateNode: 成功更新节点，ID: %s, 节点ID: %s", res.Data.Id, res.Data.NodeId)
	return res, nil
}
