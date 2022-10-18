package main

/**
v2.0
Server类型新增OnlineMap和Message属性
Handler 创建并添加用户
新增广播消息方法
新增监听广播消息channel方法

*/

import (
	"fmt"
	"net"
	"sync"
)

//Server 类型
type Server struct {
	Ip   string
	Port int
	//在线用户列表
	OnlineMap map[string]*User
	//读写锁
	mapLock sync.RWMutex

	//消息广播的channel
	Message chan string
}

//处理客户端上线
func (this *Server) Handler(conn net.Conn) {
	//fmt.Println("服务器链接成功...")

	user := NewUser(conn)

	//用户上线，将用户加入到在线用户列表 OnlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	//广播当前用户上线的消息
	this.broadCast(user, "上线了")
}

// NewServer 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		//初始化
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}

// Start 启动服务器的接口方法
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}

	//close listener
	defer listener.Close()

	//启动监听Message的goroutine
	go this.ListenMessager()

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

//广播消息的方法
func (this *Server) broadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

//监听广播消息的channel方法
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		//将msg发送给全部在线的User
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}
