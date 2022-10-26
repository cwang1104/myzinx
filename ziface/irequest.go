package ziface

//存放客户端请求的连接信息和请求数据
//我们可以从request里得到全部客户端的信息

type IRequest interface {
	// 获取请求连接信息
	GetConnection() IConnection
	// 获取请求消息的数据
	GetData() []byte

	GetMsgID() uint32
}
