package ziface

/*
	路由接口---使用框架者给链接自定的处理业务方法
	路由里的IRequest包含用该链接的连接信息和该链接的请求数据信息
*/

type IRouter interface {
	//在处理conn业务之前的钩子方法
	PreHandle(request IRequest)
	//处理conn业务的方法
	Handle(request IRequest)
	//在处理conn业务之后的钩子方法
	PostHandle(request IRequest)
}
