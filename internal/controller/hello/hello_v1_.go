package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-liteflow-editor-client/api/hello/v1"
)

func (c *ControllerV1) (ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
