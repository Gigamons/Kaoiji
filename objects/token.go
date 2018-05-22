package objects

import (
	"bytes"
	"time"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/global"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/google/uuid"
)

// Token data
type Token struct {
	token  string
	User   consts.User
	Status struct {
		Torney  bool
		Beatmap constants.ClientSendUserStatusStruct
		Info    struct {
			Permissions int8
			ClientPerm  int8
			TimeZone    int8
			CountryID   int8
			Lon         float64
			Lat         float64
			Rank        int32
		}
	}
	Leaderboard consts.Leaderboard
	LastPing    time.Time
	Output      bytes.Buffer
}

// TOKENS Global Variable for Token array.
var TOKENS []*Token

// NewToken returns a Token that has a Token with a Token
func NewToken(uuid uuid.UUID, lon float64, lat float64, u consts.User) *Token {
	t := Token{}
	t.token = uuid.String()
	t.Status.Info.Lat = lat
	t.Status.Info.Lon = lon
	t.LastPing = time.Now()
	t.User = u

	t.Status.Info.ClientPerm |= constants.Userperm
	t.Status.Info.Permissions |= constants.Userperm

	if helpers.HasPrivileges(consts.BAT, u) {
		t.Status.Info.Permissions |= constants.Administrator
		t.Status.Info.ClientPerm |= constants.BAT
	}
	if global.CONFIG.Server.FreeDirect {
		t.Status.Info.ClientPerm |= constants.Supporter
	}
	if helpers.HasPrivileges(consts.Supporter, u) {
		t.Status.Info.Permissions |= constants.Supporter
	}
	if helpers.HasPrivileges(consts.AdminDeveloper, u) {
		t.Status.Info.Permissions |= constants.Developer
		t.Status.Info.ClientPerm |= constants.Developer
	}

	TOKENS = append(TOKENS, &t)

	return &t
}

// TokenExists return a boolean, true if exists else false
func TokenExists(token string) bool {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].token == token {
			return true
		}
	}
	return false
}

// GetToken Returns a UserToken from a String if not exists return nil
func GetToken(token string) *Token {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].token == token {
			return TOKENS[i]
		}
	}
	return nil
}

// GetTokenByID Returns a UserToken from an UserID if not exists return nil
func GetTokenByID(userid int32) *Token {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].User.ID == userid {
			return TOKENS[i]
		}
	}
	return nil
}

// StartTimeoutChecker https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang/16466581
func StartTimeoutChecker() {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				for i := 0; i < len(TOKENS); i++ {
					if time.Time(TOKENS[i].LastPing).Unix() < (time.Now().Unix() - int64(1000*30)) {
						TOKENS = append(TOKENS[:i], TOKENS[i+1:]...)
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
