package objects

import (
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
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

// CHANNELS Global Variable
var CHANNELS []*Channel

func init() {
	CHANNELS = append(CHANNELS, &Channel{CInfo: ChannelInfo{ChannelName: "#osu", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{}, AutoJoin: true})
	CHANNELS = append(CHANNELS, &Channel{CInfo: ChannelInfo{ChannelName: "#announce", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{ReadOnly: true}, AutoJoin: true})
	CHANNELS = append(CHANNELS, &Channel{CInfo: ChannelInfo{ChannelName: "#userlog", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{ReadOnly: true}})
	CHANNELS = append(CHANNELS, &Channel{CInfo: ChannelInfo{ChannelName: "#lobby", ChannelDescription: "Default osu! Channel"}, CPerm: ChannelPermissions{}})
	CHANNELS = append(CHANNELS, &Channel{CInfo: ChannelInfo{ChannelName: "#admin", ChannelDescription: "Administrator channel of all Admins"}, CPerm: ChannelPermissions{AdminOnly: true}, AutoJoin: true})
}

// GetChannel returns an Channel object, if not exists return nil
func GetChannel(channelname string) *Channel {
	for i := 0; i < len(CHANNELS); i++ {
		if CHANNELS[i].CInfo.ChannelName == channelname {
			return CHANNELS[i]
		}
	}
	return nil
}

// JoinChannel add a User to the Defined channel. returns bool, if Successfull return true else false
func JoinChannel(channelname string, t *Token) bool {
	channel := GetChannel(channelname)
	if channel == nil {
		return false
	}
	if helpers.HasPrivileges(consts.AdminChatMod, t.User) && channel.CPerm.AdminOnly {
		channel.CInfo.UserCount++
		return true
	} else if !helpers.HasPrivileges(consts.AdminChatMod, t.User) && channel.CPerm.AdminOnly {
		return false
	}
	channel.CInfo.UserCount++
	return true
}

// LeaveChannel remove a User from the Defined channel. returns bool, if Successfull return true else false
func LeaveChannel(channelname string) bool {
	channel := GetChannel(channelname)
	if channel == nil {
		return false
	}
	channel.CInfo.UserCount--
	return true
}
