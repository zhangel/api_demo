// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/api/v1/information/sample/delete": {
            "post": {
                "description": "删除样本数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "删除样本列表",
                "parameters": [
                    {
                        "default": 0,
                        "description": "默认为 空",
                        "name": "level",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "integer"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "默认为 空",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ExecSuccess"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/information/sample/get": {
            "get": {
                "description": "获取样本列表数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "获取样本列表数据",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "默认为 1",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 25,
                        "description": "默认为 25",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "默认为 空",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "默认为 空",
                        "name": "token",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "默认为 空",
                        "name": "md5",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "默认为 空",
                        "name": "sha1",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JsonOut"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/model.SampleInfo"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/information/sample/insert": {
            "post": {
                "description": "新增样本数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "新增样本信息",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 70,
                        "description": "默认为 空",
                        "name": "level",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "2a57220fe8f64481b1311c892b788da5",
                        "description": "默认为 空",
                        "name": "md5",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "ff4cd7d8ee07f35037e834cc0f356f5fa159c871",
                        "description": "默认为 空",
                        "name": "sha1",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "default": "admin",
                        "description": "默认为 空",
                        "name": "operator",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "默认为 空",
                        "name": "token",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ExecSuccess"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ServerError"
                        }
                    }
                }
            }
        },
        "/api/v1/information/sample/update": {
            "get": {
                "description": "更新样本数据",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "标签"
                ],
                "summary": "更新样本信息",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "默认为 空",
                        "name": "level",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "默认为 空",
                        "name": "id",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "默认为 空",
                        "name": "token",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ExecSuccess"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/model.ServerError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ExecSuccess": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "model.JsonOut": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 200
                },
                "data": {},
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "model.SampleInfo": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string",
                    "example": "2022-10-20 11:00:01"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "level": {
                    "type": "integer",
                    "example": 70
                },
                "md5": {
                    "type": "string",
                    "example": "2a57220fe8f64481b1311c892b788da5"
                },
                "operator": {
                    "type": "string",
                    "example": "admin"
                },
                "sha1": {
                    "type": "string",
                    "example": "ff4cd7d8ee07f35037e834cc0f356f5fa159c871"
                }
            }
        },
        "model.ServerError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 500
                },
                "data": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string",
                    "example": "服务器内部错误"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
