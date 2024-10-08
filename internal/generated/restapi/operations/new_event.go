// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// NewEventHandlerFunc turns a function with the right signature into a new event handler
type NewEventHandlerFunc func(NewEventParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NewEventHandlerFunc) Handle(params NewEventParams) middleware.Responder {
	return fn(params)
}

// NewEventHandler interface for that can handle valid new event params
type NewEventHandler interface {
	Handle(NewEventParams) middleware.Responder
}

// NewNewEvent creates a new http.Handler for the new event operation
func NewNewEvent(ctx *middleware.Context, handler NewEventHandler) *NewEvent {
	return &NewEvent{Context: ctx, Handler: handler}
}

/*
	NewEvent swagger:route POST /event/new newEvent

# Создать новое событие

Добавляет новое событие в базу данных
*/
type NewEvent struct {
	Context *middleware.Context
	Handler NewEventHandler
}

func (o *NewEvent) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNewEventParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
