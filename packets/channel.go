package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Mempler/osubinary"
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
		p := constants.NewPacket(constants.BanchoChannelAvailable)
		p.SetPacketData(osubinary.Marshal(objects.CHANNELS[i].CInfo))
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

func (w *Writer) LeaveChannel(channelname string) {
	objects.LeaveChannel(channelname)
	w.KickOutOfChannel(channelname)
}

func (w *Writer) JoinChannelSuccess(channelname string) {
	p := constants.NewPacket(constants.BanchoChannelJoinSuccess)
	p.SetPacketData(osubinary.BString(channelname))
	w.Write(p.ToByteArray())
}

func (w *Writer) KickOutOfChannel(channelname string) {
	p := constants.NewPacket(constants.BanchoChannelRevoked)
	p.SetPacketData(osubinary.BString(channelname))
	w.Write(p.ToByteArray())
}
