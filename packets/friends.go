package packets

import (
	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/common/consts"
	"github.com/Gigamons/common/helpers"
	"github.com/Gigamons/common/logger"
	"github.com/Gigamons/common/tools/usertools"
	"github.com/Mempler/osubinary"
)

func (w *Writer) SendFriendlist() {
	flist := usertools.GetFriends(w._token.User)
	p := constants.NewPacket(constants.BanchoFriendsList)
	p.SetPacketData(osubinary.IntArray(flist))
	w.WritePacket(p)
}

func AddFriend(u *consts.User, f *consts.User) bool {
	db := helpers.DB
	if u == nil || f == nil {
		return false
	}
	RemoveFriend(u, f)
	_, err := db.Exec("INSERT INTO friends (userid, friendid) VALUES (?, ?)", u.ID, f.ID)
	if err != nil {
		logger.Errorln(err)
		return false
	}
	return true
}

func RemoveFriend(u *consts.User, f *consts.User) bool {
	db := helpers.DB

	if u == nil || f == nil {
		return false
	}

	_, err := db.Exec("DELETE FROM friends WHERE userid = ? AND friendid = ?", u.ID, f.ID)
	if err != nil {
		logger.Errorln(err)
		return false
	}
	return true
}
