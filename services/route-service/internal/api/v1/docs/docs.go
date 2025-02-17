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
        "/routes": {
            "get": {
                "description": "Get a list of all routes available in the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "Retrieve all routes",
                "responses": {
                    "200": {
                        "description": "List of all routes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.RouteInfo"
                            }
                        }
                    },
                    "500": {
                        "description": "Unable to retrieve routes",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new route with the specified details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "Add a new route",
                "parameters": [
                    {
                        "description": "Create Route Request",
                        "name": "route",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RouteCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created route information",
                        "schema": {
                            "$ref": "#/definitions/dto.RouteInfo"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Unable to create route",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/routes/stops/{stopId}/schedules": {
            "get": {
                "description": "Retrieve all schedules",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "responses": {
                    "200": {
                        "description": "List of schedules",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.ScheduleResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid route ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new schedule for a specific route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Stop ID",
                        "name": "stopId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Schedule information",
                        "name": "addScheduleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddScheduleRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created schedule details",
                        "schema": {
                            "$ref": "#/definitions/dto.ScheduleResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid route ID, request format, or validation error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/routes/stops/{stopId}/schedules/{id}": {
            "get": {
                "description": "Retrieve a schedule by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Stop ID",
                        "name": "stopId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Schedule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Schedule details",
                        "schema": {
                            "$ref": "#/definitions/dto.ScheduleResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid schedule ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Schedule not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of a schedule by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Schedule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated schedule information",
                        "name": "updateScheduleRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateScheduleRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated schedule details",
                        "schema": {
                            "$ref": "#/definitions/dto.ScheduleResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid schedule ID or request format",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Schedule not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a schedule by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Schedules"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Schedule ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Invalid schedule ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Schedule not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/routes/{routeId}": {
            "get": {
                "description": "Retrieve a route by its unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "Get route by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved route",
                        "schema": {
                            "$ref": "#/definitions/dto.RouteResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid route ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Unable to retrieve route",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Update the details of an existing route by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "Update a route",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "uint",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Route Update Information",
                        "name": "route",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.RouteUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated route information",
                        "schema": {
                            "$ref": "#/definitions/dto.RouteInfo"
                        }
                    },
                    "400": {
                        "description": "Invalid route ID or request format",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to update route",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a route by its unique ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "routes"
                ],
                "summary": "Delete a route",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Route successfully deleted"
                    },
                    "400": {
                        "description": "Invalid route ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Unable to delete route",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/routes/{routeId}/stops": {
            "get": {
                "description": "Retrieve a list of stops for a given route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stops"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "stops",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.StopResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid route ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "post": {
                "description": "Add a new stop to the specified route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stops"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Stop information",
                        "name": "stop",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.AddStopRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "createdStop",
                        "schema": {
                            "$ref": "#/definitions/dto.StopResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request or incorrect data",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            }
        },
        "/routes/{routeId}/stops/{id}": {
            "get": {
                "description": "Retrieve details of a stop by its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stops"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Stop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "stop",
                        "schema": {
                            "$ref": "#/definitions/dto.StopResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid route ID or stop ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route or stop not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "put": {
                "description": "Update details of a stop for a given route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stops"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Stop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated stop information",
                        "name": "updateStopRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateStopRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "updatedStop",
                        "schema": {
                            "$ref": "#/definitions/dto.StopResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid route ID or stop ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route or stop not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a stop from a given route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stops"
                ],
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Route ID",
                        "name": "routeId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "minimum": 1,
                        "type": "integer",
                        "description": "Stop ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Invalid route ID or stop ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "404": {
                        "description": "Route or stop not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.ErrorMessage"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.AddScheduleRequest": {
            "type": "object",
            "required": [
                "arrival_time",
                "departure_time",
                "stop_id"
            ],
            "properties": {
                "arrival_time": {
                    "type": "string"
                },
                "departure_time": {
                    "type": "string"
                },
                "stop_id": {
                    "type": "integer"
                }
            }
        },
        "dto.AddStopRequest": {
            "type": "object",
            "required": [
                "name",
                "sequence"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "dto.BusResponse": {
            "type": "object",
            "properties": {
                "busCode": {
                    "type": "string"
                },
                "capacity": {
                    "type": "integer"
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
                "nextServiceDate": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "dto.RouteCreateRequest": {
            "type": "object",
            "required": [
                "endLocation",
                "name",
                "startLocation"
            ],
            "properties": {
                "endLocation": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "startLocation": {
                    "type": "string"
                }
            }
        },
        "dto.RouteInfo": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "end_location": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "route_id": {
                    "type": "integer"
                },
                "start_location": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.RouteResponse": {
            "type": "object",
            "properties": {
                "buses": {
                    "description": "Optional Buses",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.BusResponse"
                    }
                },
                "endLocation": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "startLocation": {
                    "type": "string"
                },
                "stops": {
                    "description": "Optional Stops",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.StopResponse"
                    }
                }
            }
        },
        "dto.RouteUpdateRequest": {
            "type": "object",
            "properties": {
                "endLocation": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "startLocation": {
                    "type": "string"
                }
            }
        },
        "dto.ScheduleResponse": {
            "type": "object",
            "properties": {
                "arrival_time": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "departure_time": {
                    "type": "string"
                },
                "schedule_id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "dto.StopResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "schedules": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ScheduleResponse"
                    }
                },
                "sequence": {
                    "type": "integer"
                },
                "stop_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateScheduleRequest": {
            "type": "object",
            "required": [
                "arrival_time",
                "departure_time",
                "stop_id"
            ],
            "properties": {
                "arrival_time": {
                    "type": "string"
                },
                "departure_time": {
                    "type": "string"
                },
                "stop_id": {
                    "type": "integer"
                }
            }
        },
        "dto.UpdateStopRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "sequence": {
                    "type": "integer"
                }
            }
        },
        "pkg.ErrorMessage": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "My Route Service API",
	Description:      "This API serves as an interface to interact with the My Route Service platform, providing endpoints for managing bus routes, Stops, and schedules.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
