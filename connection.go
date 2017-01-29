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
	V1_7_2 Protocol = 4
	V1_7_6 = 5
	V1_8 = 47
	V1_9 = 107
	V1_9_1 = 108
	V1_9_2 = 109
	V1_9_3 = 110
	V1_10 = 210
	V1_11 = 315
	V1_11_1 = 316
)
var (
	COMPATIBLE_PROTO = [...]Protocol{V1_10,V1_11,V1_11_1}
)

func IsCompatible(proto Protocol) bool {
	for _, x := range COMPATIBLE_PROTO {
		if x == proto {
			return true
		}
	}
	return false
}

type InAddr struct {
	address string
	port uint16
}

type Player struct {
	id int
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