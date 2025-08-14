// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiteflowNodeDao is the data access object for the table liteflow_node.
type LiteflowNodeDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  LiteflowNodeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// LiteflowNodeColumns defines and stores column names for the table liteflow_node.
type LiteflowNodeColumns struct {
	Id         string // 主键ID(UUID)
	NodeId     string // 节点ID
	NodeName   string // 节点名称
	NodeType   string // 节点类型(common,switch,for,while等)
	ClassName  string // 节点实现类名
	ScriptId   string // 关联脚本ID
	NodeDesc   string // 节点描述
	Enable     string // 是否启用 1:启用 0:禁用
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// liteflowNodeColumns holds the columns for the table liteflow_node.
var liteflowNodeColumns = LiteflowNodeColumns{
	Id:         "id",
	NodeId:     "node_id",
	NodeName:   "node_name",
	NodeType:   "node_type",
	ClassName:  "class_name",
	ScriptId:   "script_id",
	NodeDesc:   "node_desc",
	Enable:     "enable",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewLiteflowNodeDao creates and returns a new DAO object for table data access.
func NewLiteflowNodeDao(handlers ...gdb.ModelHandler) *LiteflowNodeDao {
	return &LiteflowNodeDao{
		group:    "default",
		table:    "liteflow_node",
		columns:  liteflowNodeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LiteflowNodeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LiteflowNodeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LiteflowNodeDao) Columns() LiteflowNodeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LiteflowNodeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LiteflowNodeDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *LiteflowNodeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
