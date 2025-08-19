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
		go s.cleanupDisabledChains(ctx)
	} else {
		g.Log().Info(ctx, "配置为启动时不立即执行清理任务")
	}

	// 使用配置的cron表达式执行定期清理任务
	_, err := gcron.Add(ctx, config.CleanupCronExpression, s.cleanupDisabledChains)

	if err != nil {
		g.Log().Errorf(ctx, "启动清理定时任务失败: %v", err)
	} else {
		g.Log().Infof(ctx, "成功启动清理定时任务，cron表达式: %s", config.CleanupCronExpression)
	}
}

// cleanupDisabledChains 清理已禁用的链路数据
func (s *CronService) cleanupDisabledChains(ctx context.Context) {
	g.Log().Info(ctx, "开始执行定时清理任务：删除enable字段为0的链路数据")

	startTime := time.Now()

	// 查询所有enable为0的记录
	count, err := dao.LiteflowChain.Ctx(ctx).Where("enable", 0).Count()
	if err != nil {
		g.Log().Errorf(ctx, "查询待清理记录数量失败: %v", err)
		return
	}

	if count == 0 {
		g.Log().Info(ctx, "没有需要清理的记录")
		return
	}

	g.Log().Infof(ctx, "发现 %d 条待清理记录", count)

	// 执行物理删除
	result, err := dao.LiteflowChain.Ctx(ctx).Where("enable", 0).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "删除已禁用链路数据失败: %v", err)
		return
	}

	affectedRows, _ := result.RowsAffected()
	duration := time.Since(startTime)

	g.Log().Infof(ctx, "定时清理任务完成，成功删除 %d 条记录，耗时: %v", affectedRows, duration)
}

// ManualCleanup 手动执行清理任务（用于测试或紧急清理）
func (s *CronService) ManualCleanup() error {
	ctx := context.Background()
	g.Log().Info(ctx, "手动执行清理任务")

	s.cleanupDisabledChains(ctx)
	return nil
}
