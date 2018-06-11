package constants

import "github.com/Gigamons/common/consts"

// BeatmapInfo is to send the BeatmapInfo to our client. IDK How but i managed to reverse engineer it.
type BeatmapInfo struct {
	ScoreID          int16
	BeatmapID        int32
	BeatmapSetID     int32
	ForumThreadID    int32
	RankedStatus     int8
	OsuLetter        int8
	CTBLetter        int8
	TaikoLetter      int8
	ManiaLetter      int8
	BeatmapChecksumm string
}

type ClientSendUserStatusStruct struct {
	Status          int8
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        int8
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

type MessageStruct struct {
	Username string
	Message  string
	Target   string
	UserID   int32
}

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

type UserQuitStruct struct {
	UserID     int32
	ErrorState int8
}

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
