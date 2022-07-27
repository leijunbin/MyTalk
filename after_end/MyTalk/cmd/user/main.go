package main

import (
	"MyTalk/cmd/user/dal"
	user "MyTalk/kitex_gen/user/userservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Init()
}

func main() {
	Init()

	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})

	if err != nil {
		log.Panic("etcd initialization error")
	}

	addr, err := net.ResolveTCPAddr("tcp", "localhost:8888")

	if err != nil {
		log.Panic("tcp address initialization error")
	}

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "user"}),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
	)

	err = svr.Run()

	if err != nil {
		log.Panic("user service initialization err")
	}
}
