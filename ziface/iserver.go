package ziface

// 定义接口
type IServer interface {
	//开启服务
	Start()
	//停止服务
	Stop()
	//运行服务
	Serve()

	//增加路由功能：给当前服务注册一个路由业务方法，供客户端链接处理使用
	AddRouter(msgId uint32, router IRouter)

	//得到连接管理方法
	GetConnMgr() IConManager

	//设置该server的连接创建时的hook函数
	SetOnConnStart(func(IConnection))
	//设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConnection))
	//调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	//调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}
