package packetconst

type UserStats struct {
	UserID          int32
	Status          int8
	StatusText      string
	BeatmapChecksum string
	CurrentMods     uint32
	PlayMode        int8
	RankedScore     uint64
	Accuracy        float64
	PlayCount       int32
	TotalScore      uint64
	Rank            int32
	PeppyPoints     int16
}
