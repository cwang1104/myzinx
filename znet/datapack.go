package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"myzinx/utils"
	"myzinx/ziface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	return &DataPack{}
}

// 获取包头长度方法
func (d *DataPack) GetHeadLen() uint32 {
	//Id uint32 （4字节） + DataLen（uint32）4字节
	return 8
}

// 封包方法
func (d *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写datalen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataBuff.Bytes(), nil

}

// 拆包方法
func (d *DataPack) UnPack(binaryData []byte) (ziface.IMessage, error) {

	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//直接压head的信息，得到dataLen和msgID
	msg := &Message{}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	//判断datalen长度是否超出我们允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	//这里把head的数据拆包出来即可，再通过head长度再从conn读取一次数据
	return msg, nil
}
