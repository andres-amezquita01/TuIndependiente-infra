// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/greeter": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Greeter service",
                "operationId": "1",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Input name",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.GreeterResponse"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Health service",
                "operationId": "7",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.GreeterResponse"
                        }
                    }
                }
            }
        },
        "/service/cancel": {
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Cancel service",
                "operationId": "4",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "serviceID",
                        "name": "serviceID",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ServiceResponse"
                        }
                    }
                }
            }
        },
        "/service/create": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create service",
                "operationId": "2",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "address",
                        "name": "address",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.ServiceResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.GreeterResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "main.ServiceResponse": {
            "type": "object",
            "properties": {
                "result": {
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