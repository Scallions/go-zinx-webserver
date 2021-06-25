package main
import (
    "fmt"
    "./ziface"
    "./znet"
)
//ping test 自定义路由
type PingRouter struct {
    znet.BaseRouter
}
//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
    fmt.Println("Call PingRouter Handle")
    //先读取客户端的数据，再回写ping...ping...ping
    fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
    err := request.GetConnection().SendMsg(0, []byte("ping...ping...ping"))
    if err != nil {
        fmt.Println(err)
    }
}
//HelloZinxRouter Handle
type HelloZinxRouter struct {
    znet.BaseRouter
}
func (this *HelloZinxRouter) Handle(request ziface.IRequest) {
    fmt.Println("Call HelloZinxRouter Handle")
    fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))
    err := request.GetConnection().SendMsg(1, []byte("Hello Zinx Router V0.10"))
    if err != nil {
        fmt.Println(err)
    }
}

func DoConnectionBegin(conn ziface.IConnection) {
    fmt.Println("DoConnectionBegin is Called...")

    // set prop
    conn.SetProperty("Name", "Aceld")
    conn.SetProperty("Home", "https://hexo.scallions.cn")

    err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
    if err != nil {
        fmt.Println(err)
    }
}

func DoConnectionLost(conn ziface.IConnection) {
    fmt.Println("DoConneciotnLost is Called ... ")
    // query prop
    if name, err := conn.GetProperty("Name"); err == nil {
        fmt.Println("Conn Property Name = ", name)
    }
    if home, err := conn.GetProperty("Home"); err == nil {
        fmt.Println("Conn Property Home = ", home)
    }
}

func main() {
    //创建一个server句柄
    s := znet.NewServer()
    // hook
    s.SetOnConnStart(DoConnectionBegin)
    s.SetOnConnStop(DoConnectionLost)
    //配置路由
    s.AddRouter(0, &PingRouter{})
    s.AddRouter(1, &HelloZinxRouter{})
    //开启服务
    s.Serve()
}