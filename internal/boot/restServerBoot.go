package boot

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	restv1 "go-hexagonal-practice/internal/adapters/input/rest/v1"
	"go-hexagonal-practice/internal/docs"
	"log"
	"net/http"
)

// @title Swagger Network API
// @version 1.0-SNAPSHOT
func runRestApiServer(ctx context.Context, config *AppConfig, services *applicationServices) {
	baseUrl := config.RestConfig.BaseUrl
	basePort := config.RestConfig.BasePort
	authHeader := config.RestConfig.AuthHeader

	router := gin.Default()
	baseRouter := router.Group(baseUrl)
	v1Router := baseRouter.Group("/v1")

	swaggerRoutesConfig(baseRouter, baseUrl)

	iotApplicationRouter := v1Router.Group("/applications")
	iotApplicationHandler := restv1.NewIotApplicationHandler(authHeader, services.iotApplicationService)
	iotApplicationHandler.InitRoutes(iotApplicationRouter)

	srv := &http.Server{
		Addr:    ":" + basePort,
		Handler: router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	<-ctx.Done()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}

func swaggerRoutesConfig(router *gin.RouterGroup, basePath string) {
	docs.SwaggerInfo.BasePath = basePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
