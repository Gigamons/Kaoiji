package packets

import (
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
)

type UserStatsUpdate struct {
	UserId          int32

	Status          byte
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        byte
	BeatmapId       uint32

	RankedScore     int64
	Accuracy        float32
	PlayCount       int32
	TotalScore      int64
	Rank            int32
	PP              int16
}

func (pw *PacketWriter) LoginReply(reply consts.LoginReply) {
	p := new(Packet)
	p.PacketId = consts.ServerLoginReply

	p.WriteData(helpers.GetBytes(int32(reply)))

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceBundle(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceBundle

	p.WriteData(helpers.GetBytes(list))

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceSingle(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceSingle

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) FriendsList(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerFriendsList

	p.WriteData(helpers.GetBytes(list))

	pw.WritePacket(p)
}

func (pw *PacketWriter) UserQuit(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerHandleUserQuit

	p.WriteData(helpers.GetBytes(userId))
	p.WriteData(helpers.GetBytes(int32(0)))

	pw.WritePacket(p)
}

func (pw *PacketWriter) HandleUserStatsUpdate(userStatsUpdate *UserStatsUpdate) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	p.WriteData(helpers.GetBytes(userStatsUpdate.UserId))

	p.WriteData(helpers.GetBytes(userStatsUpdate.Status))
	p.WriteData(helpers.GetBytes(userStatsUpdate.StatusText, true))
	p.WriteData(helpers.GetBytes(userStatsUpdate.BeatmapChecksum, true))
	p.WriteData(helpers.GetBytes(userStatsUpdate.CurrentMods))
	p.WriteData(helpers.GetBytes(userStatsUpdate.PlayMode))
	p.WriteData(helpers.GetBytes(userStatsUpdate.BeatmapId))

	p.WriteData(helpers.GetBytes(userStatsUpdate.RankedScore))
	p.WriteData(helpers.GetBytes(userStatsUpdate.Accuracy))
	p.WriteData(helpers.GetBytes(userStatsUpdate.PlayCount))
	p.WriteData(helpers.GetBytes(userStatsUpdate.TotalScore))
	p.WriteData(helpers.GetBytes(userStatsUpdate.Rank))
	p.WriteData(helpers.GetBytes(userStatsUpdate.PP))

	pw.WritePacket(p)
}

