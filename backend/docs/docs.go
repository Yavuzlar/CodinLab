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
        "/public/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "login",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/public.LoginDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/public/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/public.RegisterDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "public.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Password is required and must be at least 8 characters",
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "description": "Username is required, must be alphanumeric and between 3-30 characters",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "public.RegisterDTO": {
            "type": "object",
            "required": [
                "name",
                "password",
                "surname",
                "username"
            ],
            "properties": {
                "githubProfile": {
                    "description": "Github Profile is must be max 30 characters long.",
                    "type": "string",
                    "maxLength": 30
                },
                "name": {
                    "description": "Name is required",
                    "type": "string"
                },
                "password": {
                    "description": "Password is required and must be at least 8 characters",
                    "type": "string",
                    "minLength": 8
                },
                "surname": {
                    "description": "Surname is required",
                    "type": "string"
                },
                "username": {
                    "description": "Username is required, must be alphanumeric and between 3-30 characters",
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "response.BaseResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "data_count": {
                    "type": "integer"
                },
                "errors": {},
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "session_id",
            "in": "cookie"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "API Service",
	Description:      "API Service for CodinLab",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
