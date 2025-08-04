// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package testinsert

import (
	"context"

	v1 "bugtest/api/testinsert/v1"
)

type ITestinsertV1 interface {
	TestInsert(ctx context.Context, req *v1.TestInsertReq) (res *v1.TestInsertRes, err error)
}
