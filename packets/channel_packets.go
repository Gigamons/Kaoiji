package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Shared/shelpers"
)

func (pw *PacketWriter) ChannelJoinSuccess(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelJoinSuccess

	shelpers.WriteBytes(&p.buffer, channel, true)

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelRevoked(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelRevoked

	shelpers.WriteBytes(&p.buffer, channel, true)

	pw.WritePacket(p)
}


func (pw *PacketWriter) ChannelAvailableAutoJoin(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailableAutojoin

	shelpers.WriteBytes(&p.buffer, channelName, true)
	shelpers.WriteBytes(&p.buffer, channelTopic, true)
	shelpers.WriteBytes(&p.buffer, userCount)

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelAvailable(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailable

	shelpers.WriteBytes(&p.buffer, channelName, true)
	shelpers.WriteBytes(&p.buffer, channelTopic, true)
	shelpers.WriteBytes(&p.buffer, userCount)

	pw.WritePacket(p)
}

func (pw *PacketWriter) SendMessage(userName string, message string, target string, senderId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSendMessage


	shelpers.WriteBytes(&p.buffer, userName, true)
	shelpers.WriteBytes(&p.buffer, message, true)
	shelpers.WriteBytes(&p.buffer, target, true)
	shelpers.WriteBytes(&p.buffer, senderId)

	pw.WritePacket(p)
}
