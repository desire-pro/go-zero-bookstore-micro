package svc

import (
	"bookstore/api/internal/config"
	"bookstore/rpc/add/adder"
	"bookstore/rpc/change/changer"
	"bookstore/rpc/check/checker"

	"github.com/zeromicro/go-zero/zrpc"
)

// 通過 ServiceContext 在不同業務邏輯之間傳遞服務依賴
type ServiceContext struct {
	Config  config.Config
	Adder   adder.Adder
	Checker checker.Checker
	Changer changer.Changer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		Adder:   adder.NewAdder(zrpc.MustNewClient(c.Add)),
		Checker: checker.NewChecker(zrpc.MustNewClient(c.Check)),
		Changer: changer.NewChanger(zrpc.MustNewClient(c.Change)),
	}
}
