package ziface

/*
封包数据和拆包数据
直接面向tcp链接中的数据流，为传输数据添加头部信息，用于处理TCP沾包问题
*/
type IDataPack interface {
	//获取包长度方法
	GetHeadLen() uint32
	//封包方法
	Pack(msg IMessage) ([]byte, error)
	//拆包方法
	UnPack([]byte) (IMessage, error)
}
