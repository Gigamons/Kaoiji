package constants

import (
	"bytes"
	"encoding/binary"

	"github.com/Gigamons/common/consts"
)

// BeatmapInfo is to send the BeatmapInfo to our client. IDK How but i managed to reverse engineer it.
type BeatmapInfo struct {
	ScoreID          uint16
	BeatmapID        uint32
	BeatmapSetID     uint32
	ForumThreadID    uint32
	RankedStatus     byte
	OsuLetter        byte
	CTBLetter        byte
	TaikoLetter      byte
	ManiaLetter      byte
	BeatmapChecksumm string
}

// ClientSendUserStatusStruct for user status
type ClientSendUserStatusStruct struct {
	Status          byte
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        byte
	BeatmapID       uint32
}

// Config is for the config.yml file
type Config struct {
	Server struct {
		Hostname   string
		Port       int
		FreeDirect bool
		Debug      bool
	}
	API struct {
		KokoroHost   string
		KokoroAPIKey string
		APIKey       string
	}
	MySQL consts.MySQLConf
	Redis struct {
		Hostname string
		Port     int
	}
}

// MessageStruct for Messages
type MessageStruct struct {
	Username string
	Message  string
	Target   string
	UserID   uint32
}

// UserPresenceStruct for User presences
type UserPresenceStruct struct {
	UserID      uint32
	Username    string
	Timezone    byte
	CountryID   byte
	Permissions byte
	Lon         float64
	Lat         float64
	Rank        uint32
}

// UserQuitStruct if a user quits/timeout
type UserQuitStruct struct {
	UserID     uint32
	ErrorState byte
}

// UserStatsStruct Stats ofc.
type UserStatsStruct struct {
	UserID          uint32
	Status          byte
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        byte
	BeatmapID       uint32
	RankedScore     uint64
	Accuracy        float32
	PlayCount       uint32
	TotalScore      uint64
	Rank            uint32
	PeppyPoints     uint16
}

// Packet struct is simply a packet
type Packet struct {
	PacketID     uint16
	PacketLength uint32
	PacketData   []byte
}

// NewPacket Create a new Packet
func NewPacket(packetid int) *Packet {
	return &Packet{uint16(packetid), 0, nil}
}

// SetPacketData set's the Packet data
func (p *Packet) SetPacketData(PacketData []byte) {
	p.PacketLength = uint32(len(PacketData))
	p.PacketData = PacketData
}

// ToByteArray retuns a ByteArray of the written Packet.
func (p *Packet) ToByteArray() []byte {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, p.PacketID)
	binary.Write(b, binary.LittleEndian, int8(0))
	binary.Write(b, binary.LittleEndian, p.PacketLength)
	binary.Write(b, binary.LittleEndian, p.PacketData)

	return b.Bytes()
}
