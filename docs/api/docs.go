// Package api Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/login": {
            "post": {
                "description": "前台用户登录接口",
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
                            "$ref": "#/definitions/api_auth.ReqUserAuth"
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
                                            "$ref": "#/definitions/api_auth.RespUserAuth"
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
        "/logout": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "前台用户退出\n前端调用该接口，无需关注结果，自行清理掉请求头的 Authorization，页面跳转至首页\n后端可以执行清理redis缓存, 设置token黑名单等操作",
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
        "/users/me": {
            "get": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "获取当前登录用户的信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取当前登录用户的信息",
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
                                            "$ref": "#/definitions/api_user.RespUserInfo"
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
        "/users/me/password/update": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "修改当前登录用户的登录密码",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "修改用户密码",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_user.ReqUpdateUserPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
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
        "/users/me/profile": {
            "get": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "当前登录用户的个人资料",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户个人资料",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
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
        "/users/me/profile/update": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "更新当前登录用户的个人资料",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户个人资料",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_user.ReqUpdateUserProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
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
        "/users/me/update": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "更新当前登录用户的基本信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户基本信息",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_user.ReqUpdateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
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
        "/users/{uuid}": {
            "get": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "通过用户UUID获取用户的信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户的信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户UUID",
                        "name": "uuid",
                        "in": "path",
                        "required": true
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
                                            "$ref": "#/definitions/api_user.RespUserInfo"
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
        "api_auth.ReqUserAuth": {
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
        "api_auth.RespUserAuth": {
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
                            "$ref": "#/definitions/api_auth.RespUserInfo"
                        }
                    ]
                }
            }
        },
        "api_auth.RespUserInfo": {
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
        "api_user.ReqUpdateUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "头像",
                    "type": "string"
                },
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "nickname": {
                    "description": "昵称",
                    "type": "string"
                },
                "uuid": {
                    "description": "用户UUID",
                    "type": "string"
                }
            }
        },
        "api_user.ReqUpdateUserPassword": {
            "type": "object",
            "properties": {
                "new_password": {
                    "description": "新密码",
                    "type": "string"
                },
                "old_password": {
                    "description": "旧密码",
                    "type": "string"
                },
                "uuid": {
                    "description": "用户UUID",
                    "type": "string"
                }
            }
        },
        "api_user.ReqUpdateUserProfile": {
            "type": "object",
            "properties": {
                "id_number": {
                    "description": "身份证号",
                    "type": "string"
                },
                "real_name": {
                    "description": "真实姓名",
                    "type": "string"
                },
                "user_id": {
                    "description": "用户ID",
                    "type": "integer"
                }
            }
        },
        "api_user.RespUserInfo": {
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
        }
    },
    "securityDefinitions": {
        "APIAuth": {
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
	Title:            "前台接口文档",
	Description:      "http状态码是200，code为0时表示正常返回；code不为0时表示有业务错误。",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
