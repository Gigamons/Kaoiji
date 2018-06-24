package objects

import (
	"bytes"
	"io"
	"sync"

	"github.com/Gigamons/common/logger"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/helpers"
	"github.com/Mempler/osubinary"
)

type Lobby struct {
	ID          uint16
	Running     bool
	Type        int8
	Mods        uint32
	Name        string
	Password    string
	BeatmapName string
	BeatmapID   uint32
	BeatmapMD5  string
	Slots       [16]LobbySlot
	Host        int32
	PlayMode    int8
	ScoreType   int8
	TeamType    int8
	FreeMods    int8
	Seed        uint32 // or random number, idk.
}

type LobbySlot struct {
	Status uint8
	UserID int32
	Team   int8
	Mods   uint32
}

var LOBBYS []*Lobby

var LobbyLock *sync.Mutex

func init() {
	LobbyLock = &sync.Mutex{}
}

// Only used for lobbys created by non osu!
func NewLobby(Name string, Password string) *Lobby {
	l := &Lobby{ID: uint16(len(LOBBYS)), Name: Name, Password: Password}
	LobbyLock.Lock()
	LOBBYS = append(LOBBYS, l)
	LobbyLock.Unlock()
	return l
}

func NewLobbyC(l *Lobby, t *Token) {
	LobbyLock.Lock()
	LOBBYS = append(LOBBYS, l)
	LobbyLock.Unlock()
	l.ID = uint16(len(LOBBYS))
	pckt := constants.NewPacket(constants.BanchoMatchNew)
	pckt.SetPacketData(WriteLobby(l, false))
	t.Write(pckt.ToByteArray())
	JoinLobby(l, l.Password, t)
}

func JoinLobby(l *Lobby, password string, t *Token) {
	if t == nil {
		return
	}
	_, s := GetLobby(l.ID)
	f := true
	for i := 0; i < len(s.Slots); i++ {
		if s.Slots[i].Status == 1 || s.Slots[i].Status == 128 {
			f = false
			if password != l.Password {
				pckt := constants.NewPacket(constants.BanchoMatchJoinFail)
				pckt.SetPacketData(WriteLobby(l, false))
				t.Write(pckt.ToByteArray())
				return
			}

			s.Slots[i].UserID = t.User.ID
			s.Slots[i].Status = constants.SlotNotReady
			s.Slots[i].Team = 0
			s.Slots[i].Mods = 0

			t.MPSlot = int8(i)
			t.MPLobby = s

			pckt := constants.NewPacket(constants.BanchoMatchJoinSuccess)
			pckt.SetPacketData(WriteLobby(l, false))
			t.Write(pckt.ToByteArray())
			break
		}
	}
	if f {
		pckt := constants.NewPacket(constants.BanchoMatchJoinFail)
		pckt.SetPacketData(WriteLobby(l, false))
		t.Write(pckt.ToByteArray())
	}
	UpdateMatch(l)
}

func LeaveLobby(t *Token) {
	if t.MPLobby == nil {
		return
	}
	for i := 0; i < len(t.MPLobby.Slots); i++ {
		if t.MPLobby.Slots[i].UserID == t.User.ID {
			if t.MPLobby.Running {
				t.MPLobby.Slots[i].UserID = -1
				t.MPLobby.Slots[i].Status = constants.SlotQuit
				t.MPSlot = 0
				UpdateMatch(t.MPLobby)
				t.MPLobby = nil
				break
			} else {
				t.MPLobby.Slots[i].UserID = -1
				t.MPLobby.Slots[i].Team = 0
				t.MPLobby.Slots[i].Status = constants.SlotOpen
				t.MPLobby.Slots[i].Mods = 0
				t.MPSlot = 0
				UpdateMatch(t.MPLobby)
				t.MPLobby = nil
				break
			}
		}
	}
}

func (l *Lobby) SwitchSlot(SlotID int8, t *Token) {
	logger.Debugln(t.User.UserName, "Switched his slot from", t.MPSlot)
	if l.Slots[SlotID].Status == 1 || l.Slots[SlotID].Status == 128 {
		l.Slots[SlotID].UserID = l.Slots[t.MPSlot].UserID
		l.Slots[SlotID].Status = l.Slots[t.MPSlot].Status
		l.Slots[SlotID].Team = l.Slots[t.MPSlot].Team
		l.Slots[SlotID].Mods = l.Slots[t.MPSlot].Mods
		l.Slots[t.MPSlot].UserID = -1
		l.Slots[t.MPSlot].Status = constants.SlotOpen
		l.Slots[t.MPSlot].Team = 1
		l.Slots[t.MPSlot].Mods = 0
		t.MPSlot = SlotID
		logger.Debugln(t.User.UserName, "Switched his slot to", t.MPSlot)
	}
	UpdateMatch(l)
}

func UpdateMatch(l *Lobby) {
	s := GetStream("lobby")
	if s != nil {
		pckt := constants.NewPacket(constants.BanchoMatchUpdate)
		pckt.SetPacketData(WriteLobby(l, true))
		s.Broadcast(pckt.ToByteArray(), nil)
	}
}

func GetLobby(ID uint16) (int, *Lobby) {
	for i := 0; i < len(LOBBYS); i++ {
		if LOBBYS[i].ID == ID {
			return i, LOBBYS[i]
		}
	}
	return -1, nil
}

func RemoveLobby(ID uint16) {
	i, l := GetLobby(ID)
	if i == -1 || l == nil {
		return
	}
	copy(LOBBYS[i:], LOBBYS[i+1:])
	LOBBYS[len(LOBBYS)-1] = nil
	LOBBYS = LOBBYS[:len(LOBBYS)-1]
}

func GetLobbys(t *Token) {
	for i := 0; i < len(LOBBYS); i++ {
		pckt := constants.NewPacket(constants.BanchoMatchNew)
		pckt.SetPacketData(WriteLobby(LOBBYS[i], true))
		t.Write(pckt.ToByteArray())
	}
}

func WriteLobby(l *Lobby, h bool) []byte {
	if l == nil {
		return nil
	}
	buf := bytes.NewBuffer(nil)
	buf.Write(osubinary.UInt16(l.ID))
	buf.Write(osubinary.Bool(l.Running))
	buf.Write(osubinary.Int8(l.Type))
	buf.Write(osubinary.UInt32(l.Mods))
	buf.Write(osubinary.BString(l.Name))
	if h {
		if len(l.Password) > 0 {
			p, _ := helpers.MD5(helpers.Pseudorandombytes(8))
			buf.Write(osubinary.BString(string(p)))
		} else {
			buf.Write(osubinary.BString(""))
		}
	} else {
		buf.Write(osubinary.BString(l.Password))
	}
	buf.Write(osubinary.BString(l.BeatmapName))
	buf.Write(osubinary.UInt32(l.BeatmapID))
	buf.Write(osubinary.BString(l.BeatmapMD5))
	for i := 0; i < 15; i++ {
		buf.Write(osubinary.UInt8(l.Slots[i].Status))
	}
	for i := 0; i < 15; i++ {
		buf.Write(osubinary.Int8(l.Slots[i].Team))
	}
	for i := 0; i < 15; i++ {
		buf.Write(osubinary.Int32(l.Slots[i].UserID))
	}
	buf.Write(osubinary.Int32(l.Host))
	buf.Write(osubinary.Int8(l.PlayMode))
	buf.Write(osubinary.Int8(l.ScoreType))
	buf.Write(osubinary.Int8(l.FreeMods))
	if l.FreeMods&constants.Freemod > 0 {
		for i := 0; i < 15; i++ {
			buf.Write(osubinary.UInt32(l.Slots[i].Mods))
		}
	}
	buf.Write(osubinary.UInt32(l.Seed))
	return buf.Bytes()
}

func ReadLobby(r io.Reader) *Lobby {
	l := &Lobby{}
	l.ID, _ = osubinary.RUInt16(r)
	l.Running, _ = osubinary.RBool(r)
	l.Type, _ = osubinary.RInt8(r)
	l.Mods, _ = osubinary.RUInt32(r)
	l.Name, _ = osubinary.RBString(r)
	l.Password, _ = osubinary.RBString(r)
	l.BeatmapName, _ = osubinary.RBString(r)
	l.BeatmapID, _ = osubinary.RUInt32(r)
	l.BeatmapMD5, _ = osubinary.RBString(r)
	for i := 0; i < 15; i++ {
		l.Slots[i].Status, _ = osubinary.RUInt8(r)
	}
	for i := 0; i < 15; i++ {
		l.Slots[i].Team, _ = osubinary.RInt8(r)
	}
	for i := 0; i < 15; i++ {
		l.Slots[i].UserID, _ = osubinary.RInt32(r)
	}
	l.PlayMode, _ = osubinary.RInt8(r)
	l.ScoreType, _ = osubinary.RInt8(r)
	l.FreeMods, _ = osubinary.RInt8(r)
	if l.FreeMods&constants.Freemod > 0 {
		for i := 0; i < 15; i++ {
			l.Slots[i].Mods, _ = osubinary.RUInt32(r)
		}
	}
	l.Seed, _ = osubinary.RUInt32(r)
	return l
}
