package main

import (
	"net"
	"log"
	"bufio"
)

type State int8
const (
	HANDSHAKING State = iota
	STATUS
	LOGIN
	PLAY
)

type Protocol uint8
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
}

func (player *Player) ReadPacket() (packet Packet, err error){
	length, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}

	id, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}

	packet, err = player.HandlePacket(id, length)
	if err != nil {
		log.Print(err)
		return
	} else if packet != nil {
		log.Println("Packet: ", packet)
		packet.Handle(player)
	}
	return
}

func (player *Player) WritePacket(packet Packet) (err error){
	buff := newVarBuffer(256)
	tmp := player.io
	player.io = &ConnReadWrite{
		rdr: tmp.rdr,
		wtr: bufio.NewWriter(buff),
	}

	buff.Write([]byte("test1"))
	player.io.wtr.Write([]byte("test2"))

	id := packet.Id()
	player.WriteVarInt(id)
	packet.Write(player)

	player.io = tmp
	player.WriteVarInt(buff.Len())
	log.Println(buff.Len())
	player.io.wtr.Write(buff.Bytes())
	return nil
}