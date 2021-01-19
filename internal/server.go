package internal

import (
	"flag"
	"log"

	"github.com/go-openapi/loads"

	"github.com/aklinker1/miasma/internal/controllers"
	"github.com/aklinker1/miasma/internal/gen/restapi"
	"github.com/aklinker1/miasma/internal/gen/restapi/operations"
)

func Start() {
	// Load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create new service API
	api := operations.NewMiasmaAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Parse flags
	var portFlag = flag.Int("port", 3000, "Port to run this service on")
	flag.Parse()
	server.Port = *portFlag

	// Use handlers
	useControllers(api)

	// Serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func useControllers(api *operations.MiasmaAPI) {
	controllers.UseHealthController(api)
	controllers.UseAppsController(api)
}
