package public

import (
	"strings"
	"time"

	"github.com/Gigamons/common/tools/usertools"
	"github.com/Mempler/osubinary"

	"github.com/Gigamons/Kaoiji/bot"
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
)

// SendMessage Send a Message from given client to given target
func SendMessage(t *objects.Token, Message string, Channel string) {
	main := objects.GetStream("main")
	p := constants.NewPacket(constants.BanchoSendMessage)
	msg := constants.MessageStruct{Message: Message, UserID: t.User.ID, Username: t.User.UserName, Target: Channel}

	if main == nil || t == nil {
		return
	}
	x := strings.Split(Message, " ")
	if len(x) < 1 {
		return
	}
	cmd := bot.GetCommand(x[0])
	if cmd != nil && (t.User.Privileges&cmd.Prev != 0 || cmd.Prev == 0) && cmd.CMD != "!faq" {
		if len(cmd.Args) > (len(x) - 1) {
			cmd.Help(t)
			return
		}
		cmd.Func(t, x...)
		return
	} else if cmd != nil && cmd.CMD == "!faq" {
		if len(cmd.Args) > (len(x) - 1) {
			cmd.Help(t)
			return
		}
		cmd.Func(t, x...)
	}
	if t.User.Status.SilencedUntil.Unix() > time.Now().Unix() {
		return
	}
	if !strings.HasPrefix(Channel, "#") {
		targetid := usertools.GetUserID(Channel)
		targettoken := objects.GetTokenByID(*targetid)
		if targettoken == nil {
			return
		}
		msg.Target = t.User.UserName
		p.SetPacketData(osubinary.Marshal(msg))
		go targettoken.Write(p.ToByteArray())
		if targettoken.User.Status.SilencedUntil.Unix() > time.Now().Unix() {
			pcktsilenced := constants.NewPacket(constants.BanchoTargetIsSilenced)
			msg.Target = targettoken.User.UserName
			pcktsilenced.SetPacketData(osubinary.Marshal(msg))
			go t.Write(pcktsilenced.ToByteArray())
		}
		return
	}
	if objects.HasChannelPermission(Channel, t) {
		p.SetPacketData(osubinary.Marshal(msg))
		if strings.HasSuffix(Channel, "spectator") {
			go t.SpectatorStream.BroadcastRaw(p.ToByteArray(), false, t, false)
			return
		}
		go main.Broadcast(p.ToByteArray(), t)
	}
}
