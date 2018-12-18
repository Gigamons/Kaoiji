package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
)

func (pw *PacketWriter) ChannelJoinSuccess(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelJoinSuccess

	helpers.WriteBytes(&p.buffer, channel, true)

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelRevoked(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelRevoked

	helpers.WriteBytes(&p.buffer, channel, true)

	pw.WritePacket(p)
}


func (pw *PacketWriter) ChannelAvailableAutoJoin(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailableAutojoin

	helpers.WriteBytes(&p.buffer, channelName, true)
	helpers.WriteBytes(&p.buffer, channelTopic, true)
	helpers.WriteBytes(&p.buffer, userCount)

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelAvailable(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailable

	helpers.WriteBytes(&p.buffer, channelName, true)
	helpers.WriteBytes(&p.buffer, channelTopic, true)
	helpers.WriteBytes(&p.buffer, userCount)

	pw.WritePacket(p)
}

func (pw *PacketWriter) SendMessage(userName string, message string, target string, senderId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSendMessage


	helpers.WriteBytes(&p.buffer, userName, true)
	helpers.WriteBytes(&p.buffer, message, true)
	helpers.WriteBytes(&p.buffer, target, true)
	helpers.WriteBytes(&p.buffer, senderId)

	pw.WritePacket(p)
}
