package main

import (
	"flag"
	"fmt"

	"cloudclass_go/internal/config"
	"cloudclass_go/internal/handler"
	"cloudclass_go/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/digitalclass-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//socket长连接初始化

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	//fmt.Print(int(time.Now().Weekday()))
	server.Start()
}
