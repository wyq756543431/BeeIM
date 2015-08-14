package s2c

import ()

const (
	PacketType_GetLoginToken = iota //获取登陆令牌
	PacketType_Login                //登陆
	PacketType_GetFList				//得到好友列表
	PacketType_SendMsg              //发送信息
	PacketType_Quit                 //退出
)

type Packet struct {
	packetType uint32
	packetData []byte
}

func NewPacket() (packet *Packet) {
	packet = &Packet{0, nil}
	return packet
}

func (self *Packet) SetType(t uint32) {
	self.packetType = t
}

func (self *Packet) GetType() (t uint32) {
	return self.packetType
}

func (self *Packet) SetData(data []byte) {
	self.packetData = data
}

func (self *Packet) GetData() (data []byte) {
	return self.packetData
}
