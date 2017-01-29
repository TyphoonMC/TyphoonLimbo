package main

import (
	"net"
	"log"
	"bytes"
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
	io ConnReadWrite
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
	log.Println("Packet size: ", length)

	id, err := player.ReadVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("Packet id: ", id)

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
	//TODO
}