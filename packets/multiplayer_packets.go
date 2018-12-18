package packets

import (
	"bytes"
	"github.com/Gigamons/Kaoiji/helpers"
	"io"
)

const maxPlayers = 16

type MultiPlayerRoom struct {
	MatchId int32

	InProgress bool

	MatchType byte
	ActiveMods uint32

	Name string
	Password string

	BeatmapName string
	BeatmapId int32
	BeatmapMd5 string

	Slots [maxPlayers]MultiPlayerSlot

	HostId int32
	PlayMode byte
	ScoringType byte
	TeamType byte
	SpecialModes byte

	Seed int32
}

type MultiPlayerSlot struct {
	Status byte
	Team   byte
	UserId int32
	Mods   uint32
}

func (r *MultiPlayerRoom) GetBytes() []byte {
	buff := new (bytes.Buffer)

	helpers.WriteBytes(buff, r.MatchId)
	helpers.WriteBytes(buff, r.InProgress)
	helpers.WriteBytes(buff, r.MatchType)
	helpers.WriteBytes(buff, r.ActiveMods)
	helpers.WriteBytes(buff, r.Name)
	helpers.WriteBytes(buff, r.Password)
	helpers.WriteBytes(buff, r.BeatmapName)
	helpers.WriteBytes(buff, r.BeatmapId)
	helpers.WriteBytes(buff, r.BeatmapMd5)

	for _, item := range r.Slots {
		helpers.WriteBytes(buff, item.Status)
	}
	for _, item := range r.Slots {
		helpers.WriteBytes(buff, item.Team)
	}
	for _, item := range r.Slots {
		helpers.WriteBytes(buff, item.UserId)
	}

	helpers.WriteBytes(buff, r.HostId)
	helpers.WriteBytes(buff, r.PlayMode)
	helpers.WriteBytes(buff, r.ScoringType)
	helpers.WriteBytes(buff, r.TeamType)
	helpers.WriteBytes(buff, r.SpecialModes)

	if r.SpecialModes != 0 {
		for _, item := range r.Slots {
			helpers.WriteBytes(buff, item.Mods)
		}
	}

	helpers.WriteBytes(buff, r.Seed)

	return buff.Bytes()
}
func (r *MultiPlayerRoom) WriteBytes(w io.Writer) error {

	if err := helpers.WriteBytes(w, r.MatchId); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.InProgress); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.MatchType); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.ActiveMods); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.Name); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.Password); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.BeatmapName); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.BeatmapId); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.BeatmapMd5); err != nil {
		return err
	}

	for _, item := range r.Slots {
		if err := helpers.WriteBytes(w, item.Status); err != nil {
			return err
		}
	}
	for _, item := range r.Slots {
		if err := helpers.WriteBytes(w, item.Team); err != nil {
			return err
		}
	}
	for _, item := range r.Slots {
		if err := helpers.WriteBytes(w, item.UserId); err != nil {
			return err
		}
	}

	if err := helpers.WriteBytes(w, r.HostId); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.PlayMode); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.ScoringType); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.TeamType); err != nil {
		return err
	}
	if err := helpers.WriteBytes(w, r.SpecialModes); err != nil {
		return err
	}

	if r.SpecialModes != 0 {
		for _, item := range r.Slots {
			if err := helpers.WriteBytes(w, item.Mods); err != nil {
				return err
			}
		}
	}

	return helpers.WriteBytes(w, r.Seed)
}
