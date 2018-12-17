package packets

import (
	"bytes"
	"fmt"
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
		err := packet.WriteBytes(buffer)
		if err != nil {
			fmt.Println(err)
			return nil
		}
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