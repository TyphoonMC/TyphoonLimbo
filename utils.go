package main

import (
	"io"
	"encoding/binary"
)

type ConnReadWrite struct {
	rdr io.Reader
	wtr io.Writer
	buffer [16]byte
}

func (rdrwtr ConnReadWrite) ReadByte() (b byte, err error) {
	buff := rdrwtr.buffer[:1]
	if _, err = rdrwtr.rdr.Read(buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (player *Player) ReadByte() (b byte, err error){
	buff := player.io.buffer[:1]
	if _, err := io.ReadFull(player.conn, buff); err != nil {
		return 0, err
	}
	return buff[0], nil
}

func (player *Player) ReadVarInt() (i int, err error){
	v, err := binary.ReadUvarint(player.io)
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (player *Player) WriteVarInt(i int) (err error){
	buff := player.io.buffer[:]
	length := binary.PutUvarint(buff, uint64(i))
	_, err = player.io.wtr.Write(buff[:length])
	if err != nil {
		return err
	}
	return nil
}

func (player *Player) ReadUInt16() (i uint16, err error){
	buff := player.io.buffer[:2]
	_, err = io.ReadFull(player.io.rdr, buff)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(buff), nil
}

func (player *Player) ReadString() (s string, err error){
	length, err := player.ReadVarInt()
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

func (player *Player) WriteString(s string) (err error){
	buff := []byte(s)
	err = player.WriteVarInt(len(buff))
	if err != nil {
		return err
	}
	_, err = player.io.wtr.Write(buff)
	if err != nil {
		return err
	}
	return nil
}