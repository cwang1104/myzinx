package znet

import (
	"fmt"
	"myzinx/ziface"
	"strconv"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter //存放每个msgID对应的处理方法
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// 马上以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgID()]
	if !ok {
		fmt.Println("api msgId not found id =  ", request.GetMsgID())
		return
	}

	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)

}

// 为消息添加具体处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {

	//判断当前msg绑定的api处理方法是否存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api, msgId = " + strconv.Itoa(int(msgId)))
	}

	mh.Apis[msgId] = router
	fmt.Println("add api msgId = ", msgId)
}
