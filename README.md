# goenv

1. 使用 `envutils.Configure(config any)` 生成默认配置到 `default.yml`。 
    + 如果 config 中的字段具有方法 `SetDefaults() , 则会被调用， 并且值会被写到 default.yml 中。
2. 同时读取并解析 `default.yml, config.yml` 到 config 结构体中。
    + 如果 config 中的字段具有方法 `Initialize()`, 则会被调用。 并进行初始化。

## Quick Start

1. 定义一个 golang 

```go
package envutils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Endpoint string `env:"endpoint"`
	client   *redis.Client
}

// 设置默认值并展示
func (r *RedisClient) SetDefaults() {
	if r.Endpoint == "" {
		r.Endpoint = "redis://:Password@localhost:6379/0"
	}
}
func (r *RedisClient) Initialize() {
	r.connect()
}

func (r *RedisClient) connect() {
	if r.client != nil {
		return
	}

	opts, err := redis.ParseURL(r.Endpoint)
	if err != nil {
		panic(err)
	}

	r.client = redis.NewClient(opts)
	_ = r.client.Conn().Ping(context.Background())
}
```

## 2. **生成变量模版**, **初始化变量结构体**

```go
var (
    Redis = &RedisClient{}
)

func init(){
    config:=&struct{
        Redis: *RedisClient
    }{
        Redis: *Redis
    }

    Configure(config)
}
```



## Usage

1. 读取 struct 并将配置文件保存在 config.yml 中

```go

// Test 3
type MysqlServer struct {
	ListenAddr string `env:"listenAddr"`
	Auth       string `env:"auth"`
	DBName     string `env:"dbName"`
}

func (my *MysqlServer) SetDefaults() {
	if my.ListenAddr == "" {
		my.ListenAddr = "localhost:3306"
	}
}

type RedisServer struct {
	DSN string `env:"dsn"`
}

func (r *RedisServer) SetDefaults() {
	if r.DSN == "" {
		r.DSN = "redis://:Password@localhost:6379/0"
	}
}

func Test_ConfP_Server(t *testing.T) {

	config := &struct {
		MysqlServer *MysqlServer
		RedisServer *RedisServer
	}{
		MysqlServer: &MysqlServer{},
		RedisServer: &RedisServer{},
	}

	// 设置默认值
	err := CallSetDefaults(config)
	if err != nil {
		panic(err)
	}

	// 序列化配置
	data, err := Marshal(config, "AppName")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}

```

2. 查看保存文件

```bash
# cat config.yml 
AppName__MysqlServer_auth: ""
AppName__MysqlServer_dbName: ""
AppName__MysqlServer_listenAddr: localhost:3306
AppName__RedisServer_dsn: redis://:Password@localhost:6379/0
```

3. 从文件中读取配置

```go

func read() {

	server := &Server{
		Address: "0.0.0.0",
	}

	config := &struct {
		Server *Server
	}{
		Server: server,
	}

	err := envutils.UnmarshalFile(config, appname, cfgfile)
	if err != nil {
		panic(err)
	}

	fmt.Println("addr=", config.Server.Address)
	fmt.Println("port=", config.Server.Port)
}
// addr= 192.168.100.100
// port= 80
```

## Todo

+ [x] 将结构体的 tag 转换为 config.yml
+ [x] 从环境变量赋值结构体
+ [x] 结构体字段支持 tag 名称
