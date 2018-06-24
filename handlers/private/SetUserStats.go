package private

import (
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/tools/usertools"
)

// SetUserStatus of given Token
func SetUserStatus(t *objects.Token) {
	t.Leaderboard = usertools.GetLeaderboard(t.User, int8(t.Status.Beatmap.PlayMode))
	t.Leaderboard.Position = usertools.GetLeaderboardPosition(t.User, int8(t.Status.Beatmap.PlayMode))
}
