package objects

import (
	"bytes"
	"sync"
	"time"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/global"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/google/uuid"
)

var lockPackets = &sync.Mutex{}
var lockAppend = &sync.Mutex{}

// Token data
type Token struct {
	Token  string
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
	AlreadyNotified bool
	SpectatorStream *SpectatorStream
	Leaderboard     consts.Leaderboard
	LastPing        time.Time
	Output          bytes.Buffer
}

// TOKENS Global Variable for Token array.
var TOKENS []*Token

// NewToken returns a Token that has a Token with a Token
func NewToken(uuid uuid.UUID, lon float64, lat float64, u consts.User) *Token {
	lockAppend.Lock()
	t := &Token{}
	t.Token = uuid.String()
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

	t.SpectatorStream = NewSpectatorStream(t)
	TOKENS = append(TOKENS, t)
	lockAppend.Unlock()
	return t
}

// DeleteToken deletes the given Token (String) from our TOKENS Array.
func DeleteToken(token string) {
	lockAppend.Lock()
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].Token == token {
			copy(TOKENS[i:], TOKENS[i+1:])
			TOKENS[len(TOKENS)-1] = nil
			TOKENS = TOKENS[:len(TOKENS)-1]
		}
	}
	lockAppend.Unlock()
}

func DeleteOldTokens(userid int32) {
	lockAppend.Lock()
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].User.ID == userid {
			copy(TOKENS[i:], TOKENS[i+1:])
			TOKENS[len(TOKENS)-1] = nil
			TOKENS = TOKENS[:len(TOKENS)-1]
		}
	}
	lockAppend.Unlock()
}

// Write writes to our Client that'll get send to client on Next/This request.
func (t *Token) Write(f []byte) {
	lockPackets.Lock()
	t.Output.Write(f)
	lockPackets.Unlock()
}

// TokenExists return a boolean, true if exists else false
func TokenExists(token string) bool {
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].Token == token {
			return true
		}
	}
	return false
}

// GetToken Returns a UserToken from a String if not exists return nil
func GetToken(token string) *Token {
	lockAppend.Lock()
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].Token == token {
			lockAppend.Unlock()
			return TOKENS[i]
		}
	}
	lockAppend.Unlock()
	return nil
}

// GetTokenByID Returns a UserToken from an UserID if not exists return nil
func GetTokenByID(userid int32) *Token {
	lockAppend.Lock()
	for i := 0; i < len(TOKENS); i++ {
		if TOKENS[i].User.ID == userid {
			lockAppend.Unlock()
			return TOKENS[i]
		}
	}
	lockAppend.Unlock()
	return nil
}
