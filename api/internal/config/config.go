package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// Config 添加服務依賴
type Config struct {
	rest.RestConf
	Add    zrpc.RpcClientConf
	Check  zrpc.RpcClientConf
	Change zrpc.RpcClientConf
}
