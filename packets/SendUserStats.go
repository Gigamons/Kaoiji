package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Mempler/osubinary"
)

// SendUserStats sends the UserStatus to the writer.
func (w *Writer) SendUserStats(x *constants.UserStatsStruct) {
	p := constants.NewPacket(constants.BanchoHandleOsuUpdate)
	p.SetPacketData(osubinary.Marshal(x))
	w.Write(p)
}
