package application

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	flags "github.com/jessevdk/go-flags"

	"data-collector/internal/config"
	"data-collector/internal/generated/restapi"
	"data-collector/internal/generated/restapi/operations"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (app *App) Run() {
	cfg := config.InitConfig()
	app.runSwaggerServer(cfg)
}

func (app *App) runSwaggerServer(cfg config.Config) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewArticleCollectorAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Article collector"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	// TODO: Add:
	//server.host = cfg.Host
	server.Port = cfg.Port
	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
