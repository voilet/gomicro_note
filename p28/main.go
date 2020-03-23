package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	_ "gomicro_note/p28/appInit"
	"gomicro_note/p28/controllers"
	"gomicro_note/p28/models"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	app := micro.NewService(
		micro.Name("api.hzde.com.user"),
		micro.Address(":8000"),
		micro.Registry(etcdReg),
	)

	models.RegisterUserServiceHandler(app.Server(), new(controllers.UserService))
	app.Init()

	app.Run()
}
