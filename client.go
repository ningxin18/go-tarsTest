package main

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/go-gppay/go-tarsTest/TestApp"
)

//tars.Communicator should only init once and be global
var comm *tars.Communicator

func main() {
	comm = tars.NewCommunicator()
	obj := "TestApp.TestServer.HelloObj@tcp -h 127.0.0.1 -p 10015 -t 60000"
	app := new(TestApp.Hello)
	comm.StringToProxy(obj, app)
	var req string="Hello World"
	var out string
	ret, err := app.TestHello(req, &out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret, out)
}