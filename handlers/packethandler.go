package handlers

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Gigamons/common/logger"
	"github.com/Mempler/osubinary"

	"github.com/Gigamons/common/tools/usertools"

	"github.com/Gigamons/Kaoiji/constants"

	"github.com/Gigamons/Kaoiji/handlers/public"
	"github.com/Gigamons/Kaoiji/objects"
	"github.com/Gigamons/Kaoiji/packets"
)

func init() {
	// Start the Timeout Check, since we don't know if the user has a timeout, we do that manually. (Only call it once)
	StartTimeoutChecker()
}

// sendUserStatus sends the UserStatus
func sendUserStatus(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	x := constants.ClientSendUserStatusStruct{}
	osubinary.Unmarshal(r, &x)
	t.Status.Beatmap = x
	if x.CurrentMods&128 > 0 || x.CurrentMods&8192 > 0 {
		if !t.AlreadyNotified {
			t.AlreadyNotified = true
			w := packets.NewWriter(t)
			w.Announce("You've enabled the RX/AP Scoreboard.\nAP has Nerfed aim PP.\nRX has Nerfed speed PP.")
			t.Write(w.Bytes())
		}
		if !t.User.Relax {
			pckt.Write(public.SendUserStats(t, true))
		} else {
			pckt.Write(public.SendUserStats(t, false))
		}
		t.User.Relax = true
	} else {
		if t.User.Relax {
			pckt.Write(public.SendUserStats(t, true))
		} else {
			pckt.Write(public.SendUserStats(t, false))
		}
		t.User.Relax = false
	}
}

// joinChannel sends a Join successfull/fail to the Client.
func joinChannel(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	yw := packets.NewWriter(t)
	xw := objects.ChannelInfo{}
	osubinary.Unmarshal(r, &xw)
	yw.JoinChannel(xw.ChannelName)
	yw.ChannelAvaible()
	pckt.Write(yw.Bytes())
}

// addFriend adds a Friend ofc!
func addFriend(r io.Reader, t *objects.Token) {
	i, err := osubinary.RInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	packets.AddFriend(t.User, usertools.GetUser(int(i)))
}

// removeFriend removes a Friend ofc!
func removeFriend(r io.Reader, t *objects.Token) {
	i, err := osubinary.RInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	packets.RemoveFriend(t.User, usertools.GetUser(int(i)))
}

// updateUserStats updates the User stats for that user, (No fetching out of SQL)
func updateUserStats(r io.Reader, pckt *bytes.Buffer) {
	_, err := osubinary.RIntArray(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	for y := 0; y < len(objects.TOKENS); y++ {
		pckt.Write(public.SendUserStats(objects.TOKENS[y], false))
	}
}

// sendUserPresence Send the User Precense of the Given UserID.
func sendUserPresence(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	i, err := osubinary.RIntArray(r)
	yw := packets.NewWriter(t)
	if err != nil {
		logger.Errorln(err)
		return
	}
	for y := 0; y < len(i); y++ {
		yw.UserPresence(objects.GetTokenByID(i[y]))
	}
	pckt.Write(yw.Bytes())
}

// sendMessage User sends a MSG Packet to us, we're handling it and send it to the User Target
func sendMessage(r io.Reader, pckt *bytes.Buffer, t *objects.Token) {
	msg := constants.MessageStruct{}
	osubinary.Unmarshal(r, &msg)
	public.SendMessage(t, msg.Message, msg.Target)
}

// disconnectUser Our user has a Timeout nor He Disconnects, we're broadcast it to Everyone that that user got a Timeout / Disconnect.
func disconnectUser(t *objects.Token) {
	main := objects.GetStream("main")
	pckt := constants.NewPacket(constants.BanchoHandleUserQuit)
	pckt.SetPacketData(osubinary.Marshal(constants.UserQuitStruct{UserID: t.User.ID, ErrorState: int8(0)}))
	objects.DeleteToken(t.Token)
	main.Broadcast(pckt.ToByteArray(), nil)
}

func startSpectate(r io.Reader, t *objects.Token) {
	i, err := osubinary.RInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	HostToken := objects.GetTokenByID(i)
	if HostToken.SpectatorStream == nil {
		objects.NewSpectatorStream(HostToken)
	}
	HostToken.SpectatorStream.AddUser(t)
}

func stopSpectate(t *objects.Token) {
	if len(t.SpectatorStream.StreamTokens) <= 1 {
		t.SpectatorStream.RemoveSpectatorStream(t)
	} else {
		t.SpectatorStream.RemoveUser(t)
	}
}

func spectatorFrame(r io.Reader, t *objects.Token) {
	Frames := &objects.SpectatorFrame{}
	osubinary.Unmarshal(r, Frames)

	t.SpectatorStream.Broadcast(r, Frames)
}

// HandlePackets is the Main Packet handler.
func HandlePackets(w http.ResponseWriter, r *http.Request, t *objects.Token) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorln(err)
		return
	}
	pkgs := packets.GetPackets(b)
	pckt := bytes.NewBuffer([]byte{})
	t.LastPing = time.Now()

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

		case constants.ClientStartSpectating:
			startSpectate(r, t)

		case constants.ClientStopSpectating:
			stopSpectate(t)

		case constants.ClientSpectateFrames:
			spectatorFrame(r, t)

		default:
			logPacket(&pkg)

		}
	}
	w.Write(t.Output.Bytes())
	w.Write(pckt.Bytes())
	t.Output.Reset()
}

func logPacket(pkg *constants.Packet) {
	logger.Debugln("---------Packet---------")
	logger.Debugln("PacketID:", pkg.PacketID)
	logger.Debugln("Length:", pkg.PacketLength)
	logger.Debugln("PacketData:", pkg.PacketData)
	logger.Debugln("------------------------")
}

// StartTimeoutChecker Check for all Timeouts (Should only called once.)
func StartTimeoutChecker() {
	go func() {
		for {
			for i := 0; i < len(objects.TOKENS); i++ {
				if objects.TOKENS[i].LastPing.Unix() < (time.Now().Unix()-int64(30)) && objects.TOKENS[i].User.ID != 100 {
					logger.Infof("%s got an Timeout!\n", objects.TOKENS[i].User.UserName)
					disconnectUser(objects.TOKENS[i])
				}
			}
			time.Sleep(time.Second * 5)
		}
	}()
}
