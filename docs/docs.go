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
        "/bookings": {
            "get": {
                "description": "Returns all bookings.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get bookings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.BookingCompleteDTO"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Post a new booking.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Post booking",
                "parameters": [
                    {
                        "description": "BookingDTO JSON",
                        "name": "bookingDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/bookings/{id}": {
            "get": {
                "description": "Returns single booking.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Get booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingCompleteDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Updates a booking.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Put booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "BookingDTO JSON",
                        "name": "bookingDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "Deletes a booking.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "bookings"
                ],
                "summary": "Delete booking",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Booking Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.BookingCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
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
                                "$ref": "#/definitions/dtos.ClassCompleteDTO"
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
                        "description": "ClassDTO JSON",
                        "name": "ClassDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/classes/{id}": {
            "get": {
                "description": "Returns single class.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Get class",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ClassSchedule Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassCompleteDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "put": {
                "description": "Updates a class.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Put classes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ClassSchedule Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ClassDTO JSON",
                        "name": "ClassDTO",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "delete": {
                "description": "Deletes a class.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Delete classes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ClassSchedule Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ClassCompleteDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        },
        "/classes/{id}/bookings": {
            "get": {
                "description": "Returns the bookings of a class.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "classes"
                ],
                "summary": "Get class bookings",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ClassSchedule Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Class date in RFC3339",
                        "name": "date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.BookingCompleteDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.BookingCompleteDTO": {
            "type": "object",
            "required": [
                "bookingId",
                "classId",
                "date",
                "name"
            ],
            "properties": {
                "bookingId": {
                    "type": "integer"
                },
                "classId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.BookingDTO": {
            "type": "object",
            "required": [
                "classId",
                "date",
                "name"
            ],
            "properties": {
                "classId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dtos.ClassCompleteDTO": {
            "type": "object",
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "classId": {
                    "type": "integer"
                },
                "endDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "startDate": {
                    "type": "string"
                }
            }
        },
        "dtos.ClassDTO": {
            "type": "object",
            "required": [
                "capacity",
                "endDate",
                "name",
                "startDate"
            ],
            "properties": {
                "capacity": {
                    "type": "integer"
                },
                "endDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "startDate": {
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
