// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LiteflowLogDao is the data access object for the table liteflow_log.
type LiteflowLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LiteflowLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LiteflowLogColumns defines and stores column names for the table liteflow_log.
type LiteflowLogColumns struct {
	Id            string // 主键ID(UUID)
	RequestId     string // 请求ID
	ChainId       string // 链路ID
	ChainName     string // 链路名称
	NodeId        string // 节点ID
	NodeName      string // 节点名称
	ExecuteStatus string // 执行状态(SUCCESS,FAILED,RUNNING)
	ExecuteTime   string // 执行耗时(毫秒)
	ErrorMsg      string // 错误信息
	InputData     string // 输入数据
	OutputData    string // 输出数据
	CreateTime    string // 创建时间
}

// liteflowLogColumns holds the columns for the table liteflow_log.
var liteflowLogColumns = LiteflowLogColumns{
	Id:            "id",
	RequestId:     "request_id",
	ChainId:       "chain_id",
	ChainName:     "chain_name",
	NodeId:        "node_id",
	NodeName:      "node_name",
	ExecuteStatus: "execute_status",
	ExecuteTime:   "execute_time",
	ErrorMsg:      "error_msg",
	InputData:     "input_data",
	OutputData:    "output_data",
	CreateTime:    "create_time",
}

// NewLiteflowLogDao creates and returns a new DAO object for table data access.
func NewLiteflowLogDao(handlers ...gdb.ModelHandler) *LiteflowLogDao {
	return &LiteflowLogDao{
		group:    "default",
		table:    "liteflow_log",
		columns:  liteflowLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LiteflowLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LiteflowLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LiteflowLogDao) Columns() LiteflowLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LiteflowLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LiteflowLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LiteflowLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
