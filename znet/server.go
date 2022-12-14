package znet

import (
	"errors"
	"fmt"
	"myzinx/utils"
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

	//当前Server由用户绑定的回调router，也就是server注册的链接对应的处理业务
	msgHandler ziface.IMsgHandle

	//集成连接管理
	ConnMgr ziface.IConManager

	// =======================
	//新增两个hook函数原型

	//该Server的连接创建时Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该Server的连接断开时的Hook函数
	OnConnStop func(conn ziface.IConnection)
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
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d,  MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	//避免单个阻塞，支持多客户端连接
	go func() {

		s.msgHandler.StartWorkerPool()

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

			//超过最大连接数则直接关闭连接
			if s.ConnMgr.Len() >= utils.GlobalObject.MaxConn {
				conn.Close()
				continue
			}

			//将处理新连接的业务方法和conn进行绑定，得到我们的连接模块
			fmt.Printf("----%+v\n", conn)
			dealConn := NewConnection(s, conn, cid, s.msgHandler)

			cid++

			//启动当前的连接业务处理
			go dealConn.Start()

		}
	}()

}

// 停止服务
func (s *Server) Stop() {
	fmt.Println("[stop] ", s.Name)
	s.ConnMgr.ClearConn()
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

func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.msgHandler.AddRouter(msgId, router)
	fmt.Println("add router success")
}

func (s *Server) GetConnMgr() ziface.IConManager {
	return s.ConnMgr
}

// 设置该Server的连接创建时Hook函数
func (s *Server) SetOnConnStart(hookFunc func(ziface.IConnection)) {
	s.OnConnStart = hookFunc
}

// 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(hookFunc func(ziface.IConnection)) {
	s.OnConnStop = hookFunc
}

// 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		fmt.Println("---> CallOnConnStart....")
		s.OnConnStart(conn)
	}
}

// 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		fmt.Println("---> CallOnConnStop....")
		s.OnConnStop(conn)
	}
}

// 初始化server
func NewServer() ziface.IServer {
	s := &Server{
		Name:       utils.GlobalObject.Name,
		IPVersion:  "tcp4",
		IP:         utils.GlobalObject.Host,
		Port:       utils.GlobalObject.TcpPort,
		msgHandler: NewMsgHandle(),
		ConnMgr:    NewConnManager(),
	}
	return s
}
