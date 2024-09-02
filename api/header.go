package api

import "encoding/binary"

type Header []byte

func getType(header Header) uint16 {
	return binary.LittleEndian.Uint16(header[:2])
}
