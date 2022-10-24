package znet

import (
	"errors"
	"fmt"
	"myzinx/ziface"
	"net"
	"time"
)

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

// 后续由用户指定
func CallbackClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[conn handle callbacktoclient]")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("callbackclient error")
	}
	return nil
}

// 开启服务
func (s *Server) Start() {
	fmt.Printf("[start] server listenner at ip %s, port: %d is starting", s.IP, s.Port)

	//避免单个阻塞，支持多客户端连接
	go func() {
		//1、获取一个tcp的addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}
		//2、监听服务器的地址
		listenner, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listenTcp error: ", err)
			return
		}
		fmt.Println("start zinx server success,", s.Name, " is listening")
		//3、阻塞的等待客户端连接，处理客户端连接业务

		var cid uint32 = 0

		for {
			//如果有客户端连接过来，阻塞会返回
			conn, err := listenner.AcceptTCP()
			if err != nil {
				fmt.Println("accept error", err)
				continue
			}

			//将处理新连接的业务方法和conn进行绑定，得到我们的连接模块
			dealConn := NewConnection(conn, cid, CallbackClient)

			cid++

			//启动当前的连接业务处理
			go dealConn.Start()

		}
	}()

}

// 停止服务
func (s *Server) Stop() {
	fmt.Println("[stop] ", s.Name)

}

// 运行服务
func (s *Server) Serve() {
	//启动服务功能，只进行监听和处理功能
	s.Start()

	//todo:可做一些启动服务器之后的额外业务

	//处于阻塞
	for {
		time.Sleep(time.Second * 10)
	}
}

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
