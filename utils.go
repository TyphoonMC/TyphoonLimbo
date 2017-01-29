package main

import (
	"io"
	"encoding/binary"
)

type ConnReadWrite struct {
	rdr io.Reader
	wtr io.ByteWriter
	buffer [16]byte
}

func (rdrwtr ConnReadWrite) ReadByte() (b byte, err error) {
	buff := rdrwtr.buffer[:1]
	if _, err = rdrwtr.rdr.Read(buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (player *Player) readByte() (b byte, err error){
	buff := player.io.buffer[:1]
	if _, err := io.ReadFull(player.conn, buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (player *Player) readVarInt() (i int, err error){
	v, err := binary.ReadUvarint(player.io)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (player *Player) readUInt16() (i uint16, err error){
	buff := player.io.buffer[:2]
	_, err = io.ReadFull(player.io.rdr, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buff), nil
}

func (player *Player) readString() (s string, err error){
	length, err := player.readVarInt()
	if err != nil {
		return "", err
	}
	buffer := make([]byte,length)
	_, err = io.ReadFull(player.io.rdr, buffer)
	if err != nil {
		return "", err
	}
	return string(buffer), nil
}