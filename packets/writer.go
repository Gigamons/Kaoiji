package packets

import (
	"bytes"
)

type PacketWriter struct {
	packetList []*Packet
}

func (pw *PacketWriter) WritePacket(p *Packet) {
	pw.packetList = append(pw.packetList, p)
}

func (pw *PacketWriter) GetBytes() []byte {
	buffer := new(bytes.Buffer)

	for _, packet := range pw.packetList {
		buffer.Write(packet.GetBytes())
	}


	return buffer.Bytes()
}
