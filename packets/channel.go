package packets

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"git.gigamons.de/Gigamons/Kaoiji/objects"
)

func (w *Writer) AutoJoinChannel() {
	for i := 0; i < len(objects.CHANNELS); i++ {
		if objects.CHANNELS[i].AutoJoin {
			p := NewPacket(constants.BanchoChannelJoinSuccess)
			p.SetPacketData(BString(objects.CHANNELS[i].CInfo.ChannelName))
			w.Write(p.ToByteArray())
		}
	}
}

func (w *Writer) ChannelAvaible() {
	for i := 0; i < len(objects.CHANNELS); i++ {
		p := NewPacket(constants.BanchoChannelAvailable)
		p.SetPacketData(MarshalBinary(&objects.CHANNELS[i].CInfo))
		w.Write(p.ToByteArray())
	}
}
