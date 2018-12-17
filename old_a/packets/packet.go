package packets

import (
	"bytes"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/logger"
	"github.com/Mempler/osubinary"
)

// GetPackets get a packet
func GetPackets(pkg []byte) []*constants.Packet {
	var packetList []*constants.Packet
	b := bytes.NewReader(pkg)

	for {
		PacketID, err := osubinary.RUInt16(b)
		if err != nil {
			break
		}
		_, err = osubinary.RInt8(b)
		if err != nil {
			break
		}
		PacketLength, err := osubinary.RUInt32(b)
		if err != nil {
			break
		}
		PacketData := make([]byte, PacketLength)
		lngth, err := b.Read(PacketData)
		if lngth < int(PacketLength) {
			logger.Errorln("Unexpected Packet length! maybe invalid packet?")
			continue
		}
		packetList = append(packetList, &constants.Packet{PacketID, PacketLength, PacketData})
	}

	return packetList
}
