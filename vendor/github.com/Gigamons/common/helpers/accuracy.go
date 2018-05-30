package helpers

import (
	"github.com/Gigamons/common/consts"
)

// CalculateAccuracy calculates every playModes accuracy.
func CalculateAccuracy(count300 int64, count100 int64, count50 int64, countMiss int64, countGeki int64, countKatu int64, playMode int8) float64 {
	var thp int64
	var th int64
	var acc float64
	if count300 == 0 {
		count300 = 1
	}
	if count100 == 0 {
		count100 = 1
	}
	if count50 == 0 {
		count50 = 1
	}
	if countMiss == 0 {
		countMiss = 1
	}
	switch playMode {
	case consts.STD:
		thp = int64((count50*50 + count100*100 + count300*300))
		th = int64(countMiss + count50 + count300 + count100)
		acc = float64(thp / (th * 300))
		return acc * 100
	case consts.Taiko:
		thp = int64((count50*50 + count300*100))
		th = int64((countMiss + count100 + count300))
		acc = float64(thp / (th * 100))
		return acc * 100
	case consts.CTB:
		thp = int64(count300 + count100 + count50)
		th = int64(thp + int64(count300+countKatu))
		acc = float64(thp / th)
		return acc * 100
	case consts.Mania:
		thp = int64(count50*50 + count100*100 + countKatu*200 + count300*300 + countGeki*300)
		th = int64(countMiss + count50 + count100 + count300 + countGeki + countKatu)
		acc = float64(thp / (th * 300))
		return acc * 100
	default:
		return 0
	}
}
