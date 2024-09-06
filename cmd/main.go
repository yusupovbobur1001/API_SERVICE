package main

import (
	"api_gateway/api"
	_ "api_gateway/api/docs"
	"api_gateway/api/handler"
	"api_gateway/configs"
	"api_gateway/grpc"
	"api_gateway/pkg/logger"
	"fmt"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := configs.Load()

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case configs.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case configs.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger("api_gateway", loggerLevel)
	defer logger.Cleanup(log)

	fmt.Println(cfg.AuthServiceGrpcHost, "+++++++++")

	service, err := grpc.NewGrpcClient(cfg, log)
	if err != nil {
		log.Error(err.Error())
		return
	}

	handler1 := handler.NewHandler(cfg, log, service)

	enforcer, err := casbin.NewEnforcer("casbin/model.conf", "casbin/policy.csv")
	if err != nil {
		log.Error(err.Error())
		return
	}

	r := api.New(handler1, enforcer)

	r.Run()
}
