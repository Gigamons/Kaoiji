package objects

import (
	"bytes"
	"time"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/constants/packets"
	"github.com/Gigamons/Kaoiji/constants/privileges"
	"github.com/Gigamons/Kaoiji/global"
	"github.com/Gigamons/Kaoiji/helpers"
	"github.com/google/uuid"
)

// Token data
type Token struct {
	token  string
	User   constants.User
	Status struct {
		Torney  bool
		Beatmap packetconst.ClientSendUserStatus
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
	Leaderboard constants.Leaderboard
	LastPing    time.Time
	Output      bytes.Buffer
}

var TOKENS []Token

// NewToken returns a Token that has a Token with a Token
func NewToken(uuid uuid.UUID, lon float64, lat float64, u constants.User) Token {
	t := Token{}
	t.token = uuid.String()
	t.Status.Info.Lat = lat
	t.Status.Info.Lon = lon
	t.LastPing = time.Now()
	t.User = u

	t.Status.Info.ClientPerm |= constants.Userperm
	t.Status.Info.Permissions |= constants.Userperm

	if helpers.HasPrivileges(privileges.BAT, u) {
		t.Status.Info.Permissions |= constants.Administrator
		t.Status.Info.ClientPerm |= constants.BAT
	}
	if global.CONFIG.Server.FreeDirect {
		t.Status.Info.ClientPerm |= constants.Supporter
	}
	if helpers.HasPrivileges(privileges.Supporter, u) {
		t.Status.Info.Permissions |= constants.Supporter
	}
	if helpers.HasPrivileges(privileges.AdminDeveloper, u) {
		t.Status.Info.Permissions |= constants.Developer
		t.Status.Info.ClientPerm |= constants.Developer
	}

	TOKENS = append(TOKENS, t)

	return t
}

func TokenExists(token string) bool {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].token == token {
			return true
		}
	}
	return false
}

func GetToken(token string) *Token {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].token == token {
			return &TOKENS[i]
		}
	}
	return nil
}

func GetTokenByID(userid int32) *Token {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].User.ID == userid {
			return &TOKENS[i]
		}
	}
	return nil
}

// https://stackoverflow.com/questions/16466320/is-there-a-way-to-do-repetitive-tasks-at-intervals-in-golang/16466581
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
