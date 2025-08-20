package node

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"

	v1 "go-liteflow-editor-client/api/node/v1"
	"go-liteflow-editor-client/internal/constant"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
)

/*
* AddNode 新增节点
* 该方法用于创建新的LiteFlow节点，包含以下功能：
* 1. 输入参数验证（节点ID、名称和类型不能为空）
* 2. 生成唯一的UUID作为主键ID
* 3. 检查节点ID唯一性
* 4. 设置默认的启用状态和创建时间
* 5. 使用事务确保数据一致性
* 6. 返回完整的节点信息
* 参数:
*   - ctx: 上下文信息
*   - req: 新增节点请求参数，包含节点基本信息
* 返回值:
*   - res: 新增节点的响应信息，包含完整的节点数据
*   - err: 错误信息，如果成功则为nil
 */

func (c *ControllerV1) AddNode(ctx context.Context, req *v1.AddNodeReq) (res *v1.AddNodeRes, err error) {
	res = &v1.AddNodeRes{}

	// 输入验证
	if req.NodeId == "" {
		g.Log().Errorf(ctx, "AddNode: 节点ID不能为空")
		return res, fmt.Errorf("节点ID不能为空")
	}
	if req.NodeName == "" {
		g.Log().Errorf(ctx, "AddNode: 节点名称不能为空")
		return res, fmt.Errorf("节点名称不能为空")
	}
	if req.NodeType == "" {
		g.Log().Errorf(ctx, "AddNode: 节点类型不能为空")
		return res, fmt.Errorf("节点类型不能为空")
	}

	// 生成UUID作为主键ID
	id := uuid.New().String()

	// 设置默认值
	if req.Enable == 0 {
		req.Enable = constant.StatusEnabled // 默认启用
	}

	// 创建节点实体
	node := entity.LiteflowNode{
		Id:         id,
		NodeId:     req.NodeId,
		NodeName:   req.NodeName,
		NodeType:   req.NodeType,
		ClassName:  req.ClassName,
		ScriptId:   req.ScriptId,
		NodeDesc:   req.NodeDesc,
		Enable:     req.Enable,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}

	// 使用事务确保数据一致性
	err = dao.LiteflowNode.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 检查节点ID是否已存在
		existingNode := &entity.LiteflowNode{}
		err := dao.LiteflowNode.Ctx(ctx).TX(tx).Where("node_id", req.NodeId).Scan(existingNode)
		if err != nil {
			// 如果没有找到记录，说明节点ID是唯一的，这是正常情况
			if err.Error() == "sql: no rows in result set" {
				// 节点ID唯一，可以继续创建新节点
			} else {
				// 其他数据库错误
				return fmt.Errorf("检查节点ID唯一性失败: %v", err)
			}
		}

		if existingNode.Id != "" {
			// 节点ID已存在
			if existingNode.Enable == constant.StatusDisabled {
				// 如果Enable为0，则设置为1
				_, err = dao.LiteflowNode.Ctx(ctx).TX(tx).Data(g.Map{
					"enable":      constant.StatusEnabled,
					"update_time": gtime.Now(),
				}).Where("id", existingNode.Id).Update()
				if err != nil {
					return fmt.Errorf("更新节点状态失败: %v", err)
				}

				// 更新existingNode的Enable字段，用于后续返回
				existingNode.Enable = constant.StatusEnabled
				existingNode.UpdateTime = gtime.Now()

				// 将existingNode赋值给node，用于后续返回
				node = *existingNode

				g.Log().Infof(ctx, "AddNode: 节点ID已存在且已禁用，已重新启用，ID: %s", existingNode.Id)
				return nil
			} else {
				return fmt.Errorf("节点ID已存在且已启用")
			}
		}

		// 插入新节点
		_, err = dao.LiteflowNode.Ctx(ctx).TX(tx).Data(node).Insert()
		if err != nil {
			return fmt.Errorf("新增节点失败: %v", err)
		}

		return nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "AddNode: %v", err)
		return res, err
	}

	// 设置返回数据
	res.Data.Id = node.Id
	res.Data.NodeId = node.NodeId
	res.Data.NodeName = node.NodeName
	res.Data.NodeType = node.NodeType
	res.Data.ClassName = node.ClassName
	res.Data.ScriptId = node.ScriptId
	res.Data.NodeDesc = node.NodeDesc
	res.Data.Enable = node.Enable
	res.Data.CreateTime = node.CreateTime.String()

	g.Log().Infof(ctx, "AddNode: 成功新增节点，ID: %s, 节点ID: %s", node.Id, node.NodeId)
	return res, nil
}
