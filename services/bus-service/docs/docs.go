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
        "/": {
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
                            "$ref": "#/definitions/pkg.APIResponse"
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
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to create bus",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/routes/{routeId}": {
            "get": {
                "description": "Retrieve a list of all buses that operate on a specific route.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Get all buses by route ID",
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
                        "description": "List of all buses",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.BusResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid route ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/status": {
            "get": {
                "description": "Retrieve a list of all buses with the specified status.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Get all buses by status",
                "parameters": [
                    {
                        "enum": [
                            "active",
                            "maintenance",
                            "decommissioned"
                        ],
                        "type": "string",
                        "description": "status",
                        "name": "status",
                        "in": "query",
                        "required": true
                    }
                ],
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
                    "400": {
                        "description": "Bad Request - Invalid status",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busID}/seats/": {
            "get": {
                "description": "Retrieves the seat inventory for a specific bus",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "busID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of Seat objects",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.SeatResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new seat to a specific bus",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "busID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Seat Data",
                        "name": "seat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateSeatRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created seat",
                        "schema": {
                            "$ref": "#/definitions/dto.SeatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busID}/seats/availability": {
            "get": {
                "description": "Retrieves all seats that are currently available",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "responses": {
                    "200": {
                        "description": "Array of available seats",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.SeatResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busID}/seats/status/{status}": {
            "get": {
                "description": "Retrieves all seats with a specific status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "enum": [
                            "booked",
                            "available",
                            "reserved"
                        ],
                        "type": "string",
                        "description": "Seat Status",
                        "name": "status",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Array of seats with the specified status",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.SeatResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busID}/seats/{id}": {
            "get": {
                "description": "Retrieves details of a specific seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully retrieved seat",
                        "schema": {
                            "$ref": "#/definitions/dto.SeatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the details or status of a specific seat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Seat update data",
                        "name": "seat",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSeatRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated seat",
                        "schema": {
                            "$ref": "#/definitions/dto.SeatResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Removes a seat from the system",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully deleted seat",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busID}/seats/{id}/status": {
            "put": {
                "description": "Updates the status of a seat (booked, available, reserved)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "seats"
                ],
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Seat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Seat status data",
                        "name": "status",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateSeatRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully updated seat status",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busId}": {
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
                        "name": "busId",
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
                            "$ref": "#/definitions/pkg.APIResponse"
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
                        "name": "busId",
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
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - Bus not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Could not update bus",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
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
                        "name": "busId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bus deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid bus ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found - Bus not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Unable to delete bus",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{busId}/service-dates": {
            "put": {
                "description": "Update the last and next service dates for a specific bus.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "buses"
                ],
                "summary": "Update bus service dates",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Bus ID",
                        "name": "busId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Service Dates",
                        "name": "serviceDates",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UpdateBusServiceDatesDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Bus service dates updated successfully",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request - Invalid bus ID or service dates",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
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
                        "maintenance",
                        "decommissioned"
                    ]
                },
                "year": {
                    "description": "Assuming year range from 1900 to 2100",
                    "type": "integer"
                }
            }
        },
        "dto.CreateSeatRequest": {
            "type": "object",
            "required": [
                "bus_id",
                "class_type",
                "seat_number",
                "seat_status"
            ],
            "properties": {
                "bus_id": {
                    "type": "integer"
                },
                "class_type": {
                    "enum": [
                        "Regular",
                        "Business"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SeatClassType"
                        }
                    ]
                },
                "is_available": {
                    "description": "Optional, defaults to true if not specified.",
                    "type": "boolean"
                },
                "seat_number": {
                    "type": "string"
                },
                "seat_status": {
                    "enum": [
                        "Booked",
                        "Available",
                        "Reserved"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SeatStatus"
                        }
                    ]
                }
            }
        },
        "dto.SeatResponse": {
            "type": "object",
            "properties": {
                "bus_id": {
                    "type": "integer"
                },
                "class_type": {
                    "$ref": "#/definitions/models.SeatClassType"
                },
                "id": {
                    "type": "integer"
                },
                "is_available": {
                    "type": "boolean"
                },
                "seat_number": {
                    "type": "string"
                },
                "seat_status": {
                    "$ref": "#/definitions/models.SeatStatus"
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
                        "maintenance",
                        "decommissioned"
                    ]
                },
                "year": {
                    "description": "Assuming year range from 1900 to 2100",
                    "type": "integer"
                }
            }
        },
        "dto.UpdateBusServiceDatesDTO": {
            "type": "object",
            "required": [
                "lastServiceDate",
                "nextServiceDate"
            ],
            "properties": {
                "lastServiceDate": {
                    "type": "string"
                },
                "nextServiceDate": {
                    "type": "string"
                }
            }
        },
        "dto.UpdateSeatRequest": {
            "type": "object",
            "properties": {
                "bus_id": {
                    "description": "Use pointers to differentiate between zero value and omitted field.",
                    "type": "integer"
                },
                "class_type": {
                    "enum": [
                        "Regular",
                        "Business"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SeatClassType"
                        }
                    ]
                },
                "is_available": {
                    "description": "Optional, can be omitted.",
                    "type": "boolean"
                },
                "seat_number": {
                    "type": "string"
                },
                "seat_status": {
                    "enum": [
                        "Booked",
                        "Available",
                        "Reserved"
                    ],
                    "allOf": [
                        {
                            "$ref": "#/definitions/models.SeatStatus"
                        }
                    ]
                }
            }
        },
        "models.SeatClassType": {
            "type": "string",
            "enum": [
                "Regular",
                "Business"
            ],
            "x-enum-varnames": [
                "ClassRegular",
                "ClassBusiness"
            ]
        },
        "models.SeatStatus": {
            "type": "string",
            "enum": [
                "Booked",
                "Available",
                "Reserved"
            ],
            "x-enum-varnames": [
                "StatusBooked",
                "StatusAvailable",
                "StatusReserved"
            ]
        },
        "pkg.APIResponse": {
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
	Host:             "localhost:8082",
	BasePath:         "/api/v1/buses",
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
