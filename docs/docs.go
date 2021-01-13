// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/Base/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户注册账号",
                "parameters": [
                    {
                        "description": "用户名, 昵称, 密码, 角色ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SysUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"注册成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/base/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Base"
                ],
                "summary": "用户登录",
                "operationId": "/base/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"登陆成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/casbin/getPolicyPathByAuthorityId": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Casbin"
                ],
                "summary": "获取权限列表",
                "parameters": [
                    {
                        "description": "权限id, 权限模型列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CasbinInReceive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/casbin/updateCasbin": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Casbin"
                ],
                "summary": "更新角色api权限",
                "parameters": [
                    {
                        "description": "权限id, 权限模型列表",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CasbinInReceive"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"更新成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/demo/bind": {
            "post": {
                "description": "测试数据绑定",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "测试数据绑定",
                "operationId": "/demo/bind",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.DemoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.DemoInput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user/ListPage": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "SysUser"
                ],
                "summary": "分页获取用户列表",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.PageInfo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"success\":true,\"data\":{},\"msg\":\"获取成功\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.DemoInput": {
            "type": "object",
            "required": [
                "age",
                "name",
                "passwd"
            ],
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 20
                },
                "name": {
                    "type": "string",
                    "example": "姓名"
                },
                "passwd": {
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "dto.LoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                },
                "stack": {
                    "type": "object"
                },
                "trace_id": {
                    "type": "object"
                }
            }
        },
        "model.SysAuthority": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "authorityName": {
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "createdAt": {
                    "type": "string"
                },
                "dataAuthorityId": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "deletedAt": {
                    "type": "string"
                },
                "menus": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenus"
                    }
                },
                "parentId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.SysBaseMenuParameter": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "sysBaseMenuID": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                },
                "value": {
                    "type": "string"
                }
            }
        },
        "model.SysBaseMenus": {
            "type": "object",
            "properties": {
                "authoritys": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysAuthority"
                    }
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenus"
                    }
                },
                "component": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "defaultMenu": {
                    "type": "boolean"
                },
                "hidden": {
                    "type": "boolean"
                },
                "icon": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "keepAlive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "parameters": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.SysBaseMenuParameter"
                    }
                },
                "parentId": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                },
                "sort": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "model.SysUser": {
            "type": "object",
            "properties": {
                "authority": {
                    "type": "object",
                    "$ref": "#/definitions/model.SysAuthority"
                },
                "authorityId": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "headerImg": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "request.CasbinInReceive": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "casbinInfos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.CasbinInfo"
                    }
                }
            }
        },
        "request.CasbinInfo": {
            "type": "object",
            "properties": {
                "method": {
                    "type": "string"
                },
                "path": {
                    "type": "string"
                }
            }
        },
        "request.PageInfo": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "pageSize": {
                    "type": "integer"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
