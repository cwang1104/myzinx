package znet

type Message struct {
	Id      uint32
	DataLen uint32
	Data    []byte
}

func NewMessagePackage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		Data:    data,
		DataLen: uint32(len(data)),
	}
}

func (mh *Message) GetDataLen() uint32 {
	return mh.DataLen
}

// 获取消息ID
func (mh *Message) GetMsgId() uint32 {
	return mh.Id
}

// 获取消息内容
func (mh *Message) GetData() []byte {
	return mh.Data
}

// 设置消息数据段长度
func (mh *Message) SetDataLen(len uint32) {
	mh.DataLen = len
}

// 设计消息ID
func (mh *Message) SetMsgId(msgId uint32) {
	mh.Id = msgId
}

// 设计消息内容
func (mh *Message) SetData(data []byte) {
	mh.Data = data
}
