// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package node

import (
	"context"

	"go-liteflow-editor-client/api/node/v1"
)

type INodeV1 interface {
	GetNode(ctx context.Context, req *v1.GetNodeReq) (res *v1.GetNodeRes, err error)
	AddNode(ctx context.Context, req *v1.AddNodeReq) (res *v1.AddNodeRes, err error)
	DeleteNode(ctx context.Context, req *v1.DeleteNodeReq) (res *v1.DeleteNodeRes, err error)
	UpdateNode(ctx context.Context, req *v1.UpdateNodeReq) (res *v1.UpdateNodeRes, err error)
}
