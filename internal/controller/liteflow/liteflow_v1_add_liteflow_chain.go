package liteflow

import (
	"context"
	"fmt"
	"time"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
	"go-liteflow-editor-client/internal/constant"
	"go-liteflow-editor-client/internal/dao"
	"go-liteflow-editor-client/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/google/uuid"
)

// AddLiteflowChain 新增LiteFlow链路
//
// 该方法用于创建新的LiteFlow执行链路，包含以下功能：
// 1. 输入参数验证（链名称和描述不能为空）
// 2. 生成唯一的UUID作为主键ID
// 3. 生成唯一的ChainId（使用时间戳+随机数）
// 4. 设置默认的ChainDsl和ElData
// 5. 检查链名称唯一性
// 6. 使用事务确保数据一致性
// 7. 返回完整的链路信息
//
// 参数:
//   - ctx: 上下文信息
//   - req: 新增链路请求参数，包含链名称和描述
//
// 返回值:
//   - res: 新增链路的响应信息，包含完整的链路数据
//   - err: 错误信息，如果成功则为nil
func (c *ControllerV1) AddLiteflowChain(ctx context.Context, req *v1.AddLiteflowChainReq) (res *v1.AddLiteflowChainRes, err error) {
	res = &v1.AddLiteflowChainRes{}

	// 输入验证
	if req.ChainName == "" {
		g.Log().Errorf(ctx, "AddLiteflowChain: 链名称不能为空")
		return res, fmt.Errorf("链名称不能为空")
	}
	if req.ChainDesc == "" {
		g.Log().Errorf(ctx, "AddLiteflowChain: 链描述不能为空")
		return res, fmt.Errorf("链描述不能为空")
	}

	// 生成UUID作为主键ID
	id := uuid.New().String()

	// 生成唯一的ChainId（使用时间戳+随机数确保唯一性）
	chainId := fmt.Sprintf("chain_%d_%s", time.Now().UnixNano(), uuid.New().String()[:8])

	// 设置默认的ChainDsl和ElData
	defaultChainDsl := `{"nodes": [], "edges": []}`
	defaultElData := ""

	chain := entity.LiteflowChain{
		Id:         id,
		ChainId:    chainId,
		ChainName:  req.ChainName,
		ChainDesc:  req.ChainDesc,
		ChainDsl:   defaultChainDsl,
		ElData:     defaultElData,
		Enable:     constant.StatusEnabled, // 默认启用
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}

	// 使用事务确保数据一致性
	err = dao.LiteflowChain.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 检查链名称是否已存在
		existingChain := &entity.LiteflowChain{}
		err := dao.LiteflowChain.Ctx(ctx).TX(tx).Where("chain_name", req.ChainName).Scan(existingChain)
		if err != nil {
			return fmt.Errorf("检查链名称唯一性失败: %v", err)
		}

		if existingChain.Id != "" {
			// 链名称已存在
			if existingChain.Enable == constant.StatusDisabled {
				// 如果Enable为0，则设置为1
				_, err = dao.LiteflowChain.Ctx(ctx).TX(tx).Data(g.Map{
					"enable":      constant.StatusEnabled,
					"update_time": gtime.Now(),
				}).Where("id", existingChain.Id).Update()
				if err != nil {
					return fmt.Errorf("更新链路状态失败: %v", err)
				}

				// 更新existingChain的Enable字段，用于后续返回
				existingChain.Enable = constant.StatusEnabled
				existingChain.UpdateTime = gtime.Now()

				// 将existingChain赋值给chain，用于后续返回
				chain = *existingChain

				g.Log().Infof(ctx, "AddLiteflowChain: 链名称已存在且已禁用，已重新启用，ID: %s", existingChain.Id)
				return nil
			} else {
				return fmt.Errorf("链名称已存在且已启用")
			}
		}

		// 插入新链路
		_, err = dao.LiteflowChain.Ctx(ctx).TX(tx).Data(chain).Insert()
		if err != nil {
			return fmt.Errorf("新增链路失败: %v", err)
		}

		return nil
	})

	if err != nil {
		g.Log().Errorf(ctx, "AddLiteflowChain: %v", err)
		return res, err
	}

	// 设置返回数据
	res.Data.Id = chain.Id
	res.Data.ChainId = chain.ChainId
	res.Data.ChainName = chain.ChainName
	res.Data.ChainDesc = chain.ChainDesc
	res.Data.ChainDsl = chain.ChainDsl
	res.Data.ElData = chain.ElData
	res.Data.Enable = chain.Enable
	res.Data.CreateTime = chain.CreateTime.Format("2006-01-02 15:04:05")

	// 根据操作类型记录不同的日志
	if chain.Id == id {
		// 新创建的链路
		g.Log().Infof(ctx, "AddLiteflowChain: 成功新增链路，ID: %s, ChainId: %s, ChainName: %s", id, chainId, req.ChainName)
	} else {
		// 重新启用的现有链路
		g.Log().Infof(ctx, "AddLiteflowChain: 成功重新启用现有链路，ID: %s, ChainId: %s, ChainName: %s", chain.Id, chain.ChainId, req.ChainName)
	}

	return res, nil
}
