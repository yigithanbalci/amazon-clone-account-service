package main

import (
	"fmt"

	"github.com/yigithanbalci/amazon-clone-account-service/dbclient"
	"github.com/yigithanbalci/amazon-clone-account-service/service"
)

var appname = "accountservice"

func main() {
	fmt.Printf("starting %v\n", appname)
	initializeBoltClient()
	service.StartWebServer("6767")
}

func initializeBoltClient() {
	service.Dbclient = &dbclient.BoltClient{}
	service.Dbclient.OpenBoltDb()
	service.Dbclient.Seed()
}
