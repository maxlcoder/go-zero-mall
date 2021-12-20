# go-zero-mall

### 目录结构

```
- common
- logs
- service
  - user
    - api
    - model
    - rpc
```

目录结构遵循官方推荐的多个服务结构，使用 `goctl new api xxx` 或者 `goctl new rpc xxx` 命令不能很好的生成需要目录结构，
我这边使用的是在指定的目录下先编写 api 定义文件和 rpc 定义文件，再进入目录使用

```shell
goctl api go -api user.api -dir . -style gozero

goctl rpc proto -src user.proto -dir .
```
生成指定的目录结构

### 模版管理

由于默认的 api 定义中如果需要定义如下通用结构的返回，你需要对 api 做嵌套定义。

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    ...
  }
}
```
而实际上这些嵌套都是重复的，所以按照官方的文档，我这边使用自定义模版。在 `response/response.go` 
中定义了通用的返回。有关自定义模版的使用，建议参考官方说明，说明一下使用时，如果指定了自定义模版初始化目录，
在使用相关命令时需要带上 `--home xxx` 的参数

### 错误处理

一开始，我是想通过自定义模版来修改成功的返回，和错误的返回都返回对应的 json 结构，
或者在中间件中做各种参数的校验处理，结果发现，这个都会于默认的 api 定义的 json 序列化有冲突。
所以需要使用官方说明中的错误处理，使用 `httpx.SetErrorHandler` 在启动的 `main` 函数中设置
错误响应，这样就能让各处的错误响应保持一致。具体的可以参考代码

### ORM 替换

替换了 `sqlx` 为 `gorm`，具体的参考 `svc/servicecontext.go` 启动配置时进行替换，替换之后目前是失去了系统
sqlx 的缓存，但是这个可以通过后续 gorm 自身设置来来启用缓存

### jwt 和 中间件

按照文档调试了 jwt 和中间件，可以参考代码

### 日志输出

默认的日志输出到控制台，这边在配置文件中加入相应的设置配置到目录。 `etc/xx.yaml` 和 `config/config.go`
中的配置是一一对应的，这个就有点像 `gin` 开发时使用的各种加载配置的扩展一样，go-zero 已经帮我们做了配置的处理。
**这里要重点说明下，当前项目的配置都是本地测试，所以没有做配置的过滤，正确的是要在部署时做配置设置的，而不是应写在代码中**

### docker 部署 api 和 rpc

根据官方 [介绍](https://go-zero.dev/cn/goctl-other.html) 这里生成 api 和 rpc 两个 Dockerfile 文件

```shell
# 生成 Dockerfile
$ cd service/user/api && goctl docker -go user.go
$ cd service/user/rpc && goctl docker -go user.go

# 运行 mysql 容器
$ docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7

# 运行 etcd 容器
$ docker run -itd --name Etcd --env ALLOW_NONE_AUTHENTICATION=yes bitnami/etcd

# 新建镜像
$ cd ../go-zero-mall
$ docker build -t go-zero-mall-user-api:v1 -f service/user/api/Dockerfile .
$ docker build -t go-zero-mall-user-rpc:v1 -f service/user/rpc/Dockerfile .


# 开不同窗口启动 api 和 rpc 镜像
$ docker run --rm -it -p 8080:8080 go-zero-mall-user-rpc:v1
$ docker run --rm -it -p 8888:8888 go-zero-mall-user-api:v1
```

这期间主要是控制好配置
1. 启动 etcd 容器获取 etcd 的 ip 
2. 启动 mysql 容器获取 mysql 的 ip，
3. 替换掉 api 和 rpc 中的关于 etcd 和 mysql 的 ip （之前的是本地的配置）
4. 由于默认的 rpc 是监听的 127.0.0.1 的 host，镜像之后，是用的容器的 ip 所以这里 rpc 的监听是需要缓换成 0.0.0.0 的

### 待办

目前 api 和 rpc 的调用过程已经处理完成了，如果要对 go-zero 有较深的理解，
建议最好把 grpc 的原理弄清楚，这样理解起来会更快。

后续将对部署细节做仔细研究，尝试各种分离式部署





