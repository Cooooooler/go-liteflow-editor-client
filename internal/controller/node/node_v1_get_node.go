package node

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-liteflow-editor-client/api/node/v1"
)

func (c *ControllerV1) GetNode(ctx context.Context, req *v1.GetNodeReq) (res *v1.GetNodeRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
