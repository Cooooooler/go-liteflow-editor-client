// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiteflowScriptDao is the data access object for the table liteflow_script.
type LiteflowScriptDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  LiteflowScriptColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// LiteflowScriptColumns defines and stores column names for the table liteflow_script.
type LiteflowScriptColumns struct {
	Id         string // 主键ID(UUID)
	ScriptId   string // 脚本ID
	ScriptName string // 脚本名称
	ScriptType string // 脚本类型(groovy,js,python等)
	ScriptData string // 脚本内容
	ScriptDesc string // 脚本描述
	Enable     string // 是否启用 1:启用 0:禁用
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
}

// liteflowScriptColumns holds the columns for the table liteflow_script.
var liteflowScriptColumns = LiteflowScriptColumns{
	Id:         "id",
	ScriptId:   "script_id",
	ScriptName: "script_name",
	ScriptType: "script_type",
	ScriptData: "script_data",
	ScriptDesc: "script_desc",
	Enable:     "enable",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewLiteflowScriptDao creates and returns a new DAO object for table data access.
func NewLiteflowScriptDao(handlers ...gdb.ModelHandler) *LiteflowScriptDao {
	return &LiteflowScriptDao{
		group:    "default",
		table:    "liteflow_script",
		columns:  liteflowScriptColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LiteflowScriptDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LiteflowScriptDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LiteflowScriptDao) Columns() LiteflowScriptColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LiteflowScriptDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LiteflowScriptDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LiteflowScriptDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
