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

func (player *Player) readPacket() (packet Packet, err error){
	length, err := player.readVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("Packet size: ", length)

	id, err := player.readVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("Packet id: ", id)

	packet, err = player.handlePacket(id, length)
	if err != nil {
		log.Print(err)
		return
	} else if packet != nil {
		log.Println("Packet: ", packet)
		packet.handle(player)
	}
	return
}