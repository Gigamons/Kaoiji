package public

import (
	"strings"

	"github.com/Gigamons/common/tools/usertools"
	"github.com/Mempler/osubinary"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
)

// SendMessage Send a Message from given client to given target
func SendMessage(t *objects.Token, Message string, Channel string) {
	main := objects.GetStream("main")
	p := packets.NewPacket(constants.BanchoSendMessage)
	msg := constants.MessageStruct{Message: Message, UserID: t.User.ID, Username: t.User.UserName, Target: Channel}

	if main == nil {
		return
	}
	if t == nil {
		return
	}
	if !strings.HasPrefix(Channel, "#") {
		targetid := usertools.GetUserID(Channel)
		targettoken := objects.GetTokenByID(int32(targetid))
		msg.Target = t.User.UserName
		p.SetPacketData(osubinary.Marshal(msg))
		targettoken.Write(p.ToByteArray())
		return
	}
	if objects.HasChannelPermission(Channel, t) {
		p.SetPacketData(osubinary.Marshal(msg))
		main.Broadcast(p.ToByteArray(), t)
	}
}
