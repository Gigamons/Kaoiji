package usertools

import (
	"fmt"
	"log"
	"strings"
	"time"

	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"git.gigamons.de/Gigamons/Kaoiji/global"
)

// GetUserID Returns a UserID of the given UserName
func GetUserID(username string) int {
	db := global.DB
	safe := strings.ToLower(strings.Replace(username, " ", "_", -1))
	UserID := 0
	err := db.Ping()
	if err != nil {
		fmt.Println("DB")
		return -1
	}
	rows, err := db.Query("SELECT id FROM users WHERE username_safe = ?", safe)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
		return -1
	}
	for rows.Next() {
		err := rows.Scan(&UserID)
		if err != nil {
			log.Fatal(err)
			return -1
		}
	}
	return UserID
}

// GetUser returns a User based on the Defined UserID
func GetUser(userid int) *constants.User {
	db := global.DB
	u := constants.User{}

	rows, err := db.Query("SELECT id, username, username_safe, email, password, privileges FROM users WHERE id = ?", userid)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&u.ID, &u.UserName, &u.UserNameSafe, &u.EMail, &u.BCryptPassword, &u.Privileges)
		if err != nil {
			log.Fatal(err)
		}
	}

	u.Status = getUserStatus(userid)

	return &u
}

func getUserStatus(userid int) constants.Status {
	db := global.DB
	u := constants.Status{}
	banneduntil := ""
	silenceduntil := ""
	rows, err := db.Query("SELECT country, lat, lon, banned, banned_until, banned_reason, silenced, silenced_until, silenced_reason, verified FROM users_status WHERE id = ?", userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&u.Country, &u.Lat, &u.Lon, &u.Banned, &banneduntil, &u.BannedReason, &u.Silenced, &silenceduntil, &u.SilencedReason, &u.Verified)
		if err != nil {
			log.Fatal(err)
		}
	}

	buntil := time.Now()
	buntil.Format(banneduntil)

	u.BannedUntil = buntil

	suntil := time.Now()
	suntil.Format(silenceduntil)

	u.SilencedUntil = suntil

	return u
}

func GetLeaderboard(u constants.User, playMode int8) constants.Leaderboard {
	var db = global.DB
	var m string
	if u.Relax {
		m = "_rx"
	} else {
		m = ""
	}
	var lb constants.Leaderboard
	pm := constants.ToPlaymodeString(playMode)
	if pm != "" {
		pm = "_" + pm
	}
	rows, err := db.Query("SELECT rankedscore"+pm+", totalscore"+pm+", count_300"+pm+", count_100"+pm+", count_50"+pm+", count_miss"+pm+", playcount"+pm+", pp"+pm+" FROM leaderboard"+m+" WHERE id = ?", u.ID)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		err := rows.Scan(&lb.RankedScore, &lb.TotalScore, &lb.Count300, &lb.Count100, &lb.Count50, &lb.CountMiss, &lb.Playcount, &lb.PeppyPoints)
		if err != nil {
			log.Fatal(err)
		}
	}

	return lb
}

func GetLeaderboardPosition(u constants.User, playMode int8) int32 {
	var userids []int
	db := global.DB
	var m string
	if u.Relax {
		m = "_rx"
	} else {
		m = ""
	}
	pm := constants.ToPlaymodeString(playMode)
	if pm != "" {
		pm = "_" + pm
	}
	rows, err := db.Query("SELECT id FROM leaderboard" + m + " ORDER BY pp" + pm + " DESC, rankedscore" + pm + " DESC")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var useridstr int
		err := rows.Scan(&useridstr)
		userids = append(userids, useridstr)
		if err != nil {
			log.Fatal(err)
		}
	}

	for i := 0; i < len(userids); i++ {
		if int32(userids[i]) == u.ID {
			return int32(i + 1)
		}
	}

	return 0
}
