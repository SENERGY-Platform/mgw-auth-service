// Package swagger_docs Code generated by swaggo/swag. DO NOT EDIT
package swagger_docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache-2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/identities": {
            "get": {
                "description": "List all mgw-core users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "Get identities",
                "parameters": [
                    {
                        "type": "string",
                        "description": "filter by identity type",
                        "name": "type",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "identities",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/model.Identity"
                            }
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new mgw-core user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "Create identity",
                "parameters": [
                    {
                        "description": "user info",
                        "name": "identity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NewIdentityRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "identity ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/identities/{id}": {
            "get": {
                "description": "Get a mgw-core user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "Get identity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "identity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "identity info",
                        "schema": {
                            "$ref": "#/definitions/model.Identity"
                        }
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Remove a mgw-core user.",
                "tags": [
                    "Identities"
                ],
                "summary": "Delete identity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "identity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Change mgw-core user information or password.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Identities"
                ],
                "summary": "Update identity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "identity ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "identity info",
                        "name": "identity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateIdentityRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/info": {
            "get": {
                "description": "Get basic service and runtime information.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service Information"
                ],
                "summary": "Get stats",
                "responses": {
                    "200": {
                        "description": "info",
                        "schema": {
                            "$ref": "#/definitions/lib.SrvInfo"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pairing/close": {
            "patch": {
                "description": "Close paring endpoint and cancel paring session.",
                "tags": [
                    "Device Pairing"
                ],
                "summary": "Stop pairing",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pairing/open": {
            "patch": {
                "description": "Open paring endpoint and create a paring session.",
                "tags": [
                    "Device Pairing"
                ],
                "summary": "Start paring",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "set session duration in nanoseconds (default=5m)",
                        "name": "duration",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/pairing/request": {
            "post": {
                "description": "Transmit device information to pair a device.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Device Pairing"
                ],
                "summary": "Pair device",
                "parameters": [
                    {
                        "description": "device information",
                        "name": "meta",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "generated credentials",
                        "schema": {
                            "$ref": "#/definitions/model.CredentialsResponse"
                        }
                    },
                    "400": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error message",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "lib.MemStats": {
            "type": "object",
            "properties": {
                "alloc": {
                    "type": "integer"
                },
                "alloc_total": {
                    "type": "integer"
                },
                "gc_cycles": {
                    "type": "integer"
                },
                "sys_total": {
                    "type": "integer"
                }
            }
        },
        "lib.SrvInfo": {
            "type": "object",
            "properties": {
                "mem_stats": {
                    "$ref": "#/definitions/lib.MemStats"
                },
                "name": {
                    "type": "string"
                },
                "up_time": {
                    "$ref": "#/definitions/time.Duration"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "model.CredentialsResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "login": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "model.Identity": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "meta": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "type": {
                    "$ref": "#/definitions/model.IdentityType"
                },
                "updated": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.IdentityType": {
            "type": "string",
            "enum": [
                "human",
                "machine"
            ],
            "x-enum-varnames": [
                "HumanType",
                "MachineType"
            ]
        },
        "model.NewIdentityRequest": {
            "type": "object",
            "properties": {
                "meta": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "secret": {
                    "type": "string"
                },
                "type": {
                    "$ref": "#/definitions/model.IdentityType"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.UpdateIdentityRequest": {
            "type": "object",
            "properties": {
                "meta": {
                    "type": "object",
                    "additionalProperties": {}
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "time.Duration": {
            "type": "integer",
            "enum": [
                1,
                1000,
                1000000,
                1000000000
            ],
            "x-enum-varnames": [
                "Nanosecond",
                "Microsecond",
                "Millisecond",
                "Second"
            ]
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.2.13",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Auth Service API",
	Description:      "Provides access to mgw-core auth functions.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
