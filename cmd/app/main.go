package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/fx"

	"mango_crm/api/v1/router"
	"mango_crm/internal/orchard/application/usecase"
	infra "mango_crm/internal/orchard/infraestructure"
	"mango_crm/internal/orchard/infraestructure/observability"
	"mango_crm/pkg/config"
	dbconst "mango_crm/pkg/constants"
	"mango_crm/pkg/db"
	errorsPkg "mango_crm/pkg/errors"
	"mango_crm/pkg/logger"
)

var appModule = fx.Options(
	fx.Provide(logger.NewLogger, config.LoadConfig, db.NewMongoClient),
	fx.Invoke(run),
)

func run(log *logger.Logger, cfg *config.Config, client *mongo.Client) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	dbconst.Init()
	shutdownOTel, promHandler, err := observability.Init(observability.Config{
		ServiceName:  "orchard-service",
		OtlpEndpoint: cfg.OtlpEndpoint,
	})
	if err != nil {
		log.Error(err.Error())
		return
	}
	defer shutdownOTel(ctx)

	col := client.Database(dbconst.DatabaseName).Collection(dbconst.CollectionName)
	repo := infra.NewMongoRepositoryImpl(col, log)
	uc := usecase.NewOrchardUseCase(repo, log)

	e := echo.New()
	e.HTTPErrorHandler = errorsPkg.HandleError
	e.Use(observability.MetricsMiddleware())
	e.GET("/metrics", echo.WrapHandler(promHandler))
	router.RegisterRoutes(router.RouterConfig{
		Echo:           e,
		OrchardUseCase: uc,
		Log:            log,
	})

	go func() {
		addr := fmt.Sprintf(":%d", cfg.HttpPort)
		if err := e.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctxShut, cancelShut := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShut()
	_ = e.Shutdown(ctxShut)
}

func main() {
	fx.New(appModule).Run()

}
