package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"mxshop/user/global"
	"mxshop/user/handler"
	"mxshop/user/initialize"
	"mxshop/user/proto"
	"net"
)

func main() {
	IP := flag.String("ip", "127.0.0.1", "ip地址")
	Port := flag.Int("port", 8080, "端口号")
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()
	flag.Parse()
	zap.S().Info("ip:", *IP)
	zap.S().Info("port", *Port)
	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserService{})
	//注册服务健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())
	cfg := api.DefaultConfig()
	//填虚拟机的地址以及端口8500
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成检查的对象
	/*check := &api.AgentServiceCheck{
		GRPC:                           "192.168.0.104:8080",
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}*/

	//生成一个注册对象
	registeration := new(api.AgentServiceRegistration)
	registeration.Port = *Port
	registeration.Name = global.ServerConfig.Name
	registeration.Address = "192.168.0.104"
	registeration.Tags = []string{"imooc", "bobby"}
	//registeration.Check = check

	err = client.Agent().ServiceRegister(registeration)
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic(err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic(err.Error())
	}
}
