package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("服务器链接成功...")
}

//  创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

//启动服务器的接口方法
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}

	//close listener
	defer listener.Close()

	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accpet err:", err)
			continue //继续下一次监听
		}
		//do handler
		go this.Handler(conn)
	}
}
