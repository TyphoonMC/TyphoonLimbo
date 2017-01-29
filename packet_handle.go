package main

import "log"

type PacketHandshake struct {
	protocol Protocol
	address string
	port uint16
	state State
}
func (packet *PacketHandshake) read(player *Player) (err error) {
	protocol, err := player.readVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.protocol = Protocol(protocol)
	packet.address, err = player.readString()
	if err != nil {
		log.Print(err)
		return
	}
	packet.port, err = player.readUInt16()
	if err != nil {
		log.Print(err)
		return
	}
	state, err := player.readVarInt()
	if err != nil {
		log.Print(err)
		return
	}
	packet.state = State(state)
	return
}
func (packet *PacketHandshake) write(player *Player) (err error) {
	return
}
func (packet *PacketHandshake) handle(player *Player) {
	player.state = packet.state
	player.protocol = packet.protocol
	player.inaddr.address = packet.address
	player.inaddr.port = packet.port
	if player.state == STATUS {

	}
	return
}
