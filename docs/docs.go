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
        "/auth/login": {
            "post": {
                "description": "Logs in a registered user and returns an access token if the credentials are valid.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Login and get an access token",
                "parameters": [
                    {
                        "description": "Login request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.loginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.authResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/blogs": {
            "get": {
                "description": "get blogs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "get blogs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Query",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "minimum": 0,
                        "type": "integer",
                        "default": 0,
                        "description": "Skip",
                        "name": "skip",
                        "in": "query"
                    },
                    {
                        "minimum": 5,
                        "type": "integer",
                        "default": 5,
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blogs data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.listBlogsResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Data not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "create a new blog",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "create blog",
                "parameters": [
                    {
                        "description": "Create blog request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.createBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blog created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.blogResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/blogs/{id}": {
            "get": {
                "description": "get blog by blog id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "get blog",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blog data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.blogResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Data not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update a blog data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "update blog",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update blog request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.putBlogRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blog updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.blogResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete a blog",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "blogs"
                ],
                "summary": "delete blog",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Blog id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Blog updated",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "create an new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "create user",
                "parameters": [
                    {
                        "description": "Create User request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.createUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.userResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "get a user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "get user",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User data",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.userResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Data not found error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update User request body",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.updateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User updated",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/handler.response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/handler.userResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "409": {
                        "description": "Data conflict error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "delete user by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "delete user",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "User id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted",
                        "schema": {
                            "$ref": "#/definitions/handler.response"
                        }
                    },
                    "400": {
                        "description": "Validation error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.authResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string",
                    "example": "eyJJ9.eyJpEzNDR9.fUjDw0"
                }
            }
        },
        "handler.blogResponse": {
            "type": "object",
            "properties": {
                "author_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string",
                    "example": "1970-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "example": "39833b12-a044-46f5-8abd-47c47345d458"
                },
                "text": {
                    "type": "string",
                    "example": "to do ..."
                },
                "title": {
                    "type": "string",
                    "example": "how to ..."
                },
                "updated_at": {
                    "type": "string",
                    "example": "1970-01-01T00:00:00Z"
                }
            }
        },
        "handler.createBlogRequest": {
            "type": "object",
            "required": [
                "text",
                "title"
            ],
            "properties": {
                "text": {
                    "type": "string",
                    "example": "adaw ..."
                },
                "title": {
                    "type": "string",
                    "example": "adw..."
                }
            }
        },
        "handler.createUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "12345678"
                },
                "username": {
                    "type": "string",
                    "minLength": 3,
                    "example": "laplala"
                }
            }
        },
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "data not found"
                    ]
                },
                "success": {
                    "type": "boolean",
                    "example": false
                }
            }
        },
        "handler.listBlogsResponse": {
            "type": "object",
            "properties": {
                "blogs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/handler.blogResponse"
                    }
                },
                "meta": {
                    "$ref": "#/definitions/handler.meta"
                }
            }
        },
        "handler.loginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "12345678"
                },
                "username": {
                    "type": "string",
                    "minLength": 3,
                    "example": "laplala"
                }
            }
        },
        "handler.meta": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer",
                    "example": 10
                },
                "skip": {
                    "type": "integer",
                    "example": 0
                },
                "total": {
                    "type": "integer",
                    "example": 100
                }
            }
        },
        "handler.putBlogRequest": {
            "type": "object",
            "required": [
                "text",
                "title"
            ],
            "properties": {
                "text": {
                    "type": "string",
                    "example": "adaw ..."
                },
                "title": {
                    "type": "string",
                    "example": "adw..."
                }
            }
        },
        "handler.response": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string",
                    "example": "Success"
                },
                "success": {
                    "type": "boolean",
                    "example": true
                }
            }
        },
        "handler.updateUserRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "12345678"
                },
                "username": {
                    "type": "string",
                    "minLength": 3,
                    "example": "laplala"
                }
            }
        },
        "handler.userResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "1970-01-01T00:00:00Z"
                },
                "id": {
                    "type": "string",
                    "example": "39833b12-a044-46f5-8abd-47c47345d458"
                },
                "updated_at": {
                    "type": "string",
                    "example": "1970-01-01T00:00:00Z"
                },
                "username": {
                    "type": "string",
                    "example": "laplala"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "Type \"Bearer\" followed by a space and the access token.",
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
	BasePath:         "/v1/api",
	Schemes:          []string{"http", "https"},
	Title:            "Go BLOG API",
	Description:      "This is a simple RESTful blog api.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
