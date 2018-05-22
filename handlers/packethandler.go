package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/Kaoiji/constants"

	"github.com/Gigamons/Kaoiji/handlers/public"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
)

func HandlePackets(w http.ResponseWriter, r *http.Request, t *objects.Token) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	pkgs := packets.GetPackets(b)
	pckt := bytes.NewBuffer([]byte{})
	for i := 0; i < len(pkgs); i++ {
		pkg := pkgs[i]
		r := bytes.NewReader(pkg.PacketData)
		switch pkg.PacketID {
		case constants.ClientSendUserStatus: // 0
			x := constants.ClientSendUserStatusStruct{}
			helpers.UnmarshalBinary(r, &x)
			t.Status.Beatmap = x
			pckt.Write(public.SendUserStats(t, false))
		case constants.ClientRequestStatusUpdate: // 3
			pckt.Write(public.SendUserStats(t, true))
		case constants.ClientPong: // 4
			t.LastPing = time.Now()
		case constants.ClientReceiveUpdates:
			pckt.Write(public.SendUserStats(t, true))
		case constants.ClientChannelJoin:
			yw := packets.NewWriter(t)
			xw := objects.ChannelInfo{}
			helpers.UnmarshalBinary(r, &xw)
			yw.JoinChannel(xw.ChannelName)
			pckt.Write(yw.Bytes())
		case 68:
			//fmt.Println(packets.ReadBeatmaps(r))

		case constants.ClientFriendAdd:
			i, err := helpers.RInt32(r)
			if err != nil {
				fmt.Println(err)
			}
			packets.AddFriend(&t.User, usertools.GetUser(int(i)))

		case constants.ClientFriendRemove:
			i, err := helpers.RInt32(r)
			if err != nil {
				fmt.Println(err)
			}
			packets.RemoveFriend(&t.User, usertools.GetUser(int(i)))

		case constants.ClientUserStatsRequest: // 84
			i, err := helpers.RIntArray(r)
			if err != nil {
				fmt.Println(err)
			}
			for y := 0; y < len(i); y++ {
				pckt.Write(public.SendUserStats(objects.GetTokenByID(i[y]), false))
			}
		case constants.ClientUserPresenceRequest: // 97
			i, err := helpers.RIntArray(r)
			yw := packets.NewWriter(t)
			if err != nil {
				fmt.Println(err)
			}
			for y := 0; y < len(i); y++ {
				yw.UserPresence(objects.GetTokenByID(i[y]))
			}
			pckt.Write(yw.Bytes())
		default:
			fmt.Println("---------Packet---------")
			fmt.Println("PacketID:", pkg.PacketID)
			fmt.Println("Length:", pkg.PacketLength)
			fmt.Println("PacketData:", pkg.PacketData)
			fmt.Println("------------------------")
		}
	}
	w.Write(t.Output.Bytes())
	t.Output.Reset()
	w.Write(pckt.Bytes())
}
