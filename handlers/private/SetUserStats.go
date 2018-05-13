package private

import (
	"git.gigamons.de/Gigamons/Kaoiji/objects"
	"git.gigamons.de/Gigamons/Kaoiji/tools/usertools"
)

func SetUserStatus(t *objects.Token) {
	t.Leaderboard = usertools.GetLeaderboard(t.User, t.Status.Beatmap.PlayMode)
	t.Leaderboard.Position = usertools.GetLeaderboardPosition(t.User, t.Status.Beatmap.PlayMode)
}
