package bot

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/common/consts"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/helpers"
	"github.com/Mempler/osubinary"
)

type Command struct {
	CMD  string
	Prev uint64
	Args []string
	Pub  bool
	Func func(*objects.Token, ...string)
}

var Commands = []*Command{
	&Command{"!setperm", consts.AdminPrivileges, []string{"Username(Safe)", "Rank"}, false, func(t *objects.Token, args ...string) {
		sql := "UPDATE users SET privileges=? WHERE username_safe = ?"
		switch strings.ToLower(args[2]) {
		case "developer":
			helpers.DB.Exec(sql, consts.AdminBanUsers+consts.AdminBeatmaps+consts.AdminChatMod+consts.AdminDeveloper+consts.AdminKickUsers+consts.AdminManageUsers+consts.AdminPannelAccess+consts.AdminPrivileges+consts.AdminReports+consts.AdminSendAnnouncements+consts.AdminSettings+consts.AdminSilenceUsers+consts.AdminWipeUsers+consts.BAT+consts.Supporter, args[1])
			message(t, "Successfully gave", args[1], "the rank Developer!")
			objects.DeleteOldTokens(*usertools.GetUserID(args[1]))

		case "moderator":
			helpers.DB.Exec(sql, consts.AdminBanUsers+consts.AdminChatMod+consts.AdminKickUsers+consts.AdminSendAnnouncements+consts.AdminReports+consts.AdminSilenceUsers+consts.AdminWipeUsers, args[1])
			message(t, "Successfully gave", args[1], "the rank Moderator!")
			objects.DeleteOldTokens(*usertools.GetUserID(args[1]))

		case "bat":
			helpers.DB.Exec(sql, consts.AdminBeatmaps+consts.BAT, args[1])
			message(t, "Successfully gave", args[1], "the rank BAT!")
			objects.DeleteOldTokens(*usertools.GetUserID(args[1]))

		case "supporter":
			helpers.DB.Exec(sql, consts.Supporter, args[1])
			message(t, "Successfully gave", args[1], "the rank Supporter!")
			objects.DeleteOldTokens(*usertools.GetUserID(args[1]))

		default:
			message(t, "Unknown Rank!")
		}
	}},
	&Command{"!announce", consts.AdminSendAnnouncements, []string{"Message"}, false, func(t *objects.Token, args ...string) {
		var xf string
		for i := 1; i < len(args); i++ {
			xf += args[i] + " "
		}
		xf = strings.Trim(xf, " ")
		announce(nil, xf)
	}},
	&Command{"!clear", consts.AdminChatMod, []string{}, false, func(t *objects.Token, args ...string) {
		pw := []byte{}
		for _, tok := range objects.TOKENS {
			pw = append(pw, deletemessages(tok.User.ID)...)
		}
		main := objects.GetStream("main")
		main.Broadcast(pw, nil)
		message(t, "Chat cleared!")
	}},
	/* TMP COMMAND */
	&Command{"!adduser", consts.AdminDeveloper, []string{"Username", "Password"}, false, func(t *objects.Token, args ...string) {
		u := args[1]
		p := args[2]
		if p == "random" {
			p = ""
			run := []rune("abcdefghijklmnopqrstuvwxyz123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
			for i := 0; i < 16; i++ {
				rand.Seed(time.Now().UnixNano())
				p += string(run[rand.Intn(len(run))])
			}
		}
		hash, err := helpers.MD5String(p)
		if err != nil {
			message(t, "could not create useraccount!", err.Error())
			return
		}
		phash, err := bcrypt.GenerateFromPassword([]byte(hex.EncodeToString(hash)), -10)
		if err != nil {
			message(t, "could not create useraccount!", err.Error())
			return
		}
		helpers.DB.Exec("INSERT INTO users (username, username_safe, password) VALUES (?, ?, ?) ", u, strings.ToLower(strings.Replace(u, " ", "_", -1)), string(phash))
		helpers.DB.Exec("INSERT INTO users_status () VALUES ()")
		helpers.DB.Exec("INSERT INTO leaderboard () VALUES ()")
		helpers.DB.Exec("INSERT INTO leaderboard_rx () VALUES ()")
		message(t, "Created user with the username of", u, "and the password of\nPass:", p, "\nMD5:", hex.EncodeToString(hash), "\nPassHash:", string(phash))
	}},
	&Command{"!silence", consts.AdminChatMod, []string{"Username", "date"}, false, func(t *objects.Token, Args ...string) {
		target := usertools.GetUser(usertools.GetUserID(Args[1]))
		if target == nil || target.ID == 0 {
			return
		}
		Args = append(Args[:0], Args[2:]...)
		var outms int64 = 0
		pim, pum, err := parsetimingstring(Args[0])
		if err != nil {
			message(t, err.Error())
			return
		}
		for xz, letter := range pim {
			tim := pum[xz]
			switch letter {
			case "s":
				outms += int64(tim)
			case "m":
				outms += int64(tim) * 60
			case "h":
				outms += int64(tim) * 60 * 60
			case "d":
				outms += int64(tim) * 60 * 60 * 24
			case "w":
				outms += int64(tim) * 60 * 60 * 24 * 7
			case "mo":
				outms += int64(tim) * 60 * 60 * 24 * 31
			case "y":
				outms += int64(tim)*60*60*24*31*12 - 60*60*24*7
			default:
				message(t, "Unknown format", strconv.Itoa(tim), string(letter))
			}
		}
		x := time.Unix(outms+time.Now().Unix(), 0)
		helpers.DB.Exec("UPDATE users_status SET silenced_until=? WHERE id=?", x, target.ID)
		helpers.DB.Exec("UPDATE users_status SET silenced_reason=? WHERE id=?", strings.Join(Args[1:], " "), target.ID)

		message(t, "Silenced", target.UserName, "for", strconv.Itoa(int(outms)), "second(s)!")
		targetToken := objects.GetTokenByID(target.ID)
		if targetToken == nil {
			return
		}
		message(targetToken, "You've got silence due this reason:", strings.Join(Args[1:], " "))
		silence(targetToken, int32(outms)) // ppfff overflow but WHO CARES ?!?
		t.User = usertools.GetUser(&target.ID)
		t.Leaderboard = usertools.GetLeaderboard(t.User, t.Status.Beatmap.PlayMode)
	}},
	&Command{"!unsilence", consts.AdminChatMod, []string{"Username"}, false, func(t *objects.Token, Args ...string) {
		target := usertools.GetUser(usertools.GetUserID(Args[1]))
		if target == nil {
			return
		}
		helpers.DB.Exec("UPDATE users_status SET silenced_until=? WHERE id=?", time.Now(), target.ID)
		message(t, "Unsilenced", target.UserName)
		targetToken := objects.GetTokenByID(target.ID)
		if targetToken == nil {
			return
		}
		silence(targetToken, int32(0)) // ppfff overflow but WHO CARES ?!?
		t.User = usertools.GetUser(&target.ID)
		t.Leaderboard = usertools.GetLeaderboard(t.User, t.Status.Beatmap.PlayMode)
	}},
	&Command{"!rtx", consts.AdminDeveloper, []string{"Username", "...reason"}, false, func(t *objects.Token, Args ...string) {
		target := usertools.GetUser(usertools.GetUserID(Args[1]))
		if target == nil || target.ID == 0 {
			return
		}
		targetToken := objects.GetTokenByID(target.ID)
		if targetToken == nil {
			return
		}
		Reason := strings.Join(append(Args[:0], Args[2:]...), " ")
		p := constants.NewPacket(constants.BanchoRTX)
		p.SetPacketData(osubinary.BString(Reason))
		targetToken.Write(p.ToByteArray())
	}},
	&Command{"!kill", consts.AdminDeveloper, []string{"Username"}, false, func(t *objects.Token, Args ...string) {
		target := usertools.GetUser(usertools.GetUserID(Args[1]))
		if target == nil || target.ID == 0 {
			return
		}
		targetToken := objects.GetTokenByID(target.ID)
		if targetToken == nil {
			return
		}
		p := constants.NewPacket(constants.BanchoLoginPermissions)
		p.SetPacketData(osubinary.Int32(int32(constants.BAT)))
		targetToken.Write(p.ToByteArray())
		p = constants.NewPacket(constants.BanchoPing)
		targetToken.Write(p.ToByteArray())
	}},
}

func init() {
	Commands = append(Commands, &Command{"!help", 0, []string{}, false, func(t *objects.Token, args ...string) {
		z := ""
		for i := 0; i < len(Commands); i++ {
			if helpers.HasPrivileges(Commands[i].Prev, t.User) {
				z += "Command: " + Commands[i].CMD + x(SToIN(Commands[i].Args...)...) + "\n"
			}
		}
		message(t, z)
	}})
}

func GetCommand(cmd string) *Command {
	for i := 0; i < len(Commands); i++ {
		if Commands[i].CMD == cmd {
			return Commands[i]
		}
	}
	return nil
}

func (cmd *Command) Help(t *objects.Token) {
	message(t, "Usage: "+cmd.CMD+""+x(SToIN(cmd.Args...)...))
}

func announce(t *objects.Token, msg ...string) {
	p := constants.NewPacket(constants.BanchoAnnounce)
	p.SetPacketData(osubinary.BString(fmt.Sprintln(SToIN(msg...)...)))
	if t == nil {
		m := objects.GetStream("main")
		m.Broadcast(p.ToByteArray(), t)
		return
	}
	t.Write(p.ToByteArray())
}

func message(t *objects.Token, msg ...string) {
	if t == nil {
		return
	}
	p := constants.NewPacket(constants.BanchoSendMessage)
	p.SetPacketData(osubinary.Marshal(constants.MessageStruct{"GigaBot", fmt.Sprintln(SToIN(msg...)...), "GigaBot", 100}))
	t.Write(p.ToByteArray())
}

func parsetimingstring(s string) ([]string, []int, error) {
	timingrgx := regexp.MustCompile(`(\d+[smhdy][mo]?)`)
	rgx := regexp.MustCompile(`([smhdy][mo]?)`)
	if timingrgx == nil || rgx == nil {
		return nil, nil, errors.New("Could not parse timing")
	}
	var x []int
	orx := rgx.Split(strings.Join(timingrgx.FindAllString(s, -1), ""), -1)
	srx := rgx.FindAllString(strings.Join(timingrgx.FindAllString(s, -1), ""), -1)
	for i := 0; i < len(orx); i++ {
		if orx[i] == "" {
			break
		}
		y, err := strconv.Atoi(orx[i])
		if err != nil {
			return nil, nil, err
		}
		x = append(x, y)
	}
	if len(srx) != len(x) {
		return nil, nil, errors.New("Could not parse timing")
	}
	return srx, x, nil
}

func deletemessages(userid uint32) []byte {
	p := constants.NewPacket(constants.BanchoUserSilenced)
	p.SetPacketData(osubinary.UInt32(userid))
	return p.ToByteArray()
}

func silence(t *objects.Token, timeout int32) {
	if timeout < 0 {
		timeout = 0
	}
	p := constants.NewPacket(constants.BanchoBanInfo)
	p.SetPacketData(osubinary.Int32(timeout))
	t.Write(p.ToByteArray())
}

func SToIN(s ...string) []interface{} {
	x := []interface{}{}
	for i := 0; i < len(s); i++ {
		x = append(x, interface{}(s[i]))
	}
	return x
}

func x(x ...interface{}) string {
	o := ""
	for i := 0; i < len(x); i++ {
		o += " <" + x[i].(string) + ">"
	}
	return o
}
