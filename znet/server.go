package znet

import "myzinx/ziface"

// IServer的接口实现
type Server struct {
	//服务器名称
	Name string
	//服务器绑定IP版本
	IPVersion string
	//服务器监听ip
	IP string
	//监听port
	Port int
}

// 开启服务
func (s *Server) Start() {}

// 停止服务
func (s *Server) Stop() {}

// 运行服务
func (s *Server) Serve() {}

// 初始化server
func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
