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
}
