package main

import (
	_ "go-liteflow-editor-client/internal/packed"

	// _ "go-liteflow-editor-client/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"go-liteflow-editor-client/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
