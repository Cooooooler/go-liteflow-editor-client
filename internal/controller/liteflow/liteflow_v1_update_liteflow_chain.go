package liteflow

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
	"go-liteflow-editor-client/internal/constant"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"
)

/*
* UpdateLiteflowChain 更新LiteFlow链路
* 该方法用于更新现有的LiteFlow执行链路，包含以下功能：
* 1. 输入参数验证（ID和ChainId不能为空）
* 2. 检查链路是否存在
* 3. 验证链名称唯一性（如果修改了链名称）
* 4. 使用事务确保数据一致性
* 5. 更新链路信息并返回完整的链路数据
* 参数:
*   - ctx: 上下文信息
*   - req: 更新链路请求参数，包含要更新的链路信息
* 返回值:
*   - res: 更新后链路的响应信息，包含完整的链路数据
*   - err: 错误信息，如果成功则为nil
 */

func (c *ControllerV1) UpdateLiteflowChain(ctx context.Context, req *v1.UpdateLiteflowChainReq) (res *v1.UpdateLiteflowChainRes, err error) {
	res = &v1.UpdateLiteflowChainRes{}

	// 输入验证
	if req.Id == "" {
		g.Log().Errorf(ctx, "UpdateLiteflowChain: ID不能为空")
		return res, gerror.NewCode(gcode.CodeInvalidParameter, "ID不能为空")
	}
	if req.ChainId == "" {
		g.Log().Errorf(ctx, "UpdateLiteflowChain: 链路ID不能为空")
		return res, gerror.NewCode(gcode.CodeInvalidParameter, "链路ID不能为空")
	}

	// 使用事务确保数据一致性
	err = dao.LiteflowChain.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 检查链路是否存在
		existingChain := &entity.LiteflowChain{}
		err := dao.LiteflowChain.Ctx(ctx).TX(tx).Where("id", req.Id).Scan(existingChain)
		if err != nil {
			return fmt.Errorf("查询链路失败: %v", err)
		}
		if existingChain.Id == "" {
			return fmt.Errorf("链路不存在")
		}

		// 如果修改了链名称，检查新名称是否与其他链路重复
		if req.ChainName != "" && req.ChainName != existingChain.ChainName {
			count, err := dao.LiteflowChain.Ctx(ctx).TX(tx).Where("chain_name", req.ChainName).WhereNot("id", req.Id).Count()
			if err != nil {
				return fmt.Errorf("检查链名称唯一性失败: %v", err)
			}
			if count > 0 {
				return fmt.Errorf("链名称已存在")
			}
		}

		// 构建更新数据
		updateData := g.Map{}

		// 只更新非空字段
		if req.ChainName != "" {
			updateData["chain_name"] = req.ChainName
		}
		if req.ChainDesc != "" {
			updateData["chain_desc"] = req.ChainDesc
		}
		if req.ChainDsl != "" {
			updateData["chain_dsl"] = req.ChainDsl
		}
		if req.ElData != "" {
			updateData["el_data"] = req.ElData
		}

		// 启用状态验证和设置
		if req.Enable == constant.StatusEnabled || req.Enable == constant.StatusDisabled {
			updateData["enable"] = req.Enable
		}

		// 设置更新时间
		updateData["update_time"] = gtime.Now()

		// 执行更新
		_, err = dao.LiteflowChain.Ctx(ctx).TX(tx).Data(updateData).Where("id", req.Id).Update()
		if err != nil {
			return fmt.Errorf("更新链路失败: %v", err)
		}

		return nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "UpdateLiteflowChain: %v", err)
		return res, err
	}

	// 查询更新后的完整数据
	var updatedChain entity.LiteflowChain
	err = dao.LiteflowChain.Ctx(ctx).Where("id", req.Id).Scan(&updatedChain)
	if err != nil {
		g.Log().Errorf(ctx, "UpdateLiteflowChain: 查询更新后数据失败: %v", err)
		return res, fmt.Errorf("查询更新后数据失败: %v", err)
	}

	// 设置返回数据
	res.Data.Id = updatedChain.Id
	res.Data.ChainId = updatedChain.ChainId
	res.Data.ChainName = updatedChain.ChainName
	res.Data.ChainDesc = updatedChain.ChainDesc
	res.Data.ChainDsl = updatedChain.ChainDsl
	res.Data.ElData = updatedChain.ElData
	res.Data.Enable = updatedChain.Enable

	// 格式化时间
	if updatedChain.CreateTime != nil {
		res.Data.CreateTime = updatedChain.CreateTime.Format("2006-01-02 15:04:05")
	}
	if updatedChain.UpdateTime != nil {
		res.Data.UpdateTime = updatedChain.UpdateTime.Format("2006-01-02 15:04:05")
	}

	// 记录成功日志
	g.Log().Infof(ctx, "UpdateLiteflowChain: 链路更新成功，ID: %s, ChainId: %s", req.Id, req.ChainId)

	return res, nil
}
