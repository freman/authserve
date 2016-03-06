package main

var swaggerDefinition = []byte(`{
	"swagger": "2.0",
	"info": {
		"description": "Authentication microservice",
		"title": "AuthService",
		"version": "1.0.0"
	},
	"consumes": [
		"application/json"
	],
	"produces": [
		"application/json"
	],
	"schemes": [
		"http"
	],
	"paths": {
		"/authenticate": {
			"post": {
				"tags": [
					"auth"
				],
				"description": "Authenticates a user and returns a token",
				"security": [],
				"parameters": [
					{
						"name": "body",
						"in": "body",
						"schema": {
							"$ref": "#/definitions/authentication"
						}
					}
				],
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/token"
						}
					},
					"401": {
						"description": "Unauthorized",
						"schema": {
							"$ref": "#/definitions/error"
						}
					},
					"default": {
						"description": "error",
						"schema": {
							"$ref": "#/definitions/error"
						}
					}
				}
			}
		},
		"/new": {
			"post": {
				"tags": [
					"auth"
				],
				"description": "Creates a user and returns a token",
				"security": [],
				"parameters": [
					{
						"name": "body",
						"in": "body",
						"schema": {
							"$ref": "#/definitions/create"
						}
					}
				],
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/token"
						}
					},
					"default": {
						"description": "error",
						"schema": {
							"$ref": "#/definitions/error"
						}
					}
				}
			}
		},
		"/refresh": {
			"get": {
				"tags": [
					"auth"
				],
				"description": "Refresh token with new TTL - no db access",
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/token"
						}
					},
					"401": {
						"description": "Unauthorized",
						"schema": {
							"$ref": "#/definitions/error"
						}
					},
					"default": {
						"description": "error",
						"schema": {
							"$ref": "#/definitions/error"
						}
					}
				}
			}
		},
		"/reload": {
			"get": {
				"tags": [
					"auth"
				],
				"description": "Reload user details from DB",
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/token"
						}
					},
					"401": {
						"description": "Unauthorized",
						"schema": {
							"$ref": "#/definitions/error"
						}
					},
					"default": {
						"description": "error",
						"schema": {
							"$ref": "#/definitions/error"
						}
					}
				}
			}
		},
		"/password": {
			"post": {
				"tags": [
					"auth"
				],
				"description": "Change the users password",
				"parameters": [
					{
						"name": "body",
						"in": "body",
						"schema": {
							"$ref": "#/definitions/password"
						}
					}
				],
				"responses": {
					"200": {
						"description": "OK",
						"schema": {
							"$ref": "#/definitions/token"
						}
					},
					"401": {
						"description": "Unauthorized",
						"schema": {
							"$ref": "#/definitions/error"
						}
					},
					"default": {
						"description": "error",
						"schema": {
							"$ref": "#/definitions/error"
						}
					}
				}
			}
		}
	},
	"securityDefinitions": {
		"jwtAuth": {
			"description": "JSON Web Token",
			"type": "apiKey",
			"in": "header",
			"name": "Authorization"
		}
	},
	"security": [
		{
			"jwtAuth": []
		}
	],
	"definitions": {
		"error": {
			"type": "object",
			"required": [
				"message"
			],
			"properties": {
				"code": {
					"type": "integer",
					"format": "int64"
				},
				"message": {
					"type": "string"
				}
			}
		},
		"create": {
			"type": "object",
			"description": "New user details",
			"required": [
				"email",
				"firstname",
				"lastname",
				"password"
			],
			"properties": {
				"email": {
					"type": "string",
					"minLength": 3
				},
				"firstname": {
					"type": "string"
				},
				"lastname": {
					"type": "string"
				},
				"password": {
					"type": "string",
					"format": "password",
					"minLength": 8
				}
			}
		},
		"authentication": {
			"type": "object",
			"description": "Authentication credentials",
			"required": [
				"email",
				"password"
			],
			"properties": {
				"email": {
					"type": "string",
					"minLength": 3
				},
				"password": {
					"type": "string",
					"format": "password",
					"minLength": 8
				}
			}
		},
		"password": {
			"type": "object",
			"description": "New password",
			"required": [
				"password"
			],
			"properties": {
				"password": {
					"type": "string",
					"format": "password",
					"minLength": 8
				}
			}
		},
		"token": {
			"type": "object",
			"description": "JSON Web Token",
			"required": [
				"token"
			],
			"properties": {
				"token": {
					"type": "string"
				}
			}
		}
	}
}`)
