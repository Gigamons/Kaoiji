package constants

type ClientSendUserStatusStruct struct {
	Status          int8
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        int8
	BeatmapID       int32
}
