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
	RankedStatus     int8
	OsuLetter        int8
	CTBLetter        int8
	TaikoLetter      int8
	ManiaLetter      int8
	BeatmapChecksumm string
}

// ClientSendUserStatusStruct for user status
type ClientSendUserStatusStruct struct {
	Status          int8
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        uint8
	BeatmapID       int32
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
	UserID   int32
}

// UserPresenceStruct for User presences
type UserPresenceStruct struct {
	UserID      int32
	Username    string
	Timezone    int8
	CountryID   int8
	Permissions int8
	Lon         float64
	Lat         float64
	Rank        int32
}

// UserQuitStruct if a user quits/timeout
type UserQuitStruct struct {
	UserID     int32
	ErrorState int8
}

// UserStatsStruct Stats ofc.
type UserStatsStruct struct {
	UserID          int32
	Status          int8
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        int8
	BeatmapID       int32
	RankedScore     uint64
	Accuracy        float32
	PlayCount       int32
	TotalScore      uint64
	Rank            int32
	PeppyPoints     int16
}

// Packet struct is simply a packet
type Packet struct {
	PacketID     int16
	PacketLength int32
	PacketData   []byte
}

// NewPacket Create a new Packet
func NewPacket(packetid int) Packet {
	return Packet{int16(packetid), 0, nil}
}

// SetPacketData set's the Packet data
func (p *Packet) SetPacketData(PacketData []byte) {
	p.PacketLength = int32(len(PacketData))
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
