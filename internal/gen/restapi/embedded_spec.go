// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running applications",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createApp",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "operationId": "getApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "operationId": "deleteApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "required": [
                "version"
              ],
              "properties": {
                "version": {
                  "description": "Miasma's current version",
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image"
      ],
      "properties": {
        "hidden": {
          "description": "Wether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Miasma",
    "version": "0.0.0"
  },
  "paths": {
    "/api/apps": {
      "get": {
        "summary": "List all the running applications",
        "operationId": "getApps",
        "parameters": [
          {
            "type": "boolean",
            "description": "Whether or not to show hidden apps",
            "name": "hidden",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/App"
              }
            }
          },
          "default": {
            "description": "Unknown Error",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "post": {
        "operationId": "createApp",
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/App"
            }
          }
        }
      }
    },
    "/api/apps/{appName}": {
      "get": {
        "operationId": "getApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      },
      "delete": {
        "operationId": "deleteApp",
        "parameters": [
          {
            "type": "string",
            "description": "App name from the ` + "`" + `-a|--app` + "`" + ` flag",
            "name": "appName",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/App"
            }
          },
          "404": {
            "description": "Not Found",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/api/health": {
      "get": {
        "summary": "Standard health check endpoint that checks all the service's statuses",
        "operationId": "getHealthCheck",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "object",
              "required": [
                "version"
              ],
              "properties": {
                "version": {
                  "description": "Miasma's current version",
                  "type": "string"
                }
              }
            }
          }
        }
      }
    }
  },
  "definitions": {
    "App": {
      "type": "object",
      "required": [
        "name",
        "image"
      ],
      "properties": {
        "hidden": {
          "description": "Wether or not the app is hidden during regular requests",
          "type": "boolean"
        },
        "image": {
          "description": "The image the app is based off of",
          "type": "string"
        },
        "name": {
          "description": "The apps name, used in the CLI with the ` + "`" + `-a|--app` + "`" + ` flag",
          "type": "string"
        }
      }
    }
  }
}`))
}
