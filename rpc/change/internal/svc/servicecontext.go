package svc

import (
	"bookstore/rpc/change/internal/config"
	"bookstore/rpc/check/checker"
	"bookstore/rpc/model/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	c       config.Config
	Model   *model.BookModel
	Checker checker.Checker // 3. ServiceContext 添加Check服務
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		c:       c,
		Model:   model.NewBookModel(sqlx.NewMysql(c.DataSource), c.Cache, c.Table),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
	}
}
