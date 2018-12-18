package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Kaoiji/helpers"
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

	helpers.WriteBytes(&p.buffer, int32(reply))

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceBundle(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceBundle

	helpers.WriteBytes(&p.buffer, list)

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceSingle(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceSingle

	helpers.WriteBytes(&p.buffer, userId)

	pw.WritePacket(p)
}

func (pw *PacketWriter) FriendsList(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerFriendsList

	helpers.WriteBytes(&p.buffer, list)

	pw.WritePacket(p)
}

func (pw *PacketWriter) UserQuit(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerHandleUserQuit

	helpers.WriteBytes(&p.buffer, userId)
	helpers.WriteBytes(&p.buffer, int32(0))

	pw.WritePacket(p)
}

func (pw *PacketWriter) HandleUserStatsUpdate(userStatsUpdate *UserStatsUpdate) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	helpers.WriteBytes(&p.buffer, userStatsUpdate.UserId)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.UserId)

	helpers.WriteBytes(&p.buffer, userStatsUpdate.Status)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.StatusText, true)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.BeatmapChecksum, true)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.CurrentMods)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.PlayMode)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.BeatmapId)

	helpers.WriteBytes(&p.buffer, userStatsUpdate.RankedScore)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.Accuracy)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.PlayCount)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.TotalScore)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.Rank)
	helpers.WriteBytes(&p.buffer, userStatsUpdate.PP)

	pw.WritePacket(p)
}

