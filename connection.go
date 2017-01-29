package main

import (
	"net"
	"log"
)

type State int8
const (
	HANDSHAKING State = iota
	STATUS
	LOGIN
	PLAY
)

type Protocol uint16
const (
	V1_10 Protocol = 210
)

type InAddr struct {
	address string
	port uint16
}

type Player struct {
	conn net.Conn
	io *ConnReadWrite
	state State
	protocol Protocol
	inaddr InAddr
	name string
	uuid string
	keepalive int
}

func (player *Player) ReadPacket() (packet Packet, err error){
	length, err := player.ReadVarInt()
	if err != nil {
		return
	}

	id, err := player.ReadVarInt()
	if err != nil {
		return
	}

	packet, err = player.HandlePacket(id, length)
	if err != nil {
		return
	} else if packet != nil {
		log.Println("->", id, packet)
		packet.Handle(player)
	}
	return
}

func (player *Player) WritePacket(packet Packet) (err error){
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: buff,
	}

	id := packet.Id()
	player.WriteVarInt(id)
	packet.Write(player)

	ln := newVarBuffer(0)
	player.io.wtr = ln
	player.WriteVarInt(buff.Len())
	player.io = tmp
	player.conn.Write(ln.Bytes())
	player.conn.Write(buff.Bytes())

	log.Println("<-", id, packet)
	return nil
}