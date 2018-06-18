package handlers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/Gigamons/Kaoiji/handlers/public"
	"github.com/Mempler/osubinary"

	"github.com/google/uuid"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/Kaoiji/packets"
)

func Err(w http.ResponseWriter) {
	w.Header().Add("cho-token", "error")
	p := constants.NewPacket(constants.BanchoLoginReply)
	p.SetPacketData(osubinary.Int32(constants.LoginException))
	w.Write(p.ToByteArray())
}

// LoginHandler main Login Handler to Handle logins... Makes sense!
func LoginHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("------------ERROR------------")
			fmt.Println(err)
			fmt.Println("---------ERROR TRACE---------")
			fmt.Println(string(debug.Stack()))
			fmt.Println("----------END ERROR----------")
			Err(w)
		}
	}()

	uuid, err := uuid.NewRandom()
	if err != nil {
		logger.Errorln(err)
		Err(w)
		return
	}

	w.Header().Add("cho-token", uuid.String())
	b, err := ioutil.ReadAll(r.Body)
	pw := packets.NewWriter(&objects.Token{})
	if err != nil {
		logger.Errorln(err)
		Err(w)
	}

	l, err := _parseLoginData(b)
	if err != nil {
		Err(w)
	}
	userid := usertools.GetUserID(l.Username)
	if userid < 1 {
		logger.Infoln(l.Username, "Failed to Login!")
		pw.UserID(constants.LoginFailed)
		w.Write(pw.Bytes())
		return
	}
	u := usertools.GetUser(userid)
	if !u.CheckPassword(l.Password) {
		logger.Infoln(l.Username, "Failed to Login!")
		pw.UserID(constants.LoginFailed)
		w.Write(pw.Bytes())
		return
	}

	objects.DeleteOldTokens(u.ID)

	main := objects.GetStream("main")

	if main == nil {
		logger.Errorln("Nil exception, main = nil")
		Err(w)
		return
	}

	var use string

	if r.Header.Get("X-Forwarded-For") != "127.0.0.1" || r.Header.Get("X-Forwarded-For") != "0.0.0.0" || r.Header.Get("X-Forwarded-For") != "" {
		use = r.Header.Get("X-Forwarded-For")
	}

	if r.Header.Get("X-Real-IP") != "127.0.0.1" || r.Header.Get("X-Real-IP") != "0.0.0.0" || r.Header.Get("X-Real-IP") != "" {
		use = r.Header.Get("X-Real-IP")
	}

	info := helpers.GetIPInfo(use)
	if info == nil {
		info = &consts.GeoIP{}
	}

	t := objects.NewToken(uuid, info.Location.Lon, info.Location.Lat, u)
	t.Status.Info.CountryID = consts.ToCountryID(info.Country)

	main.AddUser(t)

	pw.SetToken(t)

	pw.ProtocolVersion(19)
	pw.UserID(int32(t.User.ID))
	pw.UserPresence(t)
	pw.SendFriendlist()
	pw.PresenceBundle()
	pw.LoginPermissions()
	pw.AutoJoinChannel()
	pw.ChannelAvaible()
	pw.Write(public.SendUserStats(t, true))
	pasw := packets.NewWriter(t)
	pasw.Write(public.SendUserStats(t, false))
	pasw.PresenceSingle(t.User.ID)
	main.Broadcast(pasw.Bytes(), nil)

	logger.Infoln(l.Username, "has logged in!")

	w.Write(pw.Bytes())
}

type loginData struct {
	Username         string
	Password         string
	OsuVersion       string
	TimeOffset       int
	BlockNonFriendDM bool
	outDated         bool
	SecurityHash     securityHash
}

type securityHash struct {
	OsuHash   string
	DiskMD5   string
	UniqueMD5 string
}

func _parseLoginData(b []byte) (loginData, error) {
	s := string(b)
	sa := strings.Split(s, "\n")
	if len(sa) < 2 {
		return loginData{}, errors.New("no")
	}
	x := strings.Split(sa[2], "|")
	if len(x) < 4 {
		return loginData{}, errors.New("no")
	}
	y := loginData{Username: sa[0], Password: sa[1], OsuVersion: x[0]}
	timeOffset, err := strconv.Atoi(x[1])
	if err != nil {
		panic(err)
	}
	y.TimeOffset = timeOffset

	sec := strings.Split(x[3], ":")
	if len(sec) < 4 {
		return loginData{}, errors.New("no")
	}
	BlockNonFriendDM := len(x) >= 4 && x[4] == "1"
	y.BlockNonFriendDM = BlockNonFriendDM

	y.outDated = len(sec) < 4

	y.SecurityHash.DiskMD5 = sec[4]
	y.SecurityHash.UniqueMD5 = sec[5]

	return y, nil
}
