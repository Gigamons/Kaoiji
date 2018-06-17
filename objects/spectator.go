package objects

import (
	"io"
	"io/ioutil"
	"sync"

	"github.com/Mempler/osubinary"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
)

type SpectatorStream struct {
	HostToken    *Token
	streamLock   *sync.Mutex
	StreamTokens []*Token
}

type SpectatorFrame struct {
	Extra        int32
	ReplayFrames []ReplayFrame
	Action       int8
	ScoreFrame   ScoreFrame
}

type ReplayFrame struct {
	ButtonState int8
	Button      int8
	MouseX      float32
	MouseY      float32
	Time        int32
}

type ScoreFrame struct {
	Time         int32
	ID           int8
	Count300     uint16
	Count100     uint16
	Count50      uint16
	CountGeki    uint16
	CountKatu    uint16
	CountMiss    uint16
	TotalScore   int32
	MaxCombo     uint16
	CurrentCombo uint16
	FC           bool
	HP           uint8
	TagByte      int8
	ScoreV2      bool
}

func (s *ScoreFrame) ScoreV2F(comboPortion, bonusPortion float64) []byte {
	out := []byte{}
	if s.ScoreV2 {
		x := osubinary.Double(comboPortion)
		y := osubinary.Double(bonusPortion)
		out = append(y, x...)
	}
	return out
}

func NewSpectatorStream(t *Token) *SpectatorStream {
	s := &SpectatorStream{HostToken: t, streamLock: &sync.Mutex{}}
	t.SpectatorStream = s
	return s
}

func (s *SpectatorStream) RemoveSpectatorStream(t *Token) {
	for i := 0; i < len(s.StreamTokens); i++ {
		s.RemoveUser(s.StreamTokens[i])
	}
	s = NewSpectatorStream(t)
}

func (s *SpectatorStream) AddUser(t *Token) {
	if s.AlreadySpectating(t) {
		return
	}
	s.streamLock.Lock()
	s.StreamTokens = append(s.StreamTokens, t)
	s.streamLock.Unlock()
	p1 := constants.NewPacket(constants.BanchoFellowSpectatorJoined)
	p2 := constants.NewPacket(constants.BanchoSpectatorJoined)
	p3 := constants.NewPacket(constants.BanchoChannelJoinSuccess)
	p1.SetPacketData(osubinary.Int32(t.User.ID))
	p2.SetPacketData(osubinary.Int32(t.User.ID))
	p3.SetPacketData(osubinary.BString("#spectator"))
	t.SpectatorStream = s
	s.BroadcastRaw(p2.ToByteArray(), false, nil, true)
	s.BroadcastRaw(p1.ToByteArray(), false, nil, false)
	s.BroadcastRaw(p3.ToByteArray(), false, nil, false)
}

func (s *SpectatorStream) NoMap(t *Token) {
	pack := constants.NewPacket(constants.BanchoSpectatorCantSpectate)
	pack.SetPacketData(osubinary.Int32(t.User.ID))
	s.BroadcastRaw(pack.ToByteArray(), false, nil, false)
}

func (s *SpectatorStream) RemoveUser(t *Token) {
	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == t {
			p1 := constants.NewPacket(constants.BanchoFellowSpectatorLeft)
			p2 := constants.NewPacket(constants.BanchoSpectatorLeft)
			p1.SetPacketData(osubinary.Int32(s.StreamTokens[i].User.ID))
			p2.SetPacketData(osubinary.Int32(s.StreamTokens[i].User.ID))
			s.BroadcastRaw(p1.ToByteArray(), false, nil, false)
			s.BroadcastRaw(p2.ToByteArray(), false, nil, true)
			s.streamLock.Lock()
			copy(s.StreamTokens[i:], s.StreamTokens[i+1:])
			s.StreamTokens[len(s.StreamTokens)-1] = nil
			s.StreamTokens = s.StreamTokens[:len(s.StreamTokens)-1]
			s.streamLock.Unlock()
			t.SpectatorStream = nil
		}
	}
}

func (s *SpectatorStream) AlreadySpectating(t *Token) bool {
	s.streamLock.Lock()
	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == t {
			s.streamLock.Unlock()
			return true
		}
	}
	s.streamLock.Unlock()
	return false
}

func (s *SpectatorStream) Broadcast(r io.Reader) {
	// f := SpectatorFrame{}
	//
	// f.Extra, _ = osubinary.RInt32(r)
	// ReplayFrameCount, _ := osubinary.RUInt16(r)
	// for i := uint16(0); i < ReplayFrameCount; i++ {
	// 	Frame := ReplayFrame{}
	// 	Frame.ButtonState, _ = osubinary.RInt8(r)
	// 	Frame.Button, _ = osubinary.RInt8(r)
	// 	Frame.MouseX, _ = osubinary.RFloat(r)
	// 	Frame.MouseY, _ = osubinary.RFloat(r)
	// 	Frame.Time, _ = osubinary.RInt32(r)
	// }
	// osubinary.Unmarshal(r, &f.ScoreFrame)
	//
	// xbuf := bytes.NewBuffer(nil)
	// xbuf.Write(osubinary.Int32(f.Extra))
	// xbuf.Write(osubinary.UInt16(ReplayFrameCount))
	// for i := 0; i < len(f.ReplayFrames); i++ {
	// 	xbuf.Write(osubinary.Int8(f.ReplayFrames[i].ButtonState))
	// 	xbuf.Write(osubinary.Int8(f.ReplayFrames[i].Button))
	// 	xbuf.Write(osubinary.Float(f.ReplayFrames[i].MouseX))
	// 	xbuf.Write(osubinary.Float(f.ReplayFrames[i].MouseY))
	// 	xbuf.Write(osubinary.Int32(f.ReplayFrames[i].Time))
	// }
	// xbuf.Write(osubinary.Marshal(f.ScoreFrame))
	// if f.ScoreFrame.ScoreV2 {
	// 	ComboPortion, _ := osubinary.RDouble(r)
	// 	BonusPortion, _ := osubinary.RDouble(r)
	// 	xbuf.Write(f.ScoreFrame.ScoreV2F(ComboPortion, BonusPortion))
	// }
	//
	// pack := constants.NewPacket(constants.BanchoSpectateFrames)
	// pack.SetPacketData(xbuf.Bytes())
	x, err := ioutil.ReadAll(r)
	if err != nil {
		logger.Errorln(err)
	}
	logger.Debugln("Broadcasting", s.HostToken.User.UserName, "frames to", len(s.StreamTokens), "users")
	pack := constants.NewPacket(constants.BanchoSpectateFrames)
	pack.SetPacketData(x)
	go s.BroadcastRaw(pack.ToByteArray(), true, s.HostToken, false)
}

func (s *SpectatorStream) BroadcastRaw(b []byte, isFrame bool, ignoreSelf *Token, onlyHost bool) {
	if !isFrame {
		s.HostToken.Write(b)
	}
	if onlyHost {
		s.HostToken.Write(b)
		return
	}
	x := constants.UserStatsStruct{}

	x.UserID = s.HostToken.User.ID
	x.Status = s.HostToken.Status.Beatmap.Status
	x.StatusText = s.HostToken.Status.Beatmap.StatusText
	x.BeatmapChecksum = s.HostToken.Status.Beatmap.BeatmapChecksum
	x.CurrentMods = s.HostToken.Status.Beatmap.CurrentMods
	x.PlayMode = s.HostToken.Status.Beatmap.PlayMode
	x.BeatmapID = s.HostToken.Status.Beatmap.BeatmapID
	x.RankedScore = uint64(s.HostToken.Leaderboard.RankedScore)
	x.Accuracy = float32(helpers.CalculateAccuracy(s.HostToken.Leaderboard.Count300, s.HostToken.Leaderboard.Count100, s.HostToken.Leaderboard.Count50, s.HostToken.Leaderboard.CountMiss, 0, 0, 0))
	x.PlayCount = int32(s.HostToken.Leaderboard.Playcount)
	x.TotalScore = uint64(s.HostToken.Leaderboard.TotalScore)
	x.PeppyPoints = int16(s.HostToken.Leaderboard.PeppyPoints)

	pckt := constants.NewPacket(constants.ClientSendUserStatus)
	pckt.SetPacketData(osubinary.Marshal(x))

	for i := 0; i < len(s.StreamTokens); i++ {
		if s.StreamTokens[i] == ignoreSelf {
			continue
		}
		go s.StreamTokens[i].Write(pckt.ToByteArray())
		go s.StreamTokens[i].Write(b)
	}
}
