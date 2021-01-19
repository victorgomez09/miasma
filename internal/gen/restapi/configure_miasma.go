// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/dre1080/recovr"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/aklinker1/miasma/internal/gen/restapi/operations"
	customMiddleware "github.com/aklinker1/miasma/internal/middleware"
)

//go:generate swagger generate server --target ../../gen --name Miasma --spec ../../../api/swagger.yml --principal interface{} --exclude-main

func configureFlags(api *operations.MiasmaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MiasmaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.CreateAppHandler == nil {
		api.CreateAppHandler = operations.CreateAppHandlerFunc(func(params operations.CreateAppParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.CreateApp has not yet been implemented")
		})
	}
	if api.DeleteAppHandler == nil {
		api.DeleteAppHandler = operations.DeleteAppHandlerFunc(func(params operations.DeleteAppParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.DeleteApp has not yet been implemented")
		})
	}
	if api.GetAppHandler == nil {
		api.GetAppHandler = operations.GetAppHandlerFunc(func(params operations.GetAppParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetApp has not yet been implemented")
		})
	}
	if api.GetAppsHandler == nil {
		api.GetAppsHandler = operations.GetAppsHandlerFunc(func(params operations.GetAppsParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetApps has not yet been implemented")
		})
	}
	if api.GetHealthCheckHandler == nil {
		api.GetHealthCheckHandler = operations.GetHealthCheckHandlerFunc(func(params operations.GetHealthCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation operations.GetHealthCheck has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	recovery := recovr.New()
	ui := customMiddleware.UI()
	return recovery(ui(handler))
}
