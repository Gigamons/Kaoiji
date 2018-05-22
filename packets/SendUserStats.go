package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/constants/packets"
)

func (w *Writer) SendUserStats(x packetconst.UserStats) {
	p := NewPacket(constants.BanchoHandleOsuUpdate)
	p.SetPacketData(MarshalBinary(&x))
	w.Write(p.ToByteArray())
}
