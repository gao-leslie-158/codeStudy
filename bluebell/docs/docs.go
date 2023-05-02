// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "gaoleslei@gmail.com",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/post": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户必须登录才能创建帖子",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "用户创建帖子接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "创建帖子参数",
                        "name": "p",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Post"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/post/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户必须登录才能获取帖子详情",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "根据post_id获取帖子详情的接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "帖子id",
                        "name": "object",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户必须登录后根据page和size请求参数分页获取帖子详情列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "根据page，size获取帖子详情的接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "帖子起始页page和页数大小size,有默认值",
                        "name": "object",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        },
        "/posts2": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户必须登录后根据需求获取帖子详情列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关接口"
                ],
                "summary": "根据时间、帖子投票数或社区分类获取帖子详情的接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token令牌",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "可以为空",
                        "name": "community_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller._ResponsePostList"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller._ResponsePostList": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务响应状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.ApiPostDetail"
                    }
                },
                "message": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        },
        "models.ApiPostDetail": {
            "type": "object",
            "properties": {
                "community_detail": {
                    "description": "嵌入社区结构体",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.CommunityDetail"
                        }
                    ]
                },
                "post_details": {
                    "description": "嵌入帖子结构体",
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.Post"
                        }
                    ]
                },
                "user_name": {
                    "type": "string"
                },
                "vote_neg_num": {
                    "type": "string",
                    "example": "0"
                },
                "vote_pos_num": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "models.CommunityDetail": {
            "type": "object",
            "properties": {
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "introduction": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Post": {
            "type": "object",
            "required": [
                "content",
                "title"
            ],
            "properties": {
                "author_id": {
                    "type": "integer"
                },
                "community_id": {
                    "type": "integer"
                },
                "content": {
                    "type": "string"
                },
                "create_time": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "0"
                },
                "status": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "8082",
	BasePath:         "[ Base URL: 127.0.0.1:8082/api/v1 ]",
	Schemes:          []string{},
	Title:            "bluebell项目文档接口",
	Description:      "一个简单的帖子社区",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}