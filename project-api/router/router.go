package router

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"irms-api.com/project-api/api/login"
	"irms-api.com/project-api/config"
	"irms-api.com/project-api/pkg/repo/service/login_service_v1"
	"log"
	"net"
)

type Router interface {
	Register(r *gin.Engine)
}
type RegisterRouter struct {
}

func New() RegisterRouter {
	return RegisterRouter{}
}

func (RegisterRouter) Route(router Router, r *gin.Engine) {
	router.Register(r)
}

func InitRouter(r *gin.Engine) {
	router := New()
	router.Route(&login.RouterLogin{}, r)
}

type gRPCConfig struct {
	Addr         string
	RegisterFunc func(server *grpc.Server)
}

func RegisterGrpc() *grpc.Server {
	c := gRPCConfig{
		Addr: config.AppConf.GC.Addr,
		RegisterFunc: func(g *grpc.Server) {
			login_service_v1.RegisterLoginServiceServer(g, login_service_v1.New())
		},
	}
	s := grpc.NewServer()
	c.RegisterFunc(s)
	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Println("cannot listen")
	}
	go func() {
		err = s.Serve(lis)
		if err != nil {
			log.Println("server started error", err)
			return
		}
	}()
	return s
}
