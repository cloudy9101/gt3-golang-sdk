package main

import (
	"github.com/cloudy9101/gt3-golang-sdk/conf"
	"github.com/cloudy9101/gt3-golang-sdk/route"
	"github.com/cloudy9101/gt3-golang-sdk/service"
)

func main() {
	// conf init
	if err := conf.Init(); err != nil {
		panic(err)
	}
	// service init
	serv := service.New(conf.Conf)
	// http init
	route.Init(conf.Conf, serv)
}
