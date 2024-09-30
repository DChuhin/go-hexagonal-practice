package boot

import (
	"context"
	mongodbiotapplication "go-hexagonal-practice/internal/adapters/output/mongodb/iotapplication"
	inputports "go-hexagonal-practice/internal/application/ports/input"
	outputports "go-hexagonal-practice/internal/application/ports/output"
	"go-hexagonal-practice/internal/application/service/iotapplication"
	"go-hexagonal-practice/internal/integration/mongodb/repository"
)

func RunApplication(ctx context.Context) {
	config := readConfig[AppConfig]()
	outputAdapters := runOutputAdapters(ctx, config)
	services := runServices(outputAdapters, config)
	startInputAdapters(ctx, config, services)
}

type applicationServices struct {
	iotApplicationService inputports.IoTApplicationService
}

type outputAdapters struct {
	iotApplicationStorage outputports.IoTApplicationStorage
}

func runOutputAdapters(ctx context.Context, config *AppConfig) *outputAdapters {

	mongoClient := createMongoClient(config.MongoConfig.Uri)
	mongoDatabase := mongoClient.Database(config.MongoConfig.Database)

	iotApplicationRepository := repository.New[mongodbiotapplication.Document](mongoDatabase.Collection("applications"))
	iotApplicationStorage := mongodbiotapplication.NewStorage(ctx, iotApplicationRepository)

	return &outputAdapters{
		iotApplicationStorage: iotApplicationStorage,
	}
}

func runServices(outputAdapters *outputAdapters, config *AppConfig) *applicationServices {
	iotApplicationService := iotapplication.NewService(outputAdapters.iotApplicationStorage)

	return &applicationServices{
		iotApplicationService: iotApplicationService,
	}
}

func startInputAdapters(ctx context.Context, config *AppConfig, services *applicationServices) {
	runRestApiServer(ctx, config, services)
}
