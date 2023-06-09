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
        "license": {
            "name": "MIT",
            "url": "https://mit-license.org"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/balances/coin": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Fetches coin balances",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trading"
                ],
                "summary": "Get coin balances",
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            }
        },
        "/balances/usd": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Fetches usd balances",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trading"
                ],
                "summary": "Get usd balances",
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            }
        },
        "/coins": {
            "get": {
                "description": "Get data for coins",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Coins"
                ],
                "summary": "Coin data",
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "500": {
                        "description": "internal server error"
                    }
                }
            }
        },
        "/login": {
            "get": {
                "description": "Provides JWT token for existing user",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "login successful"
                    },
                    "401": {
                        "description": "invalid credentials"
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Fetches past orders",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Trading"
                ],
                "summary": "Get past orders",
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Creates new buy/sell order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Trading"
                ],
                "summary": "Buy/sell coins",
                "parameters": [
                    {
                        "description": "New order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/govulnapi_models.Order"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "order went through"
                    },
                    "401": {
                        "description": "unauthorized"
                    },
                    "404": {
                        "description": "requested coin not found"
                    },
                    "500": {
                        "description": "internal server error"
                    }
                }
            }
        },
        "/register": {
            "get": {
                "description": "Registers a user",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "User registration",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "registration successful"
                    },
                    "409": {
                        "description": "email already registered or invalid parameters"
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Fetches past transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get past transactions",
                "responses": {
                    "200": {
                        "description": "ok"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Creates new transaction",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Send coins",
                "parameters": [
                    {
                        "description": "New transaction",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/govulnapi_models.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "transaction went through"
                    },
                    "400": {
                        "description": "bad request"
                    },
                    "401": {
                        "description": "unauthorized"
                    },
                    "412": {
                        "description": "not enough coin"
                    }
                }
            }
        },
        "/user/email": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Updates user email",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update email",
                "parameters": [
                    {
                        "type": "string",
                        "description": "New email",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "email updated"
                    },
                    "400": {
                        "description": "bad request"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            }
        },
        "/user/password": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Updates user password",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update password",
                "parameters": [
                    {
                        "type": "string",
                        "description": "New password",
                        "name": "password",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "password changed"
                    },
                    "400": {
                        "description": "bad request"
                    },
                    "401": {
                        "description": "unauthorized"
                    }
                }
            }
        }
    },
    "definitions": {
        "govulnapi_models.Order": {
            "type": "object",
            "properties": {
                "coinId": {
                    "type": "string",
                    "example": "bitcoin"
                },
                "isBuy": {
                    "type": "boolean"
                },
                "qty": {
                    "type": "number",
                    "example": 1
                }
            }
        },
        "govulnapi_models.Transaction": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string",
                    "example": ""
                },
                "coinId": {
                    "type": "string",
                    "example": "bitcoin"
                },
                "note": {
                    "type": "string"
                },
                "qty": {
                    "type": "number",
                    "example": 1
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"BEARER\" followed by a space and the token.",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api",
	Schemes:          []string{"http"},
	Title:            "Govulnapi",
	Description:      "Deliberately vulnerable API written in Go",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
