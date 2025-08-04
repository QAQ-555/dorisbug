package testinsert

import (
	"context"
	"fmt"

	v1 "bugtest/api/testinsert/v1"
	"bugtest/internal/dao"
	"bugtest/internal/model/do"

	"github.com/gogf/gf/v2/os/gtime"
)

func (c *ControllerV1) TestInsert(ctx context.Context, req *v1.TestInsertReq) (res *v1.TestInsertRes, err error) {
	log := &do.Eventlog{
		EventTime: gtime.Now(),
		Adid:      req.Adid,
		AdId:      req.AdID,
	}
	fmt.Println(log)
	for i := 0; i < 20; i++ {
		dao.Eventlog.Ctx(ctx).Data(log).Insert()
	}

	return nil, nil
}
