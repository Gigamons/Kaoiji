package packets

import (
	"bytes"
	"github.com/Gigamons/Shared/shelpers"
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

	shelpers.WriteBytes(buff, r.MatchId)
	shelpers.WriteBytes(buff, r.InProgress)
	shelpers.WriteBytes(buff, r.MatchType)
	shelpers.WriteBytes(buff, r.ActiveMods)
	shelpers.WriteBytes(buff, r.Name)
	shelpers.WriteBytes(buff, r.Password)
	shelpers.WriteBytes(buff, r.BeatmapName)
	shelpers.WriteBytes(buff, r.BeatmapId)
	shelpers.WriteBytes(buff, r.BeatmapMd5)

	for _, item := range r.Slots {
		shelpers.WriteBytes(buff, item.Status)
	}
	for _, item := range r.Slots {
		shelpers.WriteBytes(buff, item.Team)
	}
	for _, item := range r.Slots {
		shelpers.WriteBytes(buff, item.UserId)
	}

	shelpers.WriteBytes(buff, r.HostId)
	shelpers.WriteBytes(buff, r.PlayMode)
	shelpers.WriteBytes(buff, r.ScoringType)
	shelpers.WriteBytes(buff, r.TeamType)
	shelpers.WriteBytes(buff, r.SpecialModes)

	if r.SpecialModes != 0 {
		for _, item := range r.Slots {
			shelpers.WriteBytes(buff, item.Mods)
		}
	}

	shelpers.WriteBytes(buff, r.Seed)

	return buff.Bytes()
}
func (r *MultiPlayerRoom) WriteBytes(w io.Writer) error {

	if err := shelpers.WriteBytes(w, r.MatchId); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.InProgress); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.MatchType); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.ActiveMods); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.Name); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.Password); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.BeatmapName); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.BeatmapId); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.BeatmapMd5); err != nil {
		return err
	}

	for _, item := range r.Slots {
		if err := shelpers.WriteBytes(w, item.Status); err != nil {
			return err
		}
	}
	for _, item := range r.Slots {
		if err := shelpers.WriteBytes(w, item.Team); err != nil {
			return err
		}
	}
	for _, item := range r.Slots {
		if err := shelpers.WriteBytes(w, item.UserId); err != nil {
			return err
		}
	}

	if err := shelpers.WriteBytes(w, r.HostId); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.PlayMode); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.ScoringType); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.TeamType); err != nil {
		return err
	}
	if err := shelpers.WriteBytes(w, r.SpecialModes); err != nil {
		return err
	}

	if r.SpecialModes != 0 {
		for _, item := range r.Slots {
			if err := shelpers.WriteBytes(w, item.Mods); err != nil {
				return err
			}
		}
	}

	return shelpers.WriteBytes(w, r.Seed)
}
