package handlers

import (
	"bytes"
	"io"
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
func sendUserStatus(r io.Reader, t *objects.Token) {
	x := constants.ClientSendUserStatusStruct{}
	osubinary.Unmarshal(r, &x)
	t.Status.Beatmap = x
	if x.CurrentMods&128 > 0 || x.CurrentMods&8192 > 0 {
		if !t.AlreadyNotified {
			t.AlreadyNotified = true
			w := packets.NewWriter(t)
			w.Announce("You've enabled the RX/AP Scoreboard.\nAP has Nerfed aim PP.\nRX has Nerfed speed PP.")
			go t.Write(w.Bytes())
		}
		if !t.User.Relax {
			t.Write(public.SendUserStats(t, true))
		} else {
			t.Write(public.SendUserStats(t, false))
		}
		t.User.Relax = true
	} else {
		if t.User.Relax {
			t.Write(public.SendUserStats(t, true))
		} else {
			t.Write(public.SendUserStats(t, false))
		}
		t.User.Relax = false
	}
}

// joinChannel sends a Join successfull/fail to the Client.
func joinChannel(r io.Reader, t *objects.Token) {
	yw := packets.NewWriter(t)
	xw := objects.ChannelInfo{}
	osubinary.Unmarshal(r, &xw)
	yw.JoinChannel(xw.ChannelName)
	yw.ChannelAvaible()
	t.Write(yw.Bytes())
}

func leaveChannel(r io.Reader, t *objects.Token) {
	yw := packets.NewWriter(t)
	xw := objects.ChannelInfo{}
	osubinary.Unmarshal(r, &xw)
	yw.LeaveChannel(xw.ChannelName)
	yw.ChannelAvaible()
	t.Write(yw.Bytes())
}

// addFriend adds a Friend ofc!
func addFriend(r io.Reader, t *objects.Token) {
	i, err := osubinary.RUInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	packets.AddFriend(t.User, usertools.GetUser(&i))
}

// removeFriend removes a Friend ofc!
func removeFriend(r io.Reader, t *objects.Token) {
	i, err := osubinary.RUInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	packets.RemoveFriend(t.User, usertools.GetUser(&i))
}

// updateUserStats updates the User stats for that user, (No fetching out of SQL)
func updateUserStats(r io.Reader, t *objects.Token) {
	_, err := osubinary.RIntArray(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	for y := 0; y < len(objects.TOKENS); y++ {
		t.Write(public.SendUserStats(objects.TOKENS[y], false))
	}
}

// sendUserPresence Send the User Precense of the Given UserID.
func sendUserPresence(r io.Reader, t *objects.Token) {
	i, err := osubinary.RUIntArray(r)
	yw := packets.NewWriter(t)
	if err != nil {
		logger.Errorln(err)
		return
	}
	for y := 0; y < len(i); y++ {
		yw.UserPresence(objects.GetTokenByID(i[y]))
	}
	t.Write(yw.Bytes())
}

// sendMessage User sends a MSG Packet to us, we're handling it and send it to the User Target
func sendMessage(r io.Reader, t *objects.Token) {
	msg := constants.MessageStruct{}
	osubinary.Unmarshal(r, &msg)
	go public.SendMessage(t, msg.Message, msg.Target)
}

// disconnectUser Our user has a Timeout nor He Disconnects, we're broadcast it to Everyone that that user got a Timeout / Disconnect.
func disconnectUser(t *objects.Token) {
	main := objects.GetStream("main")
	pckt := constants.NewPacket(constants.BanchoHandleUserQuit)
	pckt.SetPacketData(osubinary.Marshal(constants.UserQuitStruct{UserID: t.User.ID, ErrorState: 0}))
	objects.DeleteToken(t.Token)
	go main.Broadcast(pckt.ToByteArray(), nil)
}

func startSpectate(r io.Reader, t *objects.Token) {
	i, err := osubinary.RUInt32(r)
	if err != nil {
		logger.Errorln(err)
		return
	}
	HostToken := objects.GetTokenByID(i)
	if HostToken.SpectatorStream == nil {
		HostToken.SpectatorStream = objects.NewSpectatorStream(HostToken)
	}
	go HostToken.SpectatorStream.AddUser(t)
}

func stopSpectate(t *objects.Token) {
	if len(t.SpectatorStream.StreamTokens) <= 1 {
		go t.SpectatorStream.RemoveSpectatorStream(t)
	} else {
		go t.SpectatorStream.RemoveUser(t)
	}
}

func specNoMap(t *objects.Token) {
	t.SpectatorStream.NoMap(t)
}

func spectatorFrame(r io.Reader, t *objects.Token) {
	go t.SpectatorStream.Broadcast(r)
}

func beatmapInfo(r io.Reader, t *objects.Token) {
	packets.BeatmapInfoRequest(r)
}

func LobbyJoin(t *objects.Token) {
	s := objects.GetStream("lobby")
	if s == nil {
		return
	}
	s.AddUser(t)
	objects.GetLobbys(t)
}

func LobbyLeave(t *objects.Token) {
	s := objects.GetStream("lobby")
	if s == nil {
		return
	}
	s.RemoveUser(t)
}

func JoinMPLobby(r io.Reader, t *objects.Token) {
	i, _ := osubinary.RUInt32(r)
	s, _ := osubinary.RBString(r)
	_, l := objects.GetLobby(uint16(i))
	if l == nil {
		return
	}
	objects.JoinLobby(l, s, t)
}

func LeaveMPLobby(t *objects.Token) {
	objects.LeaveLobby(t)
}

func CreateMPLobby(r io.Reader, t *objects.Token) {
	l := objects.ReadLobby(r)
	objects.NewLobbyC(l, t)
}

func SwitchLobbySlot(r io.Reader, t *objects.Token) {
	slot, _ := osubinary.RByte(r)
	t.MPLobby.SwitchSlot(slot, t)
}

// HandlePackets is the Main Packet handler.
func HandlePackets(b []byte, t *objects.Token) {

	pkgs := packets.GetPackets(b)
	t.LastPing = time.Now()

	for i := 0; i < len(pkgs); i++ {
		pkg := pkgs[i]
		r := bytes.NewReader(pkg.PacketData)
		switch pkg.PacketID {

		case constants.ClientSendUserStatus:
			sendUserStatus(r, t)

		case constants.ClientExit:
			disconnectUser(t)

		case constants.ClientSendIrcMessage:
			sendMessage(r, t)

		case constants.ClientSendIrcMessagePrivate:
			sendMessage(r, t)

		case constants.ClientRequestStatusUpdate:
			t.Write(public.SendUserStats(t, true))

		case constants.ClientPong:
			t.LastPing = time.Now()

		case constants.ClientReceiveUpdates:
			t.Write(public.SendUserStats(t, true))

		case constants.ClientChannelJoin:
			joinChannel(r, t)

		case constants.ClientChannelLeave:
			leaveChannel(r, t)

		case constants.ClientFriendAdd:
			addFriend(r, t)

		case constants.ClientFriendRemove:
			removeFriend(r, t)

		case constants.ClientUserStatsRequest:
			updateUserStats(r, t)

		case constants.ClientUserPresenceRequest:
			sendUserPresence(r, t)

		case constants.ClientStartSpectating:
			startSpectate(r, t)

		case constants.ClientStopSpectating:
			stopSpectate(t)

		case constants.ClientSpectateFrames:
			spectatorFrame(r, t)

		case constants.ClientCantSpectate:
			specNoMap(t)

		case constants.ClientBeatmapInfoRequest:
			beatmapInfo(r, t)

		case constants.ClientLobbyJoin:
			LobbyJoin(t)

		case constants.ClientLobbyPart:
			LobbyLeave(t)

		case constants.ClientMatchCreate:
			CreateMPLobby(r, t)

		case constants.ClientMatchJoin:
			JoinMPLobby(r, t)

		case constants.ClientMatchPart:
			LeaveMPLobby(t)

		case constants.ClientMatchChangeSlot:
			SwitchLobbySlot(r, t)

		default:
			logPacket(pkg)

		}
	}
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
