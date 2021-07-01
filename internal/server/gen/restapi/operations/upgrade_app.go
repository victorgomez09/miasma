// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpgradeAppHandlerFunc turns a function with the right signature into a upgrade app handler
type UpgradeAppHandlerFunc func(UpgradeAppParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpgradeAppHandlerFunc) Handle(params UpgradeAppParams) middleware.Responder {
	return fn(params)
}

// UpgradeAppHandler interface for that can handle valid upgrade app params
type UpgradeAppHandler interface {
	Handle(UpgradeAppParams) middleware.Responder
}

// NewUpgradeApp creates a new http.Handler for the upgrade app operation
func NewUpgradeApp(ctx *middleware.Context, handler UpgradeAppHandler) *UpgradeApp {
	return &UpgradeApp{Context: ctx, Handler: handler}
}

/*UpgradeApp swagger:route PUT /api/apps/{appName}/upgrade upgradeApp

pull the app's image and restart it

*/
type UpgradeApp struct {
	Context *middleware.Context
	Handler UpgradeAppHandler
}

func (o *UpgradeApp) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewUpgradeAppParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
