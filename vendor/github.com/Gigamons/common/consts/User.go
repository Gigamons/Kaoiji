package consts

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User is a User Object
type User struct {
	ID             int32
	UserName       string
	UserNameSafe   string
	EMail          string
	BCryptPassword string
	Privileges     int32
	Status         Status
	Relax          bool
}

// Status status of current User
type Status struct {
	Country        int16
	Lat            float64
	Lon            float64
	Banned         bool
	BannedUntil    time.Time
	BannedReason   string
	Silenced       bool
	SilencedUntil  time.Time
	SilencedReason string
	Verified       bool
}

type Leaderboard struct {
	RankedScore int64
	TotalScore  int64
	Count300    int64
	Count100    int64
	Count50     int64
	CountMiss   int64
	Playcount   int64
	PeppyPoints int16
	Position    int32
}

// CheckPassword if valid.
func (u *User) CheckPassword(passwordMD5 string) bool {
	if bcrypt.CompareHashAndPassword([]byte(u.BCryptPassword), []byte(passwordMD5)) != nil {
		return false
	}
	return true
}
