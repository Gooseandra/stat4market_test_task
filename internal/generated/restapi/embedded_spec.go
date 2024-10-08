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
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API для работы с событиями и пользователями",
    "title": "Events API",
    "version": "1.0.0"
  },
  "host": "api.example.com",
  "basePath": "/v1",
  "paths": {
    "/event": {
      "get": {
        "description": "Возвращает список событий по указанному типу и дате",
        "produces": [
          "application/json"
        ],
        "summary": "Получить события по типу и дате",
        "operationId": "GetEvent",
        "parameters": [
          {
            "type": "string",
            "description": "Тип события",
            "name": "type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Дата в формате YYYY-MM-DD",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/event/new": {
      "post": {
        "description": "Добавляет новое событие в базу данных",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Создать новое событие",
        "operationId": "NewEvent",
        "parameters": [
          {
            "description": "Данные нового события",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEventRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/NewEventResponse"
            }
          },
          "400": {
            "description": "Плохой запрос",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/events/day/{day}": {
      "get": {
        "description": "Возвращает список событий, произошедших в указанный день",
        "produces": [
          "application/json"
        ],
        "summary": "Получить события за день",
        "operationId": "GetEventByDay",
        "parameters": [
          {
            "type": "string",
            "description": "Дата в формате YYYY-MM-DD",
            "name": "day",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventByDayResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/events/value/{value}": {
      "get": {
        "description": "Возвращает список типов событий с указанным значением",
        "produces": [
          "application/json"
        ],
        "summary": "Получить типы событий по значению",
        "operationId": "GetEventByValue",
        "parameters": [
          {
            "type": "integer",
            "description": "Значение для фильтрации типов событий",
            "name": "value",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventByValueResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/user/{value}": {
      "get": {
        "description": "Возвращает список пользователей по значению событий",
        "produces": [
          "application/json"
        ],
        "summary": "Получить пользователей по значению",
        "operationId": "GetUserByValue",
        "parameters": [
          {
            "type": "integer",
            "description": "Значение для фильтрации пользователей",
            "name": "value",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetUserByValueResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "EventDetail": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "integer"
        },
        "event_time": {
          "type": "string",
          "format": "date-time"
        },
        "event_type": {
          "type": "string"
        },
        "payload": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "EventType": {
      "type": "object",
      "properties": {
        "Types": {
          "type": "string"
        },
        "Value": {
          "type": "integer"
        }
      }
    },
    "GetEventByDayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDetail"
          }
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "GetEventByValueResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        },
        "types": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventType"
          }
        }
      }
    },
    "GetEventResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDetail"
          }
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "GetUserByValueResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        },
        "users_ids": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/UserDetail"
          }
        }
      }
    },
    "NewEventRequest": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "integer"
        },
        "event_time": {
          "type": "string",
          "format": "date-time"
        },
        "event_type": {
          "type": "string"
        },
        "payload": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "NewEventResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "UserDetail": {
      "type": "object",
      "properties": {
        "User": {
          "type": "integer"
        },
        "Value": {
          "type": "integer"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "https",
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API для работы с событиями и пользователями",
    "title": "Events API",
    "version": "1.0.0"
  },
  "host": "api.example.com",
  "basePath": "/v1",
  "paths": {
    "/event": {
      "get": {
        "description": "Возвращает список событий по указанному типу и дате",
        "produces": [
          "application/json"
        ],
        "summary": "Получить события по типу и дате",
        "operationId": "GetEvent",
        "parameters": [
          {
            "type": "string",
            "description": "Тип события",
            "name": "type",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "description": "Дата в формате YYYY-MM-DD",
            "name": "date",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/event/new": {
      "post": {
        "description": "Добавляет новое событие в базу данных",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "summary": "Создать новое событие",
        "operationId": "NewEvent",
        "parameters": [
          {
            "description": "Данные нового события",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/NewEventRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/NewEventResponse"
            }
          },
          "400": {
            "description": "Плохой запрос",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/events/day/{day}": {
      "get": {
        "description": "Возвращает список событий, произошедших в указанный день",
        "produces": [
          "application/json"
        ],
        "summary": "Получить события за день",
        "operationId": "GetEventByDay",
        "parameters": [
          {
            "type": "string",
            "description": "Дата в формате YYYY-MM-DD",
            "name": "day",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventByDayResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/events/value/{value}": {
      "get": {
        "description": "Возвращает список типов событий с указанным значением",
        "produces": [
          "application/json"
        ],
        "summary": "Получить типы событий по значению",
        "operationId": "GetEventByValue",
        "parameters": [
          {
            "type": "integer",
            "description": "Значение для фильтрации типов событий",
            "name": "value",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetEventByValueResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    },
    "/user/{value}": {
      "get": {
        "description": "Возвращает список пользователей по значению событий",
        "produces": [
          "application/json"
        ],
        "summary": "Получить пользователей по значению",
        "operationId": "GetUserByValue",
        "parameters": [
          {
            "type": "integer",
            "description": "Значение для фильтрации пользователей",
            "name": "value",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Успешный ответ",
            "schema": {
              "$ref": "#/definitions/GetUserByValueResponse"
            }
          },
          "500": {
            "description": "Ошибка сервера",
            "schema": {
              "$ref": "#/definitions/ErrorResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "ErrorResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "EventDetail": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "integer"
        },
        "event_time": {
          "type": "string",
          "format": "date-time"
        },
        "event_type": {
          "type": "string"
        },
        "payload": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "EventType": {
      "type": "object",
      "properties": {
        "Types": {
          "type": "string"
        },
        "Value": {
          "type": "integer"
        }
      }
    },
    "GetEventByDayResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDetail"
          }
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "GetEventByValueResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        },
        "types": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventType"
          }
        }
      }
    },
    "GetEventResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "events": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/EventDetail"
          }
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "GetUserByValueResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        },
        "users_ids": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/UserDetail"
          }
        }
      }
    },
    "NewEventRequest": {
      "type": "object",
      "properties": {
        "event_id": {
          "type": "integer"
        },
        "event_time": {
          "type": "string",
          "format": "date-time"
        },
        "event_type": {
          "type": "string"
        },
        "payload": {
          "type": "string"
        },
        "user_id": {
          "type": "integer"
        }
      }
    },
    "NewEventResponse": {
      "type": "object",
      "properties": {
        "error": {
          "type": "string"
        },
        "success": {
          "type": "boolean"
        }
      }
    },
    "UserDetail": {
      "type": "object",
      "properties": {
        "User": {
          "type": "integer"
        },
        "Value": {
          "type": "integer"
        }
      }
    }
  }
}`))
}
