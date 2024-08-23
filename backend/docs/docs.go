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
        "/private/admin/user": {
            "post": {
                "description": "Creates User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Creates User",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserDTO"
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
        "/private/admin/user/{ID}": {
            "get": {
                "description": "Retrieves User Profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get Profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/private/admin/users": {
            "get": {
                "description": "Retrieves All Users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.BaseResponse"
                        }
                    }
                }
            }
        },
        "/private/home/advancement": {
            "get": {
                "description": "Get User Advancement",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetUserAdvancement",
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
        "/private/home/development": {
            "get": {
                "description": "Get User Development",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetUserDevelopment",
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
        "/private/home/inventories": {
            "get": {
                "description": "Get Inventories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetInventories",
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
        "/private/home/lab": {
            "get": {
                "description": "Get Lab Content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetLabContent",
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
        "/private/home/level": {
            "get": {
                "description": "Get User Level",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetUserLevel",
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
        "/private/home/road": {
            "get": {
                "description": "Get Road Content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetRoadContent",
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
        "/private/home/welcome": {
            "get": {
                "description": "Get Welcome Content",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Home"
                ],
                "summary": "GetWelcomeContent",
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
        "/private/lab/{programmingID}/{labID}": {
            "get": {
                "description": "Get Lab By Programming Language ID \u0026 Lab ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lab"
                ],
                "summary": "GetLabByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Programming Language ID",
                        "name": "programmingID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Lab ID",
                        "name": "labID",
                        "in": "path",
                        "required": true
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
        "/private/labs/stats/{language}/{userID}": {
            "get": {
                "description": "Get User Programming Language Lab Statistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lab"
                ],
                "summary": "GetUserProgrammingLanguageLabStats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Language",
                        "name": "language",
                        "in": "path",
                        "required": true
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
        "/private/labs/stats/{userID}": {
            "get": {
                "description": "Get User General Lab Statistics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lab"
                ],
                "summary": "GetUserGeneralLabStats",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
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
        "/private/labs/{ID}": {
            "get": {
                "description": "Get Labs By Programming Language ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lab"
                ],
                "summary": "GetLabsById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Programming Language ID",
                        "name": "ID",
                        "in": "path",
                        "required": true
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
        "/private/log": {
            "get": {
                "description": "Retrieves all logs based on the provided query parameters.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Get all logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "userID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Programming ID",
                        "name": "programmingID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Log Lab or Path ID",
                        "name": "labRoadID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Log Content",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Log Type",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.LogDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/private/log/solution/byday": {
            "get": {
                "description": "Retrieves the number of lab and road solutions solved day by day.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "GetSolutionsByDay",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.SolutionsByDayDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/private/log/solution/hours": {
            "get": {
                "description": "Retrieves the total hours spent on lab and road solutions for each programming language in the last week.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "GetSolutionsHoursByProgramming",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.SolutionsHoursByProgrammingDTO"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/private/road/start": {
            "post": {
                "description": "Start",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Road"
                ],
                "summary": "Start",
                "parameters": [
                    {
                        "description": "Start",
                        "name": "start",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.StartDTO"
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
        "/private/road/{roadID}": {
            "get": {
                "description": "Get Road with Paths",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Road"
                ],
                "summary": "GetRoads",
                "parameters": [
                    {
                        "type": "string",
                        "description": "roadID",
                        "name": "roadID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.RoadDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/private/road/{roadID}/{pathID}": {
            "get": {
                "description": "Get Path By ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Road"
                ],
                "summary": "GetPathByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Road ID",
                        "name": "roadID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Path ID",
                        "name": "pathID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.PathDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/private/user/": {
            "get": {
                "description": "Gets users profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "GetProfile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserDTO"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "Updates user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "UpdateUser",
                "parameters": [
                    {
                        "description": "UpdateUser",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateUserDTO"
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
        "/private/user/password": {
            "put": {
                "description": "Updates users password",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "UpdatePassword",
                "parameters": [
                    {
                        "description": "UpdatePassword",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdatePasswordDTO"
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
                            "$ref": "#/definitions/dto.LoginDTO"
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
                            "$ref": "#/definitions/dto.RegisterDTO"
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
        "dto.CreateUserDTO": {
            "type": "object",
            "required": [
                "name",
                "password",
                "role",
                "surname",
                "username"
            ],
            "properties": {
                "githubProfile": {
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
                "role": {
                    "type": "string"
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
        "dto.LanguageDTO": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "lang": {
                    "type": "string"
                },
                "note": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.LogDTO": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "labPathID": {
                    "type": "integer"
                },
                "logType": {
                    "type": "string"
                },
                "programmingID": {
                    "type": "integer"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "dto.LoginDTO": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "dto.PathDTO": {
            "type": "object",
            "properties": {
                "difficulty": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isFinished": {
                    "type": "boolean"
                },
                "isStarted": {
                    "type": "boolean"
                },
                "languages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.LanguageDTO"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.RegisterDTO": {
            "type": "object",
            "required": [
                "name",
                "password",
                "surname",
                "username"
            ],
            "properties": {
                "githubProfile": {
                    "type": "string",
                    "maxLength": 30
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "surname": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "dto.RoadDTO": {
            "type": "object",
            "properties": {
                "iconPath": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "paths": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.PathDTO"
                    }
                }
            }
        },
        "dto.SolutionsByDayDTO": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "labCount": {
                    "type": "integer"
                },
                "roadCount": {
                    "type": "integer"
                }
            }
        },
        "dto.SolutionsHoursByProgrammingDTO": {
            "type": "object",
            "properties": {
                "labHours": {
                    "type": "number"
                },
                "programmingID": {
                    "type": "integer"
                },
                "roadHours": {
                    "type": "number"
                }
            }
        },
        "dto.StartDTO": {
            "type": "object",
            "required": [
                "programmingID"
            ],
            "properties": {
                "programmingID": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdatePasswordDTO": {
            "type": "object",
            "required": [
                "confirmPassword",
                "newPassword",
                "password"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string",
                    "minLength": 8
                },
                "newPassword": {
                    "type": "string",
                    "minLength": 8
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateUserDTO": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "githubProfile": {
                    "type": "string",
                    "maxLength": 30
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "username": {
                    "type": "string",
                    "maxLength": 30,
                    "minLength": 3
                }
            }
        },
        "dto.UserDTO": {
            "type": "object",
            "properties": {
                "bestLanguage": {
                    "type": "string"
                },
                "githubProfile": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
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
