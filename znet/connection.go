package znet

import (
	"fmt"
	"myzinx/ziface"
	"net"
)

type Connection struct {
	Conn   *net.TCPConn
	ConnID uint32

	IsClosed bool

	//该链接的处理方法router
	Router ziface.IRouter

	//告知当前连接已经退出/停止的 channel
	ExitChan chan bool
}

// 连接的读业务方法
func (c *Connection) StartReader() {
	fmt.Println("[Reader go is running]")
	defer fmt.Printf("connid=%d,reader is exit,remote addr is %v", c.ConnID, c.RemoteAddr())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err")
			continue
		}

		//得到当前客户端请求的request数据
		req := Request{
			conn: c,
			data: buf,
		}

		//从路由routers中找到注册绑定Conn对应的Handle
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

		////调用当前连接绑定的handleApi
		//err = c.handleApi(c.Conn, buf, cnt)
		//if err != nil {
		//	fmt.Println("ConnId handle is error", err)
		//	break
		//}
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start,connid=", c.ConnID)
	//启动从当前连接的读数据业务
	go c.StartReader()
	//todo:启动从当前连接写数据的业务
	for {
		select {
		case <-c.ExitChan:
			//得到退出消息 不再阻塞
			return
		}
	}

}

func (c *Connection) Stop() {

	fmt.Println("conn stop,connid=", c.ConnID)

	if c.IsClosed == true {
		return
	}

	c.IsClosed = true
	c.Conn.Close()

	//通知从缓冲队列读取数据的业务，该链接已经关闭
	c.ExitChan <- true

	//回收资源
	close(c.ExitChan)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
func (c *Connection) Send(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connID,
		Router:   router,
		IsClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
