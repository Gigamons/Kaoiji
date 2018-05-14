package objects

import (
	"git.gigamons.de/Gigamons/Kaoiji/constants/privileges"
	"git.gigamons.de/Gigamons/Kaoiji/helpers"
)

type ChannelInfo struct {
	ChannelName        string
	ChannelDescription string
	UserCount          int16
}

type ChannelPermissions struct {
	AdminOnly bool
	ReadOnly  bool
	WriteOnly bool
}

type Channel struct {
	CInfo    ChannelInfo
	CPerm    ChannelPermissions
	AutoJoin bool
}

var CHANNELS []Channel

func init() {
	CHANNELS = append(CHANNELS, Channel{CInfo: ChannelInfo{ChannelName: "#osu", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{}, AutoJoin: true})
	CHANNELS = append(CHANNELS, Channel{CInfo: ChannelInfo{ChannelName: "#announce", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{ReadOnly: true}, AutoJoin: true})
	CHANNELS = append(CHANNELS, Channel{CInfo: ChannelInfo{ChannelName: "#userlog", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{ReadOnly: true}})
	CHANNELS = append(CHANNELS, Channel{CInfo: ChannelInfo{ChannelName: "#lobby", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{}})
	CHANNELS = append(CHANNELS, Channel{CInfo: ChannelInfo{ChannelName: "#admin", ChannelDescription: "Administrator channel of all Admins"}, CPerm: ChannelPermissions{AdminOnly: true}, AutoJoin: true})
}

func GetChannel(channelname string) *Channel {
	for i := 0; i < len(CHANNELS); i++ {
		if CHANNELS[i].CInfo.ChannelName == channelname {
			return &CHANNELS[i]
		}
	}
	return nil
}

func JoinChannel(channelname string, t *Token) bool {
	channel := GetChannel(channelname)
	if channel == nil {
		return false
	}
	if helpers.HasPrivileges(privileges.AdminChatMod, t.User) && channel.CPerm.AdminOnly {
		channel.CInfo.UserCount++
		return true
	} else if !helpers.HasPrivileges(privileges.AdminChatMod, t.User) && channel.CPerm.AdminOnly {
		return false
	}
	channel.CInfo.UserCount++
	return true
}

func LeaveChannel(channelname string) bool {
	channel := GetChannel(channelname)
	if channel == nil {
		return false
	}
	channel.CInfo.UserCount--
	return true
}
