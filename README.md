# goal-http-template-simple

HTTP API服务开发模板, 不使用依赖注入(facebookgo/inject)

## 本地开发

#### 准备：

1. 将 settings/server.json.template文件 复制为 settings/server.json 文件;
2. 按实际需要修改 settings/server.json 文件内的其他各项配置;

#### 运行本地服务:

* 启动前台API服务:

```shell script
$ go run cmd/api/main.go start -c settings/local.json
```

* 启动后台管理API服务:

```shell script
$ go run cmd/admin/main.go start -c settings/local.json
```

* 停止服务：

    ctrl + c 可停止API服务


* 测试服务是否正常启动: 

    {本地IP}/api/ping

## Docker相关：

#### 构建 docker image

* API image

```shell script
$ make build-api
```

* Admin Image
```shell script
$ make build-admin
```


#### 运行Docker容器

* 运行前台API服务
```shell script
$ make run-api
```

* 运行后台API服务

```shell script
$ make run-admin
```

#### 通过 docker-compose 启动服务

将 settings/server.json.template文件 复制为 settings/compose.json 文件.

按实际需要修改 settings/compose.json 文件内的其他各项配置.

* 运行api和admin接口服务
```shell script
$ docker-compose up -d 
```

* 运行All-In-One服务(含mysql,redis,api,admin服务)
```shell script
$ docker-compose -f docker-compose-all-in-one.yml up -d 
```