/*
author:jaydenhe
 */
package s2c

import (
	"github.com/go-xweb/log"
	uuid "github.com/satori/go.uuid"
	"net"
)

/*
SessionID 64bit

|uid (32bit)|extendId (32bit)|

SessionID = uid<<32 + extendId
 */
type TypeSessionID uuid.UUID

type Session struct {
	conn     net.Conn
	incoming chan Packet
	outgoing chan Packet
	packetReader  *PacketReader
	packetWriter  *PacketWriter
	quiting  chan byte
	name     string
	id       TypeSessionID
	server   *Server
}

func (self *Session) GetName() string {
	return self.name
}

func (self *Session) SetName(name string) {
	self.name = name
}

func (self *Session) GetID() TypeSessionID {
	return self.id
}

func (self *Session) SetID(id TypeSessionID) {
	self.id = id
}

func (self *Session) GetConn() net.Conn {
	return self.conn
}

func (self *Session) GetIncoming() Packet {
	return <-self.incoming
}

func (self *Session) PutOutgoing(packet Packet) {
	self.outgoing <- packet
}

func CreateSession(server *Server,conn net.Conn) *Session {
	packetReader := NewPacketReader(conn)
	packetWriter := NewPacketWriter(conn)

	Session := &Session{
		conn:     conn,
		incoming: make(chan Packet),
		outgoing: make(chan Packet),
		quiting:  make(chan byte),
		packetReader:   packetReader,
		packetWriter:   packetWriter,
		server:server,
	}
	log.Printf("create session [ %+v ]\n",Session)
	Session.Listen(server)
	return Session
}

func (self *Session) Listen(server *Server) {
	go self.Read()
	go self.Write()
}

func (self *Session) quit() {
	self.quiting <- 0
}

func (self *Session) Read() {

	for {
		if packet,err := self.packetReader.ReadAPacket();err == nil {
			log.Printf("放入进入信息通道%+v",packet)
			self.incoming <- *packet
		}else{
			log.Println("Read error:",err)
			self.quit()
			return
		}
	}
}

func (self *Session) Write() {

	for packet := range self.outgoing {
		if err := self.packetWriter.WriteAPacket(&packet);err != nil {
			log.Println("Write a packet error:",err)
			self.quit()
			return
		}

		if err := self.packetWriter.Flush();err != nil {
			log.Println("Write error:",err)
			self.quit()
			return
		}
	}

}
