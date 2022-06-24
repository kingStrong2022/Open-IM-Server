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
        "/user/update_user_info": {
            "post": {
                "description": "修改用户信息 userID faceURL等",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户信息"
                ],
                "summary": "修改用户信息",
                "operationId": "UpdateUserInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "im token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "请求",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/base_info.UpdateSelfUserInfoReq"
                        }
                    }
                ],
                "responses": {
                    "0": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/base_info.UpdateUserInfoResp"
                        }
                    },
                    "400": {
                        "description": "errCode为400 一般为参数输入错误, token未带上等",
                        "schema": {
                            "$ref": "#/definitions/base_info.UpdateUserInfoResp"
                        }
                    },
                    "500": {
                        "description": "errCode为500 一般为服务器内部错误",
                        "schema": {
                            "$ref": "#/definitions/base_info.UpdateUserInfoResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "base_info.UpdateSelfUserInfoReq": {
            "type": "object",
            "required": [
                "operationID",
                "userID"
            ],
            "properties": {
                "birth": {
                    "type": "integer"
                },
                "email": {
                    "type": "string",
                    "maxLength": 64
                },
                "ex": {
                    "type": "string",
                    "maxLength": 1024
                },
                "faceURL": {
                    "type": "string",
                    "maxLength": 1024
                },
                "gender": {
                    "type": "integer",
                    "enum": [
                        0,
                        1,
                        2
                    ]
                },
                "nickname": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                },
                "operationID": {
                    "type": "string"
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 32
                },
                "userID": {
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1
                }
            }
        },
        "base_info.UpdateUserInfoResp": {
            "type": "object",
            "properties": {
                "errCode": {
                    "type": "integer"
                },
                "errMsg": {
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
