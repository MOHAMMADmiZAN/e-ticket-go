// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
            "name": "Mohammad Mizan",
            "url": "http://swagger.io/support",
            "email": "takbir.jcd@gmail.com"
        },
        "license": {
            "name": "Apache License Version 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/buses": {
            "get": {
                "description": "Retrieve a list of all buses from the bus service.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Get all buses",
                "responses": {
                    "200": {
                        "description": "List of all buses",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BusResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Could not retrieve buses",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new bus to the system with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Create a new bus",
                "parameters": [
                    {
                        "description": "Bus Creation Data",
                        "name": "bus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateBusDTO"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created bus",
                        "schema": {
                            "$ref": "#/definitions/dto.BusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid input provided",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to create bus",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    }
                }
            }
        },
        "/buses/{id}": {
            "get": {
                "description": "Retrieve details of a specific bus by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Get bus by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved bus",
                        "schema": {
                            "$ref": "#/definitions/dto.BusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid ID format or Bus not found",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of a specific bus by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Update bus details",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Bus Update Data",
                        "name": "bus",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBusDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bus updated successfully",
                        "schema": {
                            "$ref": "#/definitions/dto.BusResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid ID or update data format",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - Bus not found",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Could not update bus",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete the bus with the specified ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Delete a bus",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bus deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid bus ID",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - Bus not found",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to delete bus",
                        "schema": {
                            "$ref": "#/definitions/handler.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BusResponse": {
            "type": "object",
            "properties": {
                "busCode": {
                    "type": "string"
                },
                "capacity": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastServiceDate": {
                    "type": "string"
                },
                "licensePlate": {
                    "type": "string"
                },
                "makeModel": {
                    "type": "string"
                },
                "nextServiceDate": {
                    "type": "string"
                },
                "routeId": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "year": {
                    "type": "integer"
                }
            }
        },
        "dto.CreateBusDTO": {
            "type": "object",
            "required": [
                "busCode",
                "capacity",
                "lastServiceDate",
                "licensePlate",
                "makeModel",
                "nextServiceDate",
                "routeId",
                "status",
                "year"
            ],
            "properties": {
                "busCode": {
                    "type": "string",
                    "maxLength": 100
                },
                "capacity": {
                    "type": "integer"
                },
                "lastServiceDate": {
                    "type": "string"
                },
                "licensePlate": {
                    "type": "string",
                    "maxLength": 20
                },
                "makeModel": {
                    "type": "string",
                    "maxLength": 100
                },
                "nextServiceDate": {
                    "type": "string"
                },
                "routeId": {
                    "type": "integer"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "active",
                        "inactive",
                        "in_service",
                        "out_of_service"
                    ]
                },
                "year": {
                    "description": "Assuming year range from 1900 to 2100",
                    "type": "integer"
                }
            }
        },
        "dto.UpdateBusDTO": {
            "type": "object",
            "properties": {
                "busCode": {
                    "type": "string",
                    "maxLength": 100
                },
                "capacity": {
                    "type": "integer"
                },
                "lastServiceDate": {
                    "type": "string"
                },
                "licensePlate": {
                    "type": "string",
                    "maxLength": 20
                },
                "makeModel": {
                    "type": "string",
                    "maxLength": 100
                },
                "nextServiceDate": {
                    "type": "string"
                },
                "routeId": {
                    "type": "integer"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "active",
                        "inactive",
                        "in_service",
                        "out_of_service"
                    ]
                },
                "year": {
                    "description": "Assuming year range from 1900 to 2100",
                    "type": "integer"
                }
            }
        },
        "handler.APIResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "My Bus Service API",
	Description:      "This API serves as an interface to interact with the My Bus Service platform, providing endpoints for managing bus routes, bookings, and user interactions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}