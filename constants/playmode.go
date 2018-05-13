package constants

// Playmodes
const (
	STD   = 0
	Taiko = 1
	CTB   = 2
	Mania = 3
)

func ToPlaymodeString(p int8) string {
	switch p {
	case STD:
		return "std"
	case Taiko:
		return "taiko"
	case CTB:
		return "ctb"
	case Mania:
		return "mania"
	}
	return "osu"
}
