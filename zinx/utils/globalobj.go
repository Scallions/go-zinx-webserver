package utils

import (
	"encoding/json"
	"io/ioutil"

	"../ziface"
)

type GolbalObj struct {
	TcpServer ziface.IServer
	Host string
	TcpPort int 
	Name string 
	Version string
	MaxPacketSize uint32
	MaxConn int

	// worker pool
	WorkerPoolSize uint32
	MaxWorkerTaskLen uint

	ConfFilePath string
}

var GlobalObject *GolbalObj


func (g *GolbalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GolbalObj{
		Name: "ZinxServerApp",
		Version: "V0.4",
		TcpPort: 7777,
		Host: "0.0.0.0",
		MaxConn: 12000,
		MaxPacketSize: 4096,
		WorkerPoolSize: 10,
		MaxWorkerTaskLen: 1024,
		ConfFilePath: "conf/zinx.json",
	}
	GlobalObject.Reload()
}