package packets

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"git.gigamons.de/Gigamons/Kaoiji/constants/packets"
)

func (w *Writer) SendUserStats(x packetconst.UserStats) {
	p := NewPacket(constants.BanchoHandleOsuUpdate)
	p.SetPacketData(MarshalBinary(&x))
	w.Write(p.ToByteArray())
}
