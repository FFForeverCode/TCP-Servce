package gnet

import (
	"fmt"
	"github.com/FFForeverCode/ginx/giface"
	"net"
)

//server服务器实现

type Server struct {
	//server名称
	Name string

	//server的IP版本
	IPVersion string

	//server绑定的ip地址
	IP string

	//server绑定的端口
	Port int
}

// Start 开启server/**
func (s *Server) Start() {
	fmt.Println("server is starting...")
	//1. 创建addr、解析ip、端口
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("addr is err", err)
	}
	//2. 监听addr:ip,port
	tcp, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen err", err)
	}
	//3.监听端口后，循环尝试创建连接
	for {
		conn, err := tcp.AcceptTCP()
		if err != nil {
			fmt.Println("accept err", err)
			continue
		}
		//已建立连接、执行任务
		for {
			go func() {
				buf := make([]byte, 512)
				read, err2 := conn.Read(buf)
				if err2 != nil {
					fmt.Println("read err", err2)
				}
				if _, writerErr := conn.Write(buf[:read]); writerErr != nil {
					fmt.Println("write err", writerErr)
				}
			}()
		}
	}

}

// Stop 停止服务器/**
func (s *Server) Stop() {
	fmt.Println("server is stoping...")
	//todo 服务器关闭、关闭资源
}

// ListenAndServe 监听server端口/**
func (s *Server) ListenAndServe() {

}

// Serve 开启服务
func (s *Server) Serve() {
	//开启服务
	s.Start()
	//TODO 服务器已开启，执行其他业务

	select {}

}

// NewServer 创建一个server句柄
func NewServer(name string) giface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
}
