package public

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/handlers/private"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/tools/usertools"
)

func SendUserStats(t *objects.Token, forced bool) []byte {
	w := packets.Writer{}
	x := constants.UserStatsStruct{}
	if forced {
		private.SetUserStatus(t)
	}
	x.UserID = t.User.ID
	x.Status = t.Status.Beatmap.Status
	x.StatusText = t.Status.Beatmap.StatusText
	x.CurrentMods = t.Status.Beatmap.CurrentMods
	x.BeatmapChecksum = t.Status.Beatmap.BeatmapChecksum
	x.PlayMode = t.Status.Beatmap.PlayMode
	x.Accuracy = helpers.CalculateAccuracy(t.Leaderboard.Count300, t.Leaderboard.Count100, t.Leaderboard.Count50, t.Leaderboard.CountMiss, 0, 0, 0)
	x.TotalScore = uint64(t.Leaderboard.TotalScore)
	x.PlayCount = int32(t.Leaderboard.Playcount)
	x.PeppyPoints = t.Leaderboard.PeppyPoints
	x.RankedScore = uint64(t.Leaderboard.RankedScore)
	x.Rank = usertools.GetLeaderboardPosition(t.User, t.Status.Beatmap.PlayMode)
	w.SendUserStats(x)
	return w.Bytes()
}
