package packets

import (
	"github.com/Gigamons/Kaoiji/consts"
	"github.com/Gigamons/Shared/shelpers"
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

	shelpers.WriteBytes(&p.buffer, int32(reply))

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceBundle(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceBundle

	shelpers.WriteBytes(&p.buffer, list)

	pw.WritePacket(p)
}

func (pw *PacketWriter) PresenceSingle(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerUserPresenceSingle

	shelpers.WriteBytes(&p.buffer, userId)

	pw.WritePacket(p)
}

func (pw *PacketWriter) FriendsList(list []int32) {
	p := new(Packet)
	p.PacketId = consts.ServerFriendsList

	shelpers.WriteBytes(&p.buffer, list)

	pw.WritePacket(p)
}

func (pw *PacketWriter) UserQuit(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerHandleUserQuit

	shelpers.WriteBytes(&p.buffer, userId)
	shelpers.WriteBytes(&p.buffer, int32(0))

	pw.WritePacket(p)
}

func (pw *PacketWriter) HandleUserStatsUpdate(userStatsUpdate *UserStatsUpdate) {
	p := new(Packet)
	p.PacketId = consts.ServerAnnounce

	shelpers.WriteBytes(&p.buffer, userStatsUpdate.UserId)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.UserId)

	shelpers.WriteBytes(&p.buffer, userStatsUpdate.Status)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.StatusText, true)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.BeatmapChecksum, true)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.CurrentMods)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.PlayMode)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.BeatmapId)

	shelpers.WriteBytes(&p.buffer, userStatsUpdate.RankedScore)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.Accuracy)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.PlayCount)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.TotalScore)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.Rank)
	shelpers.WriteBytes(&p.buffer, userStatsUpdate.PP)

	pw.WritePacket(p)
}

