package handlers

import (
	"bytes"
	"fmt"
	"io"
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

func init() {
	StartTimeoutChecker()
}

// sendUserStatus sends the UserStatus
func sendUserStatus(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	x := constants.ClientSendUserStatusStruct{}
	helpers.UnmarshalBinary(r, &x)
	t.Status.Beatmap = x
	pckt.Write(public.SendUserStats(t, false))
}

// joinChannel sends a Join successfull/fail to the Client.
func joinChannel(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	yw := packets.NewWriter(t)
	xw := objects.ChannelInfo{}
	helpers.UnmarshalBinary(r, &xw)
	yw.JoinChannel(xw.ChannelName)
	pckt.Write(yw.Bytes())
}

// addFriend adds a Friend ofc!
func addFriend(r io.Reader, t *objects.Token) {
	i, err := helpers.RInt32(r)
	if err != nil {
		fmt.Println(err)
	}
	packets.AddFriend(&t.User, usertools.GetUser(int(i)))
}

// removeFriend removes a Friend ofc!
func removeFriend(r io.Reader, t *objects.Token) {
	i, err := helpers.RInt32(r)
	if err != nil {
		fmt.Println(err)
	}
	packets.RemoveFriend(&t.User, usertools.GetUser(int(i)))
}

func updateUserStats(r io.Reader, pckt *bytes.Buffer) {
	_, err := helpers.RIntArray(r)
	if err != nil {
		fmt.Println(err)
	}
	for y := 0; y < len(objects.TOKENS); y++ {
		pckt.Write(public.SendUserStats(objects.TOKENS[y], false))
	}
}

func sendUserPresence(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	i, err := helpers.RIntArray(r)
	yw := packets.NewWriter(t)
	if err != nil {
		fmt.Println(err)
	}
	for y := 0; y < len(i); y++ {
		yw.UserPresence(objects.GetTokenByID(i[y]))
	}
	pckt.Write(yw.Bytes())
}

func sendMessage(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	msg := constants.MessageStruct{}
	helpers.UnmarshalBinary(r, &msg)
	public.SendMessage(t, msg.Message, msg.Target)
}

func disconnectUser(t *objects.Token) {
	main := objects.GetStream("main")
	pckt := packets.NewPacket(constants.BanchoHandleUserQuit)
	pckt.SetPacketData(helpers.MarshalBinary(&constants.UserQuitStruct{UserID: t.User.ID, ErrorState: int8(0)}))
	main.Broadcast(pckt.ToByteArray(), nil)
	t = nil
}

// HandlePackets is the Main Packet handler.
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

		case constants.ClientSendUserStatus:
			sendUserStatus(r, pckt, t)

		case constants.ClientExit:
			disconnectUser(t)

		case constants.ClientSendIrcMessage:
			sendMessage(r, pckt, t)

		case constants.ClientSendIrcMessagePrivate:
			sendMessage(r, pckt, t)

		case constants.ClientRequestStatusUpdate:
			pckt.Write(public.SendUserStats(t, true))

		case constants.ClientPong:
			t.LastPing = time.Now()

		case constants.ClientReceiveUpdates:
			pckt.Write(public.SendUserStats(t, true))

		case constants.ClientChannelJoin:
			joinChannel(r, pckt, t)

		case constants.ClientFriendAdd:
			addFriend(r, t)

		case constants.ClientFriendRemove:
			removeFriend(r, t)

		case constants.ClientUserStatsRequest:
			updateUserStats(r, pckt)

		case constants.ClientUserPresenceRequest:
			sendUserPresence(r, pckt, t)

		default:
			logPacket(&pkg)

		}
	}
	w.Write(t.Output.Bytes())
	w.Write(pckt.Bytes())
	t.Output.Reset()
}

func logPacket(pkg *packets.Packet) {
	fmt.Println("---------Packet---------")
	fmt.Println("PacketID:", pkg.PacketID)
	fmt.Println("Length:", pkg.PacketLength)
	fmt.Println("PacketData:", pkg.PacketData)
	fmt.Println("------------------------")
}

func StartTimeoutChecker() {
	go func() {
		for {
			for i := 0; i < len(objects.TOKENS); i++ {
				if time.Time(objects.TOKENS[i].LastPing).Unix() < (time.Now().Unix() - int64(1000*10)) {
					disconnectUser(objects.TOKENS[i])
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()
}
