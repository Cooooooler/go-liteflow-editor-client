// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiteflowChainDao is the data access object for the table liteflow_chain.
type LiteflowChainDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  LiteflowChainColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// LiteflowChainColumns defines and stores column names for the table liteflow_chain.
type LiteflowChainColumns struct {
	Id         string // 主键ID(UUID)
	ChainId    string // 链路ID
	ChainName  string // 链路名称
	ChainDesc  string // 链路描述
	ChainDsl   string // 链路dsl
	ElData     string // EL表达式数据
	Enable     string // 是否启用 1:启用 0:禁用
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// liteflowChainColumns holds the columns for the table liteflow_chain.
var liteflowChainColumns = LiteflowChainColumns{
	Id:         "id",
	ChainId:    "chain_id",
	ChainName:  "chain_name",
	ChainDesc:  "chain_desc",
	ChainDsl:   "chain_dsl",
	ElData:     "el_data",
	Enable:     "enable",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewLiteflowChainDao creates and returns a new DAO object for table data access.
func NewLiteflowChainDao(handlers ...gdb.ModelHandler) *LiteflowChainDao {
	return &LiteflowChainDao{
		group:    "default",
		table:    "liteflow_chain",
		columns:  liteflowChainColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LiteflowChainDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LiteflowChainDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LiteflowChainDao) Columns() LiteflowChainColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LiteflowChainDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LiteflowChainDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LiteflowChainDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
