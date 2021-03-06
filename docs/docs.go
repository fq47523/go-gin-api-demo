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
        "termsOfService": "http://github.com",
        "contact": {
            "name": "API Support",
            "url": "http://www.cnblogs.com",
            "email": "×××@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/release/template/add": {
            "post": {
                "description": "新增VueData",
                "tags": [
                    "release template"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 31a165baebe6dec616b1f8f3207b4273",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "JSON数据",
                        "name": "vue",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.VueData"
                        }
                    }
                ],
                "responses": {
                    "20000": {
                        "schema": {
                            "$ref": "#/definitions/vuedata.response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Platforms": {
            "type": "object",
            "properties": {
                "name": {
                    "description": "gorm.Model",
                    "type": "string"
                },
                "vueDataID": {
                    "type": "integer"
                }
            }
        },
        "models.VueData": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "comment_disabled": {
                    "type": "boolean"
                },
                "content": {
                    "type": "string"
                },
                "content_short": {
                    "type": "string"
                },
                "display_time": {
                    "type": "string"
                },
                "forecast": {
                    "type": "number"
                },
                "image_uri": {
                    "type": "string"
                },
                "importance": {
                    "type": "integer"
                },
                "pageviews": {
                    "type": "integer"
                },
                "platforms": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Platforms"
                    }
                },
                "reviewer": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "timestamp": {
                    "description": "gorm.Model",
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "vuedata.response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
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
	Version:     "1.0",
	Host:        "127.0.0.1:8081",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Esign API",
	Description: "Golang api of demo",
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
