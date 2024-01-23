// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Luís Almas",
            "email": "la_luisalmas@hotmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/classes": {
            "get": {
                "description": "Returns all scheduled classes.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Get classes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entities.ClassSchedule"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post a new class.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Post class",
                "parameters": [
                    {
                        "description": "ClassScheduleDTO JSON",
                        "name": "classScheduleDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassScheduleDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/entities.ClassSchedule"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ClassScheduleDTO": {
            "type": "object",
            "required": [
                "capacity",
                "end_date",
                "name",
                "start_date"
            ],
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "end_date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        },
        "entities.ClassSchedule": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "end_date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "start_date": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Gym management API",
	Description:      "A book management service API in Go using Gin framework.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
