// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// CreateAppHandlerFunc turns a function with the right signature into a create app handler
type CreateAppHandlerFunc func(CreateAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateAppHandlerFunc) Handle(params CreateAppParams) middleware.Responder {
	return fn(params)
}

// CreateAppHandler interface for that can handle valid create app params
type CreateAppHandler interface {
	Handle(CreateAppParams) middleware.Responder
}

// NewCreateApp creates a new http.Handler for the create app operation
func NewCreateApp(ctx *middleware.Context, handler CreateAppHandler) *CreateApp {
	return &CreateApp{Context: ctx, Handler: handler}
}

/*CreateApp swagger:route POST /api/apps createApp

CreateApp create app API

*/
type CreateApp struct {
	Context *middleware.Context
	Handler CreateAppHandler
}

func (o *CreateApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
