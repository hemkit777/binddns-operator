package main

import (
	"github.com/hemkit777/binddns-operator/pkg/controller/router"
	"github.com/hemkit777/binddns-operator/pkg/kube"
	zlog "github.com/hemkit777/binddns-operator/pkg/utils/zaplog"
)

func main() {
	zlog.DefaultLog("./bind.log")
	// Init kubernetes client
	err := kube.InitOutOfKubernetes()
	if err != nil {
		zlog.Error(err)
		return
	}

	server := router.NewHttpServer(":8088")
	server.Start()
}
