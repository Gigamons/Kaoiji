package packets

import (
	"bytes"
	"io"
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
		_ = packet.WriteBytes(buffer) // teehee, ignoring the error. i'm bad ass
	}

	return buffer.Bytes()
}

func (pw *PacketWriter) WriteBytes(w io.Writer) (err error) {
	for _, packet := range pw.packetList {
		if err = packet.WriteBytes(w); err != nil {
			return
		}
	}

	return
}