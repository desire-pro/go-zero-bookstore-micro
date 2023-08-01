package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource string
	Table      string
	Cache      cache.CacheConf
	Check      zrpc.RpcClientConf // 2.添加 check 服務配置
}
