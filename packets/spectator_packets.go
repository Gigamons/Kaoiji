package packets

import (
	"bytes"
	"github.com/cyanidee/bancho-go/consts"
	"github.com/cyanidee/bancho-go/helpers"
	"io"
)

type ScoreFrame struct {
	Count300     uint16
	Count100     uint16
	Count50      uint16
	CountGeki    uint16
	CountKatu    uint16
	CountMiss    uint16
	CurrentCombo uint16

	CurrentHp    byte
	Id           byte  // Multiplayer = SlotId
	MaxCombo     uint16
	Perfect      bool

	ScoreV2      bool
	TagByte      byte

	Time         int32
	TotalScore   int32

	BonusPortion float64
	ComboPortion float64
}

type ReplayFrame struct {
	Button      byte
	ButtonState byte
	MouseX      float32
	MouseY      float32
	Time        int32
}

type SpectatorFrame struct {
	Action byte
	Extra  int

	ReplayFrames []ReplayFrame
	ScoreFrame   ScoreFrame
}

func (s *ScoreFrame) GetBytes() []byte {
	buff := new(bytes.Buffer)

	buff.Write(helpers.GetBytes(s.Time))
	buff.Write(helpers.GetBytes(s.Id))
	buff.Write(helpers.GetBytes(s.Count300))
	buff.Write(helpers.GetBytes(s.Count100))
	buff.Write(helpers.GetBytes(s.Count50))
	buff.Write(helpers.GetBytes(s.CountGeki))
	buff.Write(helpers.GetBytes(s.CountKatu))
	buff.Write(helpers.GetBytes(s.CountMiss))
	buff.Write(helpers.GetBytes(s.TotalScore))
	buff.Write(helpers.GetBytes(s.MaxCombo))
	buff.Write(helpers.GetBytes(s.CurrentCombo))
	buff.Write(helpers.GetBytes(s.Perfect))
	buff.Write(helpers.GetBytes(s.CurrentHp))
	buff.Write(helpers.GetBytes(s.TagByte))
	buff.Write(helpers.GetBytes(s.ScoreV2))
	if s.ScoreV2 {
		buff.Write(helpers.GetBytes(s.ComboPortion))
		buff.Write(helpers.GetBytes(s.BonusPortion))
	}

	return buff.Bytes()
}

func (r *ReplayFrame) GetBytes() []byte {
	buff := new(bytes.Buffer)
	buff.Write(helpers.GetBytes(r.ButtonState))
	buff.Write(helpers.GetBytes(r.Button))
	buff.Write(helpers.GetBytes(r.MouseX))
	buff.Write(helpers.GetBytes(r.MouseY))
	buff.Write(helpers.GetBytes(r.Time))
	return buff.Bytes()
}

func (sf *SpectatorFrame) GetBytes() []byte {
	buff := new(bytes.Buffer)

	buff.Write(helpers.GetBytes(sf.Extra))
	buff.Write(helpers.GetBytes(uint16(len(sf.ReplayFrames))))

	for _, rf := range sf.ReplayFrames  {
		buff.Write(rf.GetBytes())
	}

	buff.Write(helpers.GetBytes(sf.Action))
	buff.Write(sf.ScoreFrame.GetBytes())

	return buff.Bytes()
}

func (s *ScoreFrame) WriteBytes(w io.Writer) error {
	if _, err := w.Write(helpers.GetBytes(s.Time)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.Id)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.Count300)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.Count100)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.Count50)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.CountGeki)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.CountKatu)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.CountMiss)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.TotalScore)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.MaxCombo)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.CurrentCombo)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.Perfect)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.CurrentHp)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.TagByte)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(s.ScoreV2)); err != nil {
		return err
	}
	if s.ScoreV2 {
		if _, err := w.Write(helpers.GetBytes(s.ComboPortion)); err != nil {
			return err
		}
		if _, err := w.Write(helpers.GetBytes(s.BonusPortion)); err != nil {
			return err
		}
	}

	return nil
}

func (r *ReplayFrame) WriteBytes(w io.Writer) error {
	if _, err := w.Write(helpers.GetBytes(r.ButtonState)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(r.Button)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(r.MouseX)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(r.MouseY)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(r.Time)); err != nil {
		return err
	}
	return nil
}

func (sf *SpectatorFrame) WriteBytes(w io.Writer) error {
	if _, err := w.Write(helpers.GetBytes(sf.Extra)); err != nil {
		return err
	}
	if _, err := w.Write(helpers.GetBytes(uint16(len(sf.ReplayFrames)))); err != nil {
		return err
	}

	for _, rf := range sf.ReplayFrames  {
		if err := rf.WriteBytes(w); err != nil {
			return err
		}
	}

	if _, err := w.Write(helpers.GetBytes(sf.Action)); err != nil {
		return err
	}
	if _, err := w.Write(sf.ScoreFrame.GetBytes()); err != nil {
		return err
	}

	return nil
}



func (pw *PacketWriter) FellowSpectatorJoined(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerFellowSpectatorJoined

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) FellowSpectatorLeft(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerFellowSpectatorLeft

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) SpectatorJoined(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSpectatorJoined

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) SpectatorLeft(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSpectatorLeft

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) SpectatorCantSpectate(userId int32) {
	p := new(Packet)
	p.PacketId = consts.ServerSpectatorCantSpectate

	p.WriteData(helpers.GetBytes(userId))

	pw.WritePacket(p)
}

func (pw *PacketWriter) SpectateFrames(frame *SpectatorFrame) {
	p := new(Packet)
	p.PacketId = consts.ServerSpectateFrames

	frame.WriteBytes(&p.buffer)

	pw.WritePacket(p)
}
