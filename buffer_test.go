package main

import (
	"testing"
	"bytes"
)

func TestVarBufferAllocation(t *testing.T) {
	size := 20
	buff := newVarBuffer(size)

	if len(buff.buffer) != size {
		t.Log("VarBuffer (size =", size, ") created with size", len(buff.buffer))
		t.Fail()
	}

	if buff.used != 0 {
		t.Log("VarBuffer started with used bytes")
		t.Fail()
	}
}

func TestVarBufferWriteWithoutResize(t *testing.T) {
	size := 20
	buff := newVarBuffer(size)

	dataSize := 10
	data := make([]byte, dataSize)
	buff.Write(data)

	if len(buff.buffer) != size {
		t.Log("VarBuffer (size =", size, ") write with size", len(buff.buffer))
		t.Fail()
	}

	if buff.used != dataSize {
		t.Log("VarBuffer used bytes =", buff.used, " instead of", data)
		t.Fail()
	}

	if !bytes.Equal(data, buff.Bytes()) {
		t.Log("VarBuffer corrupted data")
		t.Fail()
	}
}

func TestVarBufferWriteWithResize(t *testing.T) {
	size := 20
	buff := newVarBuffer(size)

	dataSize := 40
	data := make([]byte, dataSize)
	buff.Write(data)

	if buff.used != dataSize {
		t.Log("VarBuffer used bytes =", buff.used, " instead of", data)
		t.Fail()
	}

	if !bytes.Equal(data, buff.Bytes()) {
		t.Log("VarBuffer corrupted data")
		t.Fail()
	}
}