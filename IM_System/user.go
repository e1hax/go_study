package main

/**
v2.0
User 类型
创建user对象的方法
监听user对应的channel消息的方法

*/

import "net"

// User 创建一个User类型
type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn
}

// ListenMessage 监听当前User  channel的方法，一旦有消息，就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

// NewUser 创建一个用户的API
func NewUser(conn net.Conn) *User {
	//获取用户的地址
	userAddr := conn.RemoteAddr().String()

	//创建一个用户
	user := &User{
		Name: userAddr,
		Addr: userAddr,
		C:    make(chan string),
		conn: conn,
	}

	//启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}
