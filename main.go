package main

import (
	"fmt"

	"github.com/yigithanbalci/amazon-clone-account-service/service"
)

var appname = "accountservice"

func main() {
	fmt.Printf("starting %v\n", appname)
	service.StartWebServer("6767")
}
