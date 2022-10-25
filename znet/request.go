package znet

import "myzinx/ziface"

type Request struct {
	// 已经和客户端建立好的连接
	conn ziface.IConnection
	data []byte //客户端请求的数据
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
