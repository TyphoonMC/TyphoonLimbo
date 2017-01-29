package main

import (
	"reflect"
	"errors"
)

var (
	packets map[int64]reflect.Type = make(map[int64]reflect.Type)
)

type Packet interface {
	Write(*Player) error
	Read(*Player) error
	Handle(*Player)
	Id() int
}

func PacketTypeHash(state State, id int) int64 {
	return int64(id)^(int64(state) << 32)
}

func InitPackets() {
	packets[PacketTypeHash(HANDSHAKING, 0x00)] = reflect.TypeOf((*PacketHandshake)(nil)).Elem()
	packets[PacketTypeHash(STATUS, 0x00)] = reflect.TypeOf((*PacketStatusRequest)(nil)).Elem()
}

func (player *Player) HandlePacket(id int, length int) (packet Packet, err error) {
	typ := packets[PacketTypeHash(player.state, id)];

	if typ == nil {
		return nil, errors.New("Unknown packet")
	}

	packet, _ = reflect.New(typ).Interface().(Packet)
	if err = packet.Read(player); err != nil {
		return nil, err
	}
	return
}