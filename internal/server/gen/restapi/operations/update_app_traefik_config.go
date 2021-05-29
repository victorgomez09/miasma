// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateAppTraefikConfigHandlerFunc turns a function with the right signature into a update app traefik config handler
type UpdateAppTraefikConfigHandlerFunc func(UpdateAppTraefikConfigParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateAppTraefikConfigHandlerFunc) Handle(params UpdateAppTraefikConfigParams) middleware.Responder {
	return fn(params)
}

// UpdateAppTraefikConfigHandler interface for that can handle valid update app traefik config params
type UpdateAppTraefikConfigHandler interface {
	Handle(UpdateAppTraefikConfigParams) middleware.Responder
}

// NewUpdateAppTraefikConfig creates a new http.Handler for the update app traefik config operation
func NewUpdateAppTraefikConfig(ctx *middleware.Context, handler UpdateAppTraefikConfigHandler) *UpdateAppTraefikConfig {
	return &UpdateAppTraefikConfig{Context: ctx, Handler: handler}
}

/*UpdateAppTraefikConfig swagger:route PUT /api/plugins/traefik/{appId} updateAppTraefikConfig

Update an app's routing config

*/
type UpdateAppTraefikConfig struct {
	Context *middleware.Context
	Handler UpdateAppTraefikConfigHandler
}

func (o *UpdateAppTraefikConfig) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpdateAppTraefikConfigParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
