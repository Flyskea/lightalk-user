package main

import (
	"flag"
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/flyskea/lightalk-user/rpc/internal/config"
	"github.com/flyskea/lightalk-user/rpc/internal/server"
	"github.com/flyskea/lightalk-user/rpc/internal/svc"
	"github.com/flyskea/lightalk-user/rpc/user"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	var dbInfo config.DBInfo
	if err := env.Parse(&dbInfo, env.Options{
		Prefix: "DB_",
	}); err != nil {
		panic(err)
	}

	ctx := svc.NewServiceContext(c, config.ConfigEnv{DBInfo: &dbInfo})

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
