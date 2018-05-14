package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"git.gigamons.de/Gigamons/Kaoiji/tools/usertools"

	"git.gigamons.de/Gigamons/Kaoiji/constants"
	"git.gigamons.de/Gigamons/Kaoiji/constants/packets"

	"git.gigamons.de/Gigamons/Kaoiji/handlers/public"
	"git.gigamons.de/Gigamons/Kaoiji/objects"
	"git.gigamons.de/Gigamons/Kaoiji/packets"
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
			x := packetconst.ClientSendUserStatus{}
			packets.UnmarshalBinary(r, &x)
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
			packets.UnmarshalBinary(r, &xw)
			yw.JoinChannel(xw.ChannelName)
			pckt.Write(yw.Bytes())
		case 68:
			//fmt.Println(packets.ReadBeatmaps(r))

		case constants.ClientFriendAdd:
			i, err := packets.RInt32(r)
			if err != nil {
				fmt.Println(err)
			}
			packets.AddFriend(&t.User, usertools.GetUser(int(i)))

		case constants.ClientFriendRemove:
			i, err := packets.RInt32(r)
			if err != nil {
				fmt.Println(err)
			}
			packets.RemoveFriend(&t.User, usertools.GetUser(int(i)))

		case constants.ClientUserStatsRequest: // 84
			i, err := packets.RIntArray(r)
			if err != nil {
				fmt.Println(err)
			}
			for y := 0; y < len(i); y++ {
				pckt.Write(public.SendUserStats(objects.GetTokenByID(i[y]), false))
			}
		case constants.ClientUserPresenceRequest: // 97
			i, err := packets.RIntArray(r)
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
