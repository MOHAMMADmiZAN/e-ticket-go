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
        "/create": {
            "post": {
                "description": "Creates a new user profile with the provided details.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Create user profile",
                "parameters": [
                    {
                        "description": "Create Profile Request",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfileRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Profile created successfully",
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid profile data",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        },
        "/{userID}": {
            "get": {
                "description": "Retrieves the user profile for a specified user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Get user profile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile fetched successfully",
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "404": {
                        "description": "Profile not found",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Updates the profile details for a given user ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Profile Request",
                        "name": "profile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfileUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Profile updated successfully",
                        "schema": {
                            "$ref": "#/definitions/dto.UserProfileResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid data",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes the profile of the specified user ID.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Delete user profile",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Profile deleted successfully"
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/pkg.APIResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserProfileRequest": {
            "type": "object",
            "required": [
                "dateOfBirth",
                "firstName",
                "lastName",
                "userID"
            ],
            "properties": {
                "dateOfBirth": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "profilePictureURL": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "dto.UserProfileResponse": {
            "type": "object",
            "properties": {
                "dateOfBirth": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "profilePictureURL": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "dto.UserProfileUpdate": {
            "type": "object",
            "properties": {
                "dateOfBirth": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "profilePictureURL": {
                    "type": "string"
                }
            }
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
	Host:             "localhost:8084",
	BasePath:         "/api/v1/profiles",
	Schemes:          []string{},
	Title:            "My Profile Service API",
	Description:      "This API serves as an interface to interact with the My Bus Service platform, providing endpoints for managing bus routes, bookings, and user interactions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
