package packets

import (
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
)

func (pw *PacketWriter) ChannelJoinSuccess(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelJoinSuccess

	p.WriteData(helpers.GetBytes(channel, true))

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelRevoked(channel string) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelRevoked

	p.WriteData(helpers.GetBytes(channel, true))

	pw.WritePacket(p)
}


func (pw *PacketWriter) ChannelAvailableAutoJoin(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailableAutojoin

	p.WriteData(helpers.GetBytes(channelName, true))
	p.WriteData(helpers.GetBytes(channelTopic, true))
	p.WriteData(helpers.GetBytes(userCount))

	pw.WritePacket(p)
}

func (pw *PacketWriter) ChannelAvailable(channelName string, channelTopic string, userCount int16) {
	p := new(Packet)
	p.PacketId = consts.ServerChannelAvailable

	p.WriteData(helpers.GetBytes(channelName, true))
	p.WriteData(helpers.GetBytes(channelTopic, true))
	p.WriteData(helpers.GetBytes(userCount))

	pw.WritePacket(p)
}

func (pw *PacketWriter) SendMessage(userName string, message string, target string, senderId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSendMessage

	p.WriteData(helpers.GetBytes(userName, true))
	p.WriteData(helpers.GetBytes(message, true))
	p.WriteData(helpers.GetBytes(target, true))
	p.WriteData(helpers.GetBytes(senderId))

	pw.WritePacket(p)
}
