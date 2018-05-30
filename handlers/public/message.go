package public

import (
	"fmt"
	"strings"

	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
	"github.com/Gigamons/common/helpers"
)

func SendMessage(t *objects.Token, Message string, Channel string) {
	main := objects.GetStream("main")
	p := packets.NewPacket(constants.BanchoSendMessage)
	msg := &constants.MessageStruct{Message: Message, UserID: t.User.ID, Username: t.User.UserName, Target: Channel}

	if msg == nil {
		fmt.Println("Msg = nil")
		return
	}
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
		p.SetPacketData(helpers.MarshalBinary(msg))
		targettoken.Write(p.ToByteArray())
		return
	}
	if objects.HasChannelPermission(Channel, t) {
		p.SetPacketData(helpers.MarshalBinary(msg))
		main.Broadcast(p.ToByteArray(), t)
	}
}
