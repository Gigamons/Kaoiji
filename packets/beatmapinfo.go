package packets

import (
	"io"

	"github.com/Mempler/osubinary"
)

func BeatmapInfoRequest(r io.Reader) []string {
	beatmaplength, _ := osubinary.RUInt32(r)

	x := make([]string, beatmaplength)

	for i := 0; i < len(x); i++ {
		beatmap, _ := osubinary.RBString(r)
		x[i] = beatmap
	}

	return x
}

func getscores() {
	//sql := "SELECT ScoreID, FileMD5, Accuracy, (SELECT BeatmapID, SetID, RankedStatus FROM beatmaps WHERE scores.FileMD5 = beatmaps.FileMD5) FROM scores WHERE UserID = ? AND FileMD5 = ? ORDER BY score DESC LIMIT 1"
}
