package liteflow

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "go-liteflow-editor-client/api/liteflow/v1"
)

func (c *ControllerV1) AddLiteflowChain(ctx context.Context, req *v1.AddLiteflowChainReq) (res *v1.AddLiteflowChainRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
