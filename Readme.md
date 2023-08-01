## 介紹：
用go-zero建置的微服務範例

## 代碼結構:

### API接口 代碼結構

```Plain Text
api
├── bookstore.api                  // api描述文件，定義api
├── bookstore.go                   // main入口定義
├── etc                             
│   └── bookstore-api.yaml         // 用來定義項目配置，所有的配置項都可以寫在bookstore-api.yaml中
└── internal
    ├── config
    │   └── config.go              // 服務的配置定義
    ├── handler                    // API 文件中定義的路由對應的handler的實現
    │   ├── addhandler.go          // 實現addHandler
    │   ├── checkhandler.go        // 實現checkHandler
    │   └── routes.go              // 定義路由處理
    ├── logic                      // 放每個路由對應的業務邏輯
    │   ├── addlogic.go            // 實現AddLogic
    │   └── checklogic.go          // 實現CheckLogic
    ├── svc
    │   └── servicecontext.go      // 定義業務邏輯處理的依賴，在main函數創建依賴，通過ServiceContext傳遞給handler和logic
    └── types
        └── types.go               // 定義了api請求和返回數據結構

```

> 區分 handler 和 logic 是為了讓業務處理部分盡可能減少依賴，把 HTTP requests 和邏輯處理代碼隔離開，便於後續拆分成 RPC service

---

### RPC 服務代碼結構

```Plain Text
rpc/add
├── add.go                      // rpc服務main函數
├── add.proto                   // rpc接口定義
├── adder
│   ├── adder.go                // 提供了外部調用方法，無需修改
│   └── types.go                // request/response結構體定義
├── etc
│   └── add.yaml                // 配置文件
├── internal
│   ├── config
│   │   └── config.go           // 配置定義
│   ├── logic
│   │   └── addlogic.go         // add業務邏輯在這實現
│   ├── server
│   │   └── adderserver.go      // 調用入口, 不需要修改
│   └── svc
│       └── servicecontext.go   // 定義ServiceContext，傳遞依賴
└── pb
    └── add.pb.go               // protoc-gen-go產生的結構檔案

```
___________________________________________________________________________________________________________________

## RPC 服務
1、定義服務接口和資料結構
2、資料庫相互隔離、通過 rpc 相互調用
3、服務間相互調用, 例如：add 之前需要調用 check 服務 
    1. 配置 add.yaml 文件添加 check 配置項
    2. Config addRpc 服務下 添加配置項
    3. ServiceContext 添加服務上下文配置
    4. 在 AddLogic 添加邏輯中 調用 check 服務

---

### 新增 RPC 服務 

1、rpc資料夾下建立需要的服務，生成 proto 文件
    執行命令：
    cd rpc; mkdir change
    goctl rpc --o change/change.proto

2、編寫 proto 內容

3、通過.proto 文件一鍵快速生成一個 rpc 服務
    執行命令：
    cd change;
    goctl rpc protoc change.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.

    >> 生成目錄結構：參見上方 RPC 服務代碼結構

4、添加配置
    etc\change.yaml

5、聲明配置類型
    internal\config\config.go

6、填充依賴
    internal\svc\servicecontext.go

7、邏輯
    internal\logic\deletelogic.go
    internal\logic\updatelogic.go

---

## API 接口新增 RPC 服務

1、 添加yaml 配置
    api\etc\bookstore-api.yaml

2、 Config 添加服務依賴
    api\internal\config\config.go

3、 完善服務依賴
    api\internal\svc\servicecontext.go

4、 編寫業務邏輯
    api\internal\logic\deletelogic.go
    api\internal\logic\updatelogic.go

___________________________________________________________________________________________________________________

## 啟動：

1、docker-compose-db 創建 mysql、redis、etcd、phpmyadmin、redisadmin 容器
    執行命令：
    docker-compose -f docker-compose-db.yml up -d

2、docker-compose-prom 創建 prometheus、grafana 容器
    執行命令：
    docker-compose -f docker-compose-prom.yml up -d

注意：
- http://127.0.0.1:9090/ 普羅米修斯 web url
- http://localhost:3000/ Grafana web url
    
3、docker-compose 創建rpc服務容器，rpc 服務基於依賴於 etcd 服務
    執行命令：
    docker-compose up      

___________________________________________________________________________________________________________________