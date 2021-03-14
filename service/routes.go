package service

import (
	"net/http"
)

type route struct {
	name        string
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

type routes []route

var routeList = routes{
	route{
		name:        "getaccount",
		method:      "GET",
		path:        "/accounts/{accountID}",
		handlerFunc: GetAccount,
	},
}
