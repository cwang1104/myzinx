package ziface

type IMessage interface {
	//获取消息数据段长度
	GetDataLen() uint32
	//获取消息id
	GetMsgId() uint32
	//获取消息内容
	GetData() []byte

	//设计消息ID
	SetMsgId(uint32)
	//设置消息内容
	SetData([]byte)
	//设置消息数据段长度
	SetDataLen(uint32)
}
