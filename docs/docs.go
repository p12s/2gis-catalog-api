// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/building/": {
            "post": {
                "description": "Add building into database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Building"
                ],
                "summary": "Add building",
                "operationId": "add",
                "parameters": [
                    {
                        "description": "Building created info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.BuildingCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/building/all-company": {
            "get": {
                "description": "Get all company registered in a building",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Building"
                ],
                "summary": "Get all company in a building",
                "operationId": "getAllCompany",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "City id",
                        "name": "city_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Street id",
                        "name": "street_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "House id",
                        "name": "house",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.CompanyResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/company/": {
            "get": {
                "description": "Get company by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Get company by id",
                "operationId": "getCompany",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Company id",
                        "name": "company_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/common.CompanyResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Creating company with minimal required data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Company"
                ],
                "summary": "Create company",
                "operationId": "createCompany",
                "parameters": [
                    {
                        "description": "Company created info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/common.CompanyCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "id",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        },
        "/api/rubric/all-company": {
            "get": {
                "description": "Get all rubric company (without pagination)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rubric"
                ],
                "summary": "Get all company in a rubric",
                "operationId": "getAllRubricCompany",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Rubric id",
                        "name": "rubric_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/common.CompanyResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "404": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/handler.errorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "common.BuildingCreateRequest": {
            "type": "object",
            "required": [
                "house"
            ],
            "properties": {
                "city": {
                    "description": "название города (если есть), используем при создании записей в БД, если такого еще нет",
                    "type": "string"
                },
                "city_id": {
                    "type": "integer"
                },
                "house": {
                    "type": "integer"
                },
                "point": {
                    "type": "string"
                },
                "street": {
                    "description": "название улицы (если есть), используем при создании записей в БД, если такой улици еще нет",
                    "type": "string"
                },
                "street_id": {
                    "type": "integer"
                }
            }
        },
        "common.CompanyCreateRequest": {
            "type": "object",
            "properties": {
                "building": {
                    "type": "object",
                    "properties": {
                        "city_id": {
                            "description": "id города должен существовать в БД",
                            "type": "integer"
                        },
                        "house": {
                            "type": "integer"
                        },
                        "point": {
                            "description": "в формате \"(1.00234567, -90.00876211)\"",
                            "type": "string"
                        },
                        "street_id": {
                            "description": "id улицы с привязкой к городу должен существовать в БД",
                            "type": "integer"
                        }
                    }
                },
                "name": {
                    "type": "string"
                },
                "phones": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "number": {
                                "description": "номер телефона",
                                "type": "string"
                            }
                        }
                    }
                },
                "rubric": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "id": {
                                "description": "id рубрики должен существовать в БД",
                                "type": "integer"
                            }
                        }
                    }
                }
            }
        },
        "common.CompanyResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phones": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.Phone"
                    }
                }
            }
        },
        "common.Phone": {
            "type": "object",
            "properties": {
                "number": {
                    "type": "string"
                }
            }
        },
        "handler.errorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "localhost:80",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Catalog App API",
	Description: "API Server for company catalog",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
