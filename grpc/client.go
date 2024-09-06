package grpc

import (
	"api_gateway/configs"
	pbu "api_gateway/genproto/auth_service"
	pb "api_gateway/genproto/task_service"
	"api_gateway/pkg/logger"

	"google.golang.org/grpc"
)

type IServiceManager interface {
	UserService() pbu.UserServiceClient
	TaskService() pb.TaskServiceClient
	DetailService() pb.DetailServiceClient
}

type GprcClient struct {
	userService   pbu.UserServiceClient
	taskService   pb.TaskServiceClient
	detailService pb.DetailServiceClient
}

func (g *GprcClient) UserService() pbu.UserServiceClient {
	return g.userService
}

func (g *GprcClient) TaskService() pb.TaskServiceClient {
	return g.taskService
}

func (g *GprcClient) DetailService() pb.DetailServiceClient {
	return g.detailService
}

func NewGrpcClient(cnf configs.Config, log logger.ILogger) (IServiceManager, error) {
	connUserService, err :=
		grpc.Dial(
			cnf.AuthServiceGrpcHost+cnf.AuthServiceGrpcPort,
			grpc.WithInsecure())

	if err != nil {
		log.Error("this error is  connUserService with dialing ", logger.Error(err))
		return nil, err
	}

	connTaskService, err :=
		grpc.Dial(
			cnf.TaskServiceGrpcHost+cnf.TaskServiceGrpcPort,
			grpc.WithInsecure())
	if err != nil {
		log.Error("this error is  connSellerService with dialing ", logger.Error(err))

	}

	return &GprcClient{
		userService:   pbu.NewUserServiceClient(connUserService),
		taskService:   pb.NewTaskServiceClient(connTaskService),
		detailService: pb.NewDetailServiceClient(connTaskService),
	}, nil
}
