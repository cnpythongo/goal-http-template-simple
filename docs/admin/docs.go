// Package admin Code generated by swaggo/swag. DO NOT EDIT
package admin

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/account/history/list": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "获取用户登录历史记录列表",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "登录历史记录列表",
                "parameters": [
                    {
                        "type": "string",
                        "example": "2023-09-01 01:30:59",
                        "description": "数据创建时间起始",
                        "name": "created_at_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "abc@abc.com",
                        "description": "用户邮箱",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2023-09-01 22:59:59",
                        "description": "数据创建时间截止",
                        "name": "last_login_at_end",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "example": 10,
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Tom",
                        "description": "用户昵称",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "13800138000",
                        "description": "用户手机号",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "example": 123,
                        "description": "用户ID",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "abcef123",
                        "description": "用户UUID",
                        "name": "user_uuid",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.ReqGetHistoryList"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/account/login": {
            "post": {
                "description": "后台管理系统登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录退出"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.ReqAdminAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/goal-app_admin_types.RespAdminAuth"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/account/logout": {
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "退出后台管理系统\n前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页\n后端可以执行清理redis缓存, 设置token黑名单等操作",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录退出"
                ],
                "summary": "退出",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    }
                }
            }
        },
        "/account/user/create": {
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "创建新用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.ReqCreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示错误",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/goal-app_admin_types.RespUserDetail"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/account/user/delete": {
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "删除单个用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.RespEmptyJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    }
                }
            }
        },
        "/account/user/detail": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "获取用户详情",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "通过用户UUID获取用户详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户UUID",
                        "name": "uuid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.RespUserDetail"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    }
                }
            }
        },
        "/account/user/list": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "获取用户列表",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "获取用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "example": "2023-09-01 01:30:59",
                        "description": "数据创建时间起始",
                        "name": "created_at_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "abc@abc.com",
                        "description": "邮箱,模糊查旬",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "description": "是否admin, true or false",
                        "name": "is_admin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2023-09-01 22:59:59",
                        "description": "最近登录时间截止",
                        "name": "last_login_at_end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2023-09-01 01:30:59",
                        "description": "最近登录时间起始",
                        "name": "last_login_at_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "example": 10,
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Tom",
                        "description": "昵称,模糊查询",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "13800138000",
                        "description": "手机号,模糊查询",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "INACTIVE",
                                "ACTIVE",
                                "FREEZE",
                                "DELETE"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "example": [
                            "FREEZE",
                            "ACTIVE"
                        ],
                        "description": "用户状态,多种状态过滤使用逗号分隔",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/goal-app_admin_types.RespGetUserList"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "result": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/goal-app_admin_types.RespUserDetail"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/account/user/update": {
            "post": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "更新用户数据",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "更新用户",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.ReqUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_admin_types.RespEmptyJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    }
                }
            }
        },
        "/system/user/list": {
            "get": {
                "security": [
                    {
                        "AdminAuth": []
                    }
                ],
                "description": "获取系统用户列表，与 /account/user/list API相同，调用时总是传 is_admin=true 即可",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "系统管理"
                ],
                "summary": "获取系统用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "example": "2023-09-01 01:30:59",
                        "description": "数据创建时间起始",
                        "name": "created_at_start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "abc@abc.com",
                        "description": "邮箱,模糊查旬",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "example": true,
                        "description": "是否admin, true or false",
                        "name": "is_admin",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2023-09-01 22:59:59",
                        "description": "最近登录时间截止",
                        "name": "last_login_at_end",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "2023-09-01 01:30:59",
                        "description": "最近登录时间起始",
                        "name": "last_login_at_start",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "example": 10,
                        "description": "每页数量",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "Tom",
                        "description": "昵称,模糊查询",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "example": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "13800138000",
                        "description": "手机号,模糊查询",
                        "name": "phone",
                        "in": "query"
                    },
                    {
                        "type": "array",
                        "items": {
                            "enum": [
                                "INACTIVE",
                                "ACTIVE",
                                "FREEZE",
                                "DELETE"
                            ],
                            "type": "string"
                        },
                        "collectionFormat": "multi",
                        "example": [
                            "FREEZE",
                            "ACTIVE"
                        ],
                        "description": "用户状态,多种状态过滤使用逗号分隔",
                        "name": "status",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/goal-app_admin_types.RespGetUserList"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "result": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/goal-app_admin_types.RespUserDetail"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "goal-app_admin_types.ReqAdminAuth": {
            "type": "object",
            "required": [
                "password",
                "phone"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string",
                    "example": "13800138000"
                }
            }
        },
        "goal-app_admin_types.ReqCreateUser": {
            "type": "object",
            "required": [
                "password",
                "password_confirm",
                "phone"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "abc@a.com"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string",
                    "example": "Tom"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "password_confirm": {
                    "description": "确认密码",
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string",
                    "example": "13800138000"
                }
            }
        },
        "goal-app_admin_types.ReqGetHistoryList": {
            "type": "object",
            "properties": {
                "created_at_start": {
                    "description": "数据创建时间起始",
                    "type": "string",
                    "example": "2023-09-01 01:30:59"
                },
                "email": {
                    "description": "用户邮箱",
                    "type": "string",
                    "example": "abc@abc.com"
                },
                "last_login_at_end": {
                    "description": "数据创建时间截止",
                    "type": "string",
                    "example": "2023-09-01 22:59:59"
                },
                "limit": {
                    "description": "每页数量",
                    "type": "integer",
                    "default": 10,
                    "example": 10
                },
                "nickname": {
                    "description": "用户昵称",
                    "type": "string",
                    "example": "Tom"
                },
                "page": {
                    "description": "页码",
                    "type": "integer",
                    "default": 1,
                    "example": 1
                },
                "phone": {
                    "description": "用户手机号",
                    "type": "string",
                    "example": "13800138000"
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "integer",
                    "example": 123
                },
                "user_uuid": {
                    "description": "用户UUID",
                    "type": "string",
                    "example": "abcef123"
                }
            }
        },
        "goal-app_admin_types.ReqUpdateUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "用户头像URL",
                    "type": "string",
                    "example": "a/b/c.jpg"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "abc@abc.com"
                },
                "gender": {
                    "description": "性别:3-保密,1-男,2-女",
                    "type": "integer",
                    "example": 3
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string",
                    "example": "Tom"
                },
                "signature": {
                    "description": "个性化签名",
                    "type": "string",
                    "example": "haha"
                },
                "status": {
                    "description": "用户状态",
                    "type": "string",
                    "example": "FREEZE"
                }
            }
        },
        "goal-app_admin_types.RespAdminAuth": {
            "type": "object",
            "properties": {
                "expire_time": {
                    "description": "过期时间",
                    "type": "string"
                },
                "token": {
                    "description": "令牌",
                    "type": "string"
                },
                "user": {
                    "description": "用户基本信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/goal-app_admin_types.RespAdminAuthUser"
                        }
                    ]
                }
            }
        },
        "goal-app_admin_types.RespAdminAuthUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "last_login_at": {
                    "description": "最近的登录时间",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "phone": {
                    "description": "带掩码的手机号",
                    "type": "string",
                    "example": "138****8000"
                },
                "uuid": {
                    "description": "用户uuid",
                    "type": "string"
                }
            }
        },
        "goal-app_admin_types.RespEmptyJson": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "结果码：0-成功，其它-失败",
                    "type": "integer",
                    "example": 0
                },
                "msg": {
                    "description": "消息, code不为0时，返回简单的错误描述",
                    "type": "string",
                    "example": "ok"
                }
            }
        },
        "goal-app_admin_types.RespGetUserList": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "result": {},
                "total": {
                    "type": "integer"
                }
            }
        },
        "goal-app_admin_types.RespUserDetail": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "账号创建时间",
                    "type": "string",
                    "example": "2023-09-01 13:30:59"
                },
                "is_admin": {
                    "description": "是否管理员",
                    "type": "boolean",
                    "example": false
                },
                "last_login_at": {
                    "description": "最近登录时间",
                    "type": "string",
                    "example": "2023-09-01 13:30:59"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string",
                    "example": "Tom"
                },
                "phone": {
                    "description": "手机号",
                    "type": "string",
                    "example": "13800138000"
                },
                "status": {
                    "description": "用户状态",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserStatusType"
                        }
                    ],
                    "example": "ACTIVE"
                },
                "uuid": {
                    "description": "用户UUID,32位字符串",
                    "type": "string",
                    "example": "826d6b1aa64d471d822d667e92218158"
                }
            }
        },
        "goal-app_pkg_render.RespJsonData": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.UserStatusType": {
            "type": "string",
            "enum": [
                "INACTIVE",
                "ACTIVE",
                "FREEZE",
                "DELETE"
            ],
            "x-enum-varnames": [
                "UserStatusInactive",
                "UserStatusActive",
                "UserStatusFreeze",
                "UserStatusDelete"
            ]
        }
    },
    "securityDefinitions": {
        "AdminAuth": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "后台管理系统接口文档",
	Description:      "http状态码是200，code为0时表示正常返回；code不为0时表示有业务错误。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
