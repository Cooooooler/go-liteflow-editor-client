package service

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"

	"go-liteflow-editor-client/internal/dao"
)

// CronService 定时任务服务
type CronService struct{}

// NewCronService 创建定时任务服务实例
func NewCronService() *CronService {
	return &CronService{}
}

// StartCleanupTask 启动清理任务
func (s *CronService) StartCleanupTask() {
	ctx := context.Background()
	config := GetCronConfig(ctx)

	if !config.EnableCleanupTask {
		g.Log().Info(ctx, "定时清理任务已禁用")
		return
	}

	// 根据配置决定是否在启动时立即执行一次清理任务
	if config.RunOnStartup {
		g.Log().Info(ctx, "启动时立即执行一次清理任务")
		go s.cleanupDisabledData(ctx)
	} else {
		g.Log().Info(ctx, "配置为启动时不立即执行清理任务")
	}

	// 使用配置的cron表达式执行定期清理任务
	_, err := gcron.Add(ctx, config.CleanupCronExpression, s.cleanupDisabledData)

	if err != nil {
		g.Log().Errorf(ctx, "启动清理定时任务失败: %v", err)
	} else {
		g.Log().Infof(ctx, "成功启动清理定时任务，cron表达式: %s", config.CleanupCronExpression)
	}
}

// cleanupDisabledData 清理已禁用的链路和节点数据
func (s *CronService) cleanupDisabledData(ctx context.Context) {
	g.Log().Info(ctx, "开始执行定时清理任务：删除enable字段为0的链路和节点数据")

	startTime := time.Now()
	var totalDeleted int64

	// 清理已禁用的链路数据
	chainCount, err := dao.LiteflowChain.Ctx(ctx).Where("enable", 0).Count()
	if err != nil {
		g.Log().Errorf(ctx, "查询待清理链路记录数量失败: %v", err)
	} else if chainCount > 0 {
		g.Log().Infof(ctx, "发现 %d 条待清理的链路记录", chainCount)

		// 执行物理删除
		result, err := dao.LiteflowChain.Ctx(ctx).Where("enable", 0).Delete()
		if err != nil {
			g.Log().Errorf(ctx, "删除已禁用链路数据失败: %v", err)
		} else {
			affectedRows, _ := result.RowsAffected()
			totalDeleted += affectedRows
			g.Log().Infof(ctx, "成功删除 %d 条已禁用的链路记录", affectedRows)
		}
	} else {
		g.Log().Info(ctx, "没有需要清理的链路记录")
	}

	// 清理已禁用的节点数据
	nodeCount, err := dao.LiteflowNode.Ctx(ctx).Where("enable", 0).Count()
	if err != nil {
		g.Log().Errorf(ctx, "查询待清理节点记录数量失败: %v", err)
	} else if nodeCount > 0 {
		g.Log().Infof(ctx, "发现 %d 条待清理的节点记录", nodeCount)

		// 执行物理删除
		result, err := dao.LiteflowNode.Ctx(ctx).Where("enable", 0).Delete()
		if err != nil {
			g.Log().Errorf(ctx, "删除已禁用节点数据失败: %v", err)
		} else {
			affectedRows, _ := result.RowsAffected()
			totalDeleted += affectedRows
			g.Log().Infof(ctx, "成功删除 %d 条已禁用的节点记录", affectedRows)
		}
	} else {
		g.Log().Info(ctx, "没有需要清理的节点记录")
	}

	duration := time.Since(startTime)

	if totalDeleted > 0 {
		g.Log().Infof(ctx, "定时清理任务完成，总共删除 %d 条记录，耗时: %v", totalDeleted, duration)
	} else {
		g.Log().Infof(ctx, "定时清理任务完成，没有需要清理的记录，耗时: %v", duration)
	}
}

// ManualCleanup 手动执行清理任务（用于测试或紧急清理）
func (s *CronService) ManualCleanup() error {
	ctx := context.Background()
	g.Log().Info(ctx, "手动执行清理任务")

	s.cleanupDisabledData(ctx)
	return nil
}
