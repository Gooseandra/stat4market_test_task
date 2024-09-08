// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetUserByValueHandlerFunc turns a function with the right signature into a get user by value handler
type GetUserByValueHandlerFunc func(GetUserByValueParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetUserByValueHandlerFunc) Handle(params GetUserByValueParams) middleware.Responder {
	return fn(params)
}

// GetUserByValueHandler interface for that can handle valid get user by value params
type GetUserByValueHandler interface {
	Handle(GetUserByValueParams) middleware.Responder
}

// NewGetUserByValue creates a new http.Handler for the get user by value operation
func NewGetUserByValue(ctx *middleware.Context, handler GetUserByValueHandler) *GetUserByValue {
	return &GetUserByValue{Context: ctx, Handler: handler}
}

/*
	GetUserByValue swagger:route GET /user/{value} getUserByValue

# Получить пользователей по значению

Возвращает список пользователей по значению событий
*/
type GetUserByValue struct {
	Context *middleware.Context
	Handler GetUserByValueHandler
}

func (o *GetUserByValue) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetUserByValueParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
