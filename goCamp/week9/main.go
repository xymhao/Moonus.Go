package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	data := WriteTo("Hello, Moonus!")
	decoding(data)
}

const (
	// MaxBodySize max proto body size
	MaxBodySize = int32(1 << 12)
)

const (
	// size
	_packSize      = 4
	_headerSize    = 2
	_verSize       = 2
	_opSize        = 4
	_seqSize       = 4
	_heartSize     = 4
	_rawHeaderSize = _packSize + _headerSize + _verSize + _opSize + _seqSize
	_maxPackSize   = MaxBodySize + int32(_rawHeaderSize)
	// offset
	_packOffset   = 0
	_headerOffset = _packOffset + _packSize
	_verOffset    = _headerOffset + _headerSize
	_opOffset     = _verOffset + _verSize
	_seqOffset    = _opOffset + _opSize
	_heartOffset  = _seqOffset + _seqSize
)

/*
goim 协议结构
4bytes PacketLen 包长度，在数据流传输过程中，先写入整个包的长度，方便整个包的数据读取。
2bytes HeaderLen 头长度，在处理数据时，会先解析头部，可以知道具体业务操作。
2bytes Version 协议版本号，主要用于上行和下行数据包按版本号进行解析。
4bytes Operation 业务操作码，可以按操作码进行分发数据包到具体业务当中。
4bytes Sequence 序列号，数据包的唯一标记，可以做具体业务处理，或者数据包去重。
PacketLen-HeaderLen Body 实际业务数据，在业务层中会进行数据解码和编码。
*/

func decoding(data []byte) {
	if len(data) <= 16 {
		fmt.Println("data len < 16.")
		return
	}
	packetLen := binary.BigEndian.Uint32(data[_packOffset:_headerOffset])
	fmt.Printf("packetLen:%v\n", packetLen)

	headerLen := binary.BigEndian.Uint16(data[_headerOffset:_verOffset])
	fmt.Printf("headerLen:%v\n", headerLen)

	version := binary.BigEndian.Uint16(data[_verOffset:_opOffset])
	fmt.Printf("version:%v\n", version)

	operation := binary.BigEndian.Uint32(data[_opOffset:_seqOffset])
	fmt.Printf("operation:%v\n", operation)

	sequence := binary.BigEndian.Uint32(data[_seqOffset:_heartOffset])
	fmt.Printf("sequence:%v\n", sequence)

	body := string(data[_heartOffset:])
	fmt.Printf("body:%v\n", body)
}

func WriteTo(body string) []byte {
	packLen := len(body) + _rawHeaderSize
	ret := make([]byte, packLen)

	binary.BigEndian.PutUint32(ret[_packOffset:], uint32(packLen))
	binary.BigEndian.PutUint16(ret[_headerOffset:], uint16(_rawHeaderSize))

	version := 5
	binary.BigEndian.PutUint16(ret[_verOffset:], uint16(version))
	operation := 6
	binary.BigEndian.PutUint32(ret[_opOffset:], uint32(operation))
	sequence := 7
	binary.BigEndian.PutUint32(ret[_seqOffset:], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[_rawHeaderSize:], byteBody)
	return ret
}
