package main

type VarBuffer struct {
	used int
	buffer []byte
}

func newVarBuffer(size int) *VarBuffer {
	return &VarBuffer{
		used: 0,
		buffer: make([]byte, size),
	}
}

func arraycpy(dst []byte, dstindex int, src []byte, srcindex int, length int) {
	for i := 0; i < length; i++ {
		dst[i+dstindex] = src[i+srcindex]
	}
	return
}

func (buff *VarBuffer) Len() int {
	return buff.used
}

func (buff *VarBuffer) Bytes() []byte {
	if len(buff.buffer) == buff.used {
		return buff.buffer
	} else {
		return buff.buffer[:buff.used]
	}
}

func (buff *VarBuffer) Write(p []byte) (n int, err error) {
	if len(buff.buffer)-buff.used < len(p) {
		size := len(p)-len(buff.buffer)-buff.used
		nbuffer := make([]byte, size)
		copy(nbuffer, buff.buffer)
		buff.buffer = nbuffer
	}
	arraycpy(buff.buffer, buff.used, p, 0, len(p))
	buff.used += len(p)
	return len(p), nil
}

func (buff *VarBuffer) Read(p []byte) (n int, err error) {
	return 0, nil
}