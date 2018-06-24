package bot

import (
	"fmt"
	"strings"

	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/common/consts"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/common/helpers"
	"github.com/Mempler/osubinary"
)

type Command struct {
	CMD  string
	Prev int32
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
			objects.DeleteOldTokens(int32(usertools.GetUserID(args[1])))

		case "moderator":
			helpers.DB.Exec(sql, consts.AdminBanUsers+consts.AdminChatMod+consts.AdminKickUsers+consts.AdminSendAnnouncements+consts.AdminReports+consts.AdminSilenceUsers+consts.AdminWipeUsers, args[1])
			message(t, "Successfully gave", args[1], "the rank Moderator!")
			objects.DeleteOldTokens(int32(usertools.GetUserID(args[1])))

		case "bat":
			helpers.DB.Exec(sql, consts.AdminBeatmaps+consts.BAT, args[1])
			message(t, "Successfully gave", args[1], "the rank BAT!")
			objects.DeleteOldTokens(int32(usertools.GetUserID(args[1])))

		case "supporter":
			helpers.DB.Exec(sql, consts.Supporter, args[1])
			message(t, "Successfully gave", args[1], "the rank Supporter!")
			objects.DeleteOldTokens(int32(usertools.GetUserID(args[1])))

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
}

func init() {
	Commands = append(Commands, &Command{"!help", 0, []string{}, false, func(t *objects.Token, args ...string) {
		z := ""
		for i := 0; i < len(Commands); i++ {
			if helpers.HasPrivileges(int(Commands[i].Prev), t.User) {
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
