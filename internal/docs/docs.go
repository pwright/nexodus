// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "The Nexodus Authors",
            "url": "https://github.com/nexodus-io/nexodus/issues"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/devices": {
            "get": {
                "description": "Lists all devices",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "List Devices",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Device"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new device",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Add Devices",
                "parameters": [
                    {
                        "description": "Add Device",
                        "name": "device",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddDevice"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/devices/{id}": {
            "get": {
                "description": "Gets a device by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Get Devices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing device and associated IPAM lease",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Delete Device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "patch": {
                "description": "Updates a device by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Update Devices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Device Update",
                        "name": "update",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UpdateDevice"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Device"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/fflags": {
            "get": {
                "description": "Lists all feature flags",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FFlag"
                ],
                "summary": "List Feature Flags",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/fflags/{name}": {
            "get": {
                "description": "Gets a Feature Flag by name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "FFlag"
                ],
                "summary": "Get Feature Flag",
                "parameters": [
                    {
                        "type": "string",
                        "description": "feature flag name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "boolean"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/organizations": {
            "get": {
                "description": "Lists all Organizations",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "List Organizations",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Organization"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a named organization with the given CIDR",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Create an Organization",
                "parameters": [
                    {
                        "description": "Add Organization",
                        "name": "Organization",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.AddOrganization"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Organization"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/organizations/{id}": {
            "get": {
                "description": "Gets a Organization by Organization ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organization"
                ],
                "summary": "Get Organizations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Organization"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes an existing organization and associated IPAM prefix",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Organizations"
                ],
                "summary": "Delete Organization",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.Organization"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "405": {
                        "description": "Method Not Allowed",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/organizations/{id}/devices": {
            "get": {
                "description": "Lists all devices for this Organization",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "List Devices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Device"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/organizations/{organization_id}/devices/{device_id}": {
            "get": {
                "description": "Gets a device in a organization by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Devices"
                ],
                "summary": "Get Device",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Device ID",
                        "name": "device_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Device"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Lists all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "List Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Gets a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "429": {
                        "description": "Too Many Requests",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        },
        "/users/{id}/organizations/{organization}": {
            "delete": {
                "description": "Deletes an existing organization associated to a user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Remove a User from an Organization",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Organization ID",
                        "name": "organization",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.BaseError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.AddDevice": {
            "type": "object"
        },
        "models.AddOrganization": {
            "type": "object",
            "properties": {
                "cidr": {
                    "type": "string",
                    "example": "172.16.42.0/24"
                },
                "description": {
                    "type": "string",
                    "example": "The Red Zone"
                },
                "hub_zone": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string",
                    "example": "zone-red"
                }
            }
        },
        "models.BaseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "something bad"
                }
            }
        },
        "models.Device": {
            "type": "object",
            "properties": {
                "allowed_ips": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "child_prefix": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "discovery": {
                    "type": "boolean"
                },
                "endpoint_local_address_ip4": {
                    "type": "string"
                },
                "hostname": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "local_ip": {
                    "type": "string"
                },
                "organization_id": {
                    "type": "string"
                },
                "organization_prefix": {
                    "type": "string"
                },
                "public_key": {
                    "type": "string"
                },
                "reflexive_ip4": {
                    "type": "string"
                },
                "relay": {
                    "type": "boolean"
                },
                "symmetric_nat": {
                    "type": "boolean"
                },
                "tunnel_ip": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.Organization": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "devices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Device"
                    }
                },
                "hubZone": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "ipCidr": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.User"
                    }
                }
            }
        },
        "models.UpdateDevice": {
            "type": "object",
            "properties": {
                "child_prefix": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "172.16.42.0/24"
                    ]
                },
                "endpoint_local_address_ip4": {
                    "type": "string",
                    "example": "1.2.3.4"
                },
                "hostname": {
                    "type": "string",
                    "example": "myhost"
                },
                "local_ip": {
                    "type": "string",
                    "example": "10.1.1.1"
                },
                "organization_id": {
                    "type": "string",
                    "example": "694aa002-5d19-495e-980b-3d8fd508ea10"
                },
                "reflexive_ip4": {
                    "type": "string"
                },
                "symmetric_nat": {
                    "type": "boolean"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "devices": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Device"
                    }
                },
                "id": {
                    "description": "Since the ID comes from the IDP, we have no control over the format...",
                    "type": "string",
                    "example": "aa22666c-0f57-45cb-a449-16efecc04f2e"
                },
                "organizations": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Organization"
                    }
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Implicit": {
            "type": "oauth2",
            "flow": "implicit",
            "authorizationUrl": "https://auth.try.nexodus.local/",
            "scopes": {
                "admin": " Grants read and write access to administrative information",
                "user": " Grants read and write access to resources owned by this user"
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Nexodus API",
	Description:      "This is the Nexodus API Server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
