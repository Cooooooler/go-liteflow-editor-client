package main

import (
	_ "go-liteflow-editor-client/internal/packed"

	// _ "go-liteflow-editor-client/internal/logic"

	// 导入MySQL驱动
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"go-liteflow-editor-client/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
