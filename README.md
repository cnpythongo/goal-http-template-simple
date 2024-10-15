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

## 生成Swagger接口文档

* 生成Admin接口文档

```shell
$ swag init -g ./cmd/admin/main.go -o docs/admin/ --exclude ./api
```
或
```shell
$ make doc-admin
```

* 生成前台API接口文档

```shell
$ swag init -g ./cmd/api/main.go -o docs/api/ --exclude ./admin
```
或
```shell
$ make doc-api
```


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

## 鸣谢：

* 代码生成器功能接口 和 菜单管理功能接口 基于 [likeadmin_go](https://gitee.com/likeadmin/likeadmin_go) 的代码生成器进行了适配改造；
