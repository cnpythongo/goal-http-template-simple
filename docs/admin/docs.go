// Code generated by swaggo/swag. DO NOT EDIT.

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
                            "$ref": "#/definitions/types.ReqAdminAuth"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.RespAdminAuth"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/types.RespFailJson"
                        }
                    }
                }
            }
        },
        "/account/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "退出后台管理系统，前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录退出"
                ],
                "summary": "退出",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/types.RespEmptyJson"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "types.ReqAdminAuth": {
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
        "types.RespAdminAuth": {
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
                            "$ref": "#/definitions/types.RespAdminAuthUser"
                        }
                    ]
                }
            }
        },
        "types.RespAdminAuthUser": {
            "type": "object",
            "properties": {
                "last_login_at": {
                    "description": "最近的登录时间",
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
        "types.RespEmptyJson": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "结果码：0-成功，其它-失败",
                    "type": "integer"
                },
                "msg": {
                    "description": "消息, code不为0时，返回简单的错误描述",
                    "type": "string"
                }
            }
        },
        "types.RespFailJson": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "结果码：0-成功，其它-失败",
                    "type": "integer"
                },
                "error": {
                    "description": "具体的错误信息",
                    "type": "string"
                },
                "msg": {
                    "description": "消息, code不为0时，返回简单的错误描述",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8200",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "后台管理系统接口文档",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
