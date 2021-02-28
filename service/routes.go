package service

import "net/http"

type route struct {
	name        string
	method      string
	path        string
	handlerFunc http.HandlerFunc
}

type routes []route

var routeList = routes{
	route{
		name:   "getaccount",
		method: "GET",
		path:   "/accounts/{accountid}",
		handlerFunc: func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("content-type", "application/json; charset=utf-8")
			w.Write([]byte("{\"result\":\"ok\"}"))
		},
	},
}
