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
        "/attachments": {
            "post": {
                "description": "用户上传附件",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "通用"
                ],
                "summary": "新增附件",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件流",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "http状态码是200，并且code是200, 表示正常返回；code不是200时表示有业务错误",
                        "schema": {
                            "$ref": "#/definitions/api_attachment.AttachmentResp"
                        }
                    }
                }
            }
        },
        "/auth/captcha": {
            "get": {
                "description": "获取验证码ID和图片base64",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "通用"
                ],
                "summary": "获取验证码ID和图片base64",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 32,
                        "example": 32,
                        "description": "高",
                        "name": "h",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "时间戳字符串，避免缓存",
                        "name": "ts",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 128,
                        "example": 128,
                        "description": "宽",
                        "name": "w",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/api_auth.RespAuthCaptcha"
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
        "/auth/logout": {
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
                    "登录认证"
                ],
                "summary": "退出",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "前台用户登录接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录认证"
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
        "/auth/signup": {
            "post": {
                "description": "前台用户注册接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "登录认证"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_auth.ReqAuthSignup"
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
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/create": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "当前登录用户上传图片任务",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImageFlix-任务"
                ],
                "summary": "当前登录用户上传图片任务",
                "parameters": [
                    {
                        "type": "file",
                        "description": "文件流",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/reduce": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "获取当前登录用户可用资源包余额",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImageFlix-资源包"
                ],
                "summary": "获取当前登录用户可用资源包余额",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_imageflix.CreditReduceReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/start": {
            "post": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "当前登录用户出发开始任务，成功后扣除点数",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImageFlix-任务"
                ],
                "summary": "当前登录用户出发开始任务",
                "parameters": [
                    {
                        "description": "请求体",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/api_imageflix.JobStartReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code不为0时表示有错误",
                        "schema": {
                            "$ref": "#/definitions/goal-app_pkg_render.RespJsonData"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/usable": {
            "get": {
                "security": [
                    {
                        "APIAuth": []
                    }
                ],
                "description": "获取当前登录用户可用资源包余额",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ImageFlix-资源包"
                ],
                "summary": "获取当前登录用户可用资源包余额",
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
                                            "$ref": "#/definitions/api_imageflix.UserCreditResp"
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
        "api_attachment.AttachmentResp": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "上传时间",
                    "type": "string"
                },
                "name": {
                    "description": "文件名",
                    "type": "string"
                },
                "size": {
                    "description": "文件大小, 单位: 字节",
                    "type": "integer"
                },
                "uuid": {
                    "description": "文件UUID",
                    "type": "string"
                }
            }
        },
        "api_auth.ReqAuthSignup": {
            "type": "object",
            "required": [
                "captcha_answer",
                "captcha_id",
                "confirm_password",
                "password"
            ],
            "properties": {
                "captcha_answer": {
                    "description": "验证码,4位",
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "captcha_id": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "confirm_password": {
                    "description": "确认密码",
                    "type": "string",
                    "example": "123456"
                },
                "email": {
                    "description": "Phone    string ` + "`" + `json:\"phone\" binding:\"required\" example:\"13800138000\"` + "`" + ` // 手机号",
                    "type": "string",
                    "example": "foo@bar.com"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "api_auth.ReqUserAuth": {
            "type": "object",
            "required": [
                "captcha_answer",
                "captcha_id",
                "email",
                "password"
            ],
            "properties": {
                "captcha_answer": {
                    "description": "验证码,4位",
                    "type": "string",
                    "maxLength": 6,
                    "minLength": 6
                },
                "captcha_id": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "email": {
                    "description": "Phone    string ` + "`" + `json:\"phone\" binding:\"required\" example:\"13800138000\"` + "`" + ` // 手机号",
                    "type": "string",
                    "example": "foo@bar.com"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "api_auth.RespAuthCaptcha": {
            "type": "object",
            "properties": {
                "captcha_id": {
                    "description": "验证码ID",
                    "type": "string"
                },
                "captcha_img": {
                    "description": "base64编码的验证码图片",
                    "type": "string"
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
                "email": {
                    "description": "邮箱",
                    "type": "string",
                    "example": "foo@bar.com"
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
        "api_imageflix.CreditReduceReq": {
            "type": "object",
            "required": [
                "point"
            ],
            "properties": {
                "point": {
                    "description": "使用点数",
                    "type": "integer"
                }
            }
        },
        "api_imageflix.JobStartReq": {
            "type": "object",
            "required": [
                "job_id"
            ],
            "properties": {
                "job_id": {
                    "description": "任务ID",
                    "type": "integer"
                }
            }
        },
        "api_imageflix.UserCreditResp": {
            "type": "object",
            "properties": {
                "usable": {
                    "description": "可用点数",
                    "type": "integer"
                }
            }
        },
        "api_user.ReqUpdateUser": {
            "type": "object",
            "properties": {
                "avatar": {
                    "description": "用户头像URL",
                    "type": "string",
                    "example": "a/b/c.jpg"
                },
                "gender": {
                    "description": "性别:3-保密,1-男,2-女",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserGender"
                        }
                    ],
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
        },
        "model.UserGender": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "UserGenderMale",
                "UserGenderFemale",
                "UserGenderUnknown"
            ]
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
