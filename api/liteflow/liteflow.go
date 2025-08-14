// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package liteflow

import (
	"context"

	"go-liteflow-editor-client/api/liteflow/v1"
)

type ILiteflowV1 interface {
	GetLiteflowChain(ctx context.Context, req *v1.GetLiteflowChainReq) (res *v1.GetLiteflowChainRes, err error)
	AddLiteflowChain(ctx context.Context, req *v1.AddLiteflowChainReq) (res *v1.AddLiteflowChainRes, err error)
}
