package tool

import (
	"encoding/binary"
)

//uint16转化为[2]byte
func Uint16toBytes(old uint16) []byte {
	var b = make([]byte, 2)
	binary.LittleEndian.PutUint16(b, old)
	return b
}

//string转化为[16]byte
func StringToByte16(s string) [16]byte {
	var b [16]byte
	for i := 0; i < 16; i++ {
		b[i] = s[i]
	}
	return b
}

//[]byte to byte[16]
func BytesToByte16(buffer []byte) [16]byte {
	var b [16]byte
	for i := 0; i < 16; i++ {
		b[i] = buffer[i]
	}
	return b
}

//byte[16] to string
func Byte16ToString(b [16]byte) string {
	var s string
	for _, v := range b {
		s += string(v)
	}
	return s
}
