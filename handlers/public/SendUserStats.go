package public

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/handlers/private"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/tools/usertools"
)

// SendUserStats sends a Status to given client. nor return a byte array
func SendUserStats(t *objects.Token, forced bool) []byte {
	w := packets.Writer{}
	x := constants.UserStatsStruct{}
	if forced {
		private.SetUserStatus(t)
	}
	x.UserID = t.User.ID
	x.Status = t.Status.Beatmap.Status
	x.StatusText = t.Status.Beatmap.StatusText
	x.BeatmapChecksum = t.Status.Beatmap.BeatmapChecksum
	x.CurrentMods = t.Status.Beatmap.CurrentMods
	x.PlayMode = t.Status.Beatmap.PlayMode
	x.BeatmapID = t.Status.Beatmap.BeatmapID
	x.RankedScore = uint64(t.Leaderboard.RankedScore)
	x.Accuracy = float32(helpers.CalculateAccuracy(t.Leaderboard.Count300, t.Leaderboard.Count100, t.Leaderboard.Count50, t.Leaderboard.CountMiss, 0, 0, 0))
	x.PlayCount = int32(t.Leaderboard.Playcount)
	x.TotalScore = uint64(t.Leaderboard.TotalScore)
	x.Rank = usertools.GetLeaderboardPosition(t.User, t.Status.Beatmap.PlayMode)
	x.PeppyPoints = int16(t.Leaderboard.PeppyPoints)
	w.SendUserStats(x)
	return w.Bytes()
}
