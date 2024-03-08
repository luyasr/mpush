// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

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
        "/api/v1/channel": {
            "post": {
                "description": "创建频道",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channel"
                ],
                "summary": "创建频道",
                "parameters": [
                    {
                        "description": "创建频道请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/channel.CreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/channel.Channel"
                        }
                    }
                }
            }
        },
        "/api/v1/channel/delete/{id}": {
            "delete": {
                "description": "删除频道",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channel"
                ],
                "summary": "删除频道",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "频道ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/channel/{id}": {
            "put": {
                "description": "更新频道",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channel"
                ],
                "summary": "更新频道",
                "parameters": [
                    {
                        "description": "更新频道请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/channel.UpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "查询频道",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "channel"
                ],
                "summary": "查询频道",
                "parameters": [
                    {
                        "description": "查询频道请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/channel.QueryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/channel.Channels"
                        }
                    }
                }
            }
        },
        "/api/v1/message": {
            "post": {
                "description": "客户端发送消息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "message"
                ],
                "summary": "客户端发送消息",
                "parameters": [
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/message.ProducerReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "成功",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/api/v1/token/login": {
            "post": {
                "description": "登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/token.Tk"
                        }
                    }
                }
            }
        },
        "/api/v1/token/logout": {
            "post": {
                "description": "登出",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token"
                ],
                "summary": "登出",
                "parameters": [
                    {
                        "description": "登出请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.Tk"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/token/refresh": {
            "post": {
                "description": "刷新",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "token"
                ],
                "summary": "刷新",
                "parameters": [
                    {
                        "description": "刷新请求参数",
                        "name": "Object",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/token.Tk"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "post": {
                "description": "创建用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "创建用户请求参数",
                        "name": "Object",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/user.CreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "description": "查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "channel.Channel": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt 创建时间",
                    "type": "integer"
                },
                "id": {
                    "description": "Id 频道ID",
                    "type": "integer"
                },
                "name": {
                    "description": "Name 频道名称",
                    "type": "string"
                },
                "secret": {
                    "description": "Secret 频道密钥",
                    "type": "string"
                },
                "token": {
                    "description": "Token 频道token",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt 更新时间",
                    "type": "integer"
                },
                "url": {
                    "description": "Url 频道地址",
                    "type": "string"
                },
                "user_id": {
                    "description": "UserId 用户ID",
                    "type": "integer"
                }
            }
        },
        "channel.Channels": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/channel.Channel"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "channel.CreateReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "channel.QueryReq": {
            "type": "object",
            "properties": {
                "keywords": {
                    "type": "string"
                },
                "page_number": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                }
            }
        },
        "channel.UpdateReq": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "message.ProducerReq": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "content": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "token.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "token.Tk": {
            "type": "object",
            "required": [
                "access_token",
                "refresh_token"
            ],
            "properties": {
                "access_token": {
                    "description": "登录token",
                    "type": "string"
                },
                "refresh_token": {
                    "description": "刷新token",
                    "type": "string"
                }
            }
        },
        "user.CreateReq": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 6
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "user.Role": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "RoleAnonymous",
                "RoleAdmin"
            ]
        },
        "user.Status": {
            "type": "integer",
            "enum": [
                0,
                1
            ],
            "x-enum-varnames": [
                "StatusNormal",
                "StatusDeleted"
            ]
        },
        "user.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "description": "CreatedAt 创建时间",
                    "type": "integer"
                },
                "deleted_at": {
                    "description": "DeletedAt 删除时间",
                    "type": "integer"
                },
                "email": {
                    "description": "Email 邮箱",
                    "type": "string"
                },
                "id": {
                    "description": "ID 用户ID",
                    "type": "integer"
                },
                "nickname": {
                    "description": "Nickname 昵称",
                    "type": "string"
                },
                "password": {
                    "description": "Password 密码",
                    "type": "string"
                },
                "phone": {
                    "description": "Phone 手机号",
                    "type": "string"
                },
                "role": {
                    "description": "Role 角色",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Role"
                        }
                    ]
                },
                "status": {
                    "description": "Status 状态",
                    "allOf": [
                        {
                            "$ref": "#/definitions/user.Status"
                        }
                    ]
                },
                "updated_at": {
                    "description": "UpdatedAt 更新时间",
                    "type": "integer"
                },
                "username": {
                    "description": "Username 用户名",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
