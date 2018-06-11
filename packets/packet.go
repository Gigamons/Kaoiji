package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
)

// Packet struct is simply a packet
type Packet struct {
	PacketID     int16
	PacketLength int32
	PacketData   []byte
}

// NewPacket Create a new Packet
func NewPacket(packetid int) Packet {
	return Packet{PacketID: int16(packetid)}
}

// SetPacketData set's the Packet data
func (p *Packet) SetPacketData(PacketData []byte) {
	p.PacketLength = int32(len(PacketData))
	p.PacketData = PacketData
}

// ToByteArray retuns a ByteArray of the written Packet.
func (p *Packet) ToByteArray() []byte {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, p.PacketID)
	binary.Write(b, binary.LittleEndian, int8(0))
	binary.Write(b, binary.LittleEndian, p.PacketLength)
	binary.Write(b, binary.LittleEndian, p.PacketData)

	return b.Bytes()
}

// Unmarshal a packet
func GetPackets(pkg []byte) []Packet {
	packetList := []Packet{}
	b := bytes.NewReader(pkg)

	for {
		PacketID, err := helpers.RInt16(b)
		if err != nil {
			break
		}
		_, err = helpers.RInt8(b)
		if err != nil {
			break
		}
		PacketLength, err := helpers.RInt32(b)
		if err != nil {
			break
		}
		PacketData := make([]byte, PacketLength)
		lngth, err := b.Read(PacketData)
		if lngth < int(PacketLength) {
			logger.Errorln("Unexpected Packet length! maybe invalid packet?")
			continue
		}
		packetList = append(packetList, Packet{PacketID: PacketID, PacketLength: PacketLength, PacketData: PacketData})
	}

	return packetList
}
