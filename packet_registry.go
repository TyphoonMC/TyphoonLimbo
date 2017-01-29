package main

import (
	"reflect"
	"errors"
)

var (
	packets map[int64]reflect.Type = make(map[int64]reflect.Type)
)

type Packet interface {
	write(*Player) error
	read(*Player) error
	handle(*Player)
}

func packetTypeHash(state State, id int) int64 {
	return int64(id)^(int64(state) << 32)
}

func initPackets() {
	packets[packetTypeHash(HANDSHAKING, 0x00)] = reflect.TypeOf((*PacketHandshake)(nil)).Elem()
}

func (player *Player) handlePacket(id int, length int) (packet Packet, err error) {
	typ := packets[packetTypeHash(player.state, id)];

	if typ == nil {
		return nil, errors.New("Unknown packet")
	}

	packet, _ = reflect.New(typ).Interface().(Packet)
	if err = packet.read(player); err != nil {
		return nil, err
	}
	return
}