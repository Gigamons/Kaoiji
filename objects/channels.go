package objects

type channelinfo struct {
	ChannelName        string
	ChannelDescription string
	UserCount          int16
}

type channelpermissions struct {
	AdminOnly bool
	ReadOnly  bool
	WriteOnly bool
}

type channel struct {
	CInfo    channelinfo
	CPerm    channelpermissions
	AutoJoin bool
}

var CHANNELS []channel

func init() {
	CHANNELS = append(CHANNELS, channel{CInfo: channelinfo{ChannelName: "#osu", ChannelDescription: "Default osu! Channel"}, CPerm: channelpermissions{}, AutoJoin: true})
	CHANNELS = append(CHANNELS, channel{CInfo: channelinfo{ChannelName: "#announce", ChannelDescription: "Default osu! Channel"}, CPerm: channelpermissions{ReadOnly: true}, AutoJoin: true})
	CHANNELS = append(CHANNELS, channel{CInfo: channelinfo{ChannelName: "#userlog", ChannelDescription: "Default osu! Channel"}, CPerm: channelpermissions{ReadOnly: true}})
	CHANNELS = append(CHANNELS, channel{CInfo: channelinfo{ChannelName: "#lobby", ChannelDescription: "Default osu! Channel"}, CPerm: channelpermissions{}})
}
