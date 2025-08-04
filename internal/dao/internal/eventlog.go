// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventlogDao is the data access object for the table eventlog.
type EventlogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  EventlogColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// EventlogColumns defines and stores column names for the table eventlog.
type EventlogColumns struct {
	EventTime string //
	Adid      string //
	AdId      string //
}

// eventlogColumns holds the columns for the table eventlog.
var eventlogColumns = EventlogColumns{
	EventTime: "event_time",
	Adid:      "adid",
	AdId:      "ad_id",
}

// NewEventlogDao creates and returns a new DAO object for table data access.
func NewEventlogDao(handlers ...gdb.ModelHandler) *EventlogDao {
	return &EventlogDao{
		group:    "default",
		table:    "eventlog",
		columns:  eventlogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EventlogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EventlogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EventlogDao) Columns() EventlogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EventlogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EventlogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EventlogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
