package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"go-liteflow-editor-client/internal/controller/liteflow"
	"go-liteflow-editor-client/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 启动定时任务服务
			cronService := service.NewCronService()
			cronService.StartCleanupTask()

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					liteflow.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
