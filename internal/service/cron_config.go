package service

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// CronConfig 定时任务配置
type CronConfig struct {
	// 清理任务的cron表达式，默认每天凌晨2点执行
	CleanupCronExpression string
	// 是否启用定时清理任务
	EnableCleanupTask bool
	// 是否在启动时立即执行一次清理任务
	RunOnStartup bool
	// 清理任务执行超时时间（秒）
	CleanupTimeout int
}

// GetCronConfig 获取定时任务配置
func GetCronConfig(ctx context.Context) *CronConfig {
	config := &CronConfig{
		CleanupCronExpression: "0 0 2 * * *", // 每天凌晨2点 (秒 分 时 日 月 周)
		EnableCleanupTask:     true,          // 默认启用
		RunOnStartup:          true,          // 默认启动时立即执行
		CleanupTimeout:        300,           // 默认5分钟超时
	}

	// 从配置文件读取配置
	if g.Config().Available(ctx) {
		// 读取cron表达式配置
		if cronExpr := g.Config().MustGet(ctx, "cron.cleanup.expression").String(); cronExpr != "" {
			config.CleanupCronExpression = cronExpr
		}

		// 读取是否启用配置
		if enable := g.Config().MustGet(ctx, "cron.cleanup.enable").Bool(); !enable {
			config.EnableCleanupTask = false
		}

		// 读取是否在启动时立即执行配置
		if runOnStartup := g.Config().MustGet(ctx, "cron.cleanup.runOnStartup").Bool(); !runOnStartup {
			config.RunOnStartup = false
		}

		// 读取超时时间配置
		if timeout := g.Config().MustGet(ctx, "cron.cleanup.timeout").Int(); timeout > 0 {
			config.CleanupTimeout = timeout
		}
	}

	return config
}
