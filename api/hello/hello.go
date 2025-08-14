// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"

	"go-liteflow-editor-client/api/hello/v1"
)

type IHelloV1 interface {
	(ctx context.Context, req *v1.Req) (res *v1.Res, err error)
}
