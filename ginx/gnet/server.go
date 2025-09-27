package gnet

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
func (s *Server) Start() {}

// Stop 停止服务器/**
func (s *Server) Stop() {}

// ListenAndServe 监听server端口/**
func (s *Server) ListenAndServe() {}
