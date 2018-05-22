package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
)

func (w *Writer) AutoJoinChannel() {
	for i := 0; i < len(objects.CHANNELS); i++ {
		if objects.CHANNELS[i].AutoJoin {
			w.JoinChannel(objects.CHANNELS[i].CInfo.ChannelName)
		}
	}
}

func (w *Writer) ChannelAvaible() {
	for i := 0; i < len(objects.CHANNELS); i++ {
		if !helpers.HasPrivileges(consts.AdminChatMod, w._token.User) && objects.CHANNELS[i].CPerm.AdminOnly {
			continue
		}
		p := NewPacket(constants.BanchoChannelAvailable)
		p.SetPacketData(helpers.MarshalBinary(&objects.CHANNELS[i].CInfo))
		w.Write(p.ToByteArray())
	}
}

func (w *Writer) JoinChannel(channelname string) {
	if objects.JoinChannel(channelname, w._token) {
		w.JoinChannelSuccess(channelname)
	} else {
		w.KickOutOfChannel(channelname)
	}
}

func (w *Writer) JoinChannelSuccess(channelname string) {
	p := NewPacket(constants.BanchoChannelJoinSuccess)
	p.SetPacketData(helpers.BString(channelname))
	w.Write(p.ToByteArray())
}

func (w *Writer) KickOutOfChannel(channelname string) {
	p := NewPacket(constants.BanchoChannelRevoked)
	p.SetPacketData(helpers.BString(channelname))
	w.Write(p.ToByteArray())
}
