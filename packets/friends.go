package packets

import (
	"fmt"

	"github.com/Gigamons/Kaoiji/constants"
	"github.com/Gigamons/Kaoiji/global"
	"github.com/Gigamons/Kaoiji/tools/usertools"
)

func (w *Writer) SendFriendlist() {
	flist := usertools.GetFriends(&w._token.User)
	p := NewPacket(constants.BanchoFriendsList)
	p.SetPacketData(IntArray(flist))
	w.Write(p.ToByteArray())
}

func AddFriend(u *constants.User, f *constants.User) bool {
	db := global.DB
	if u == nil || f == nil {
		return false
	}
	RemoveFriend(u, f)
	_, err := db.Exec("INSERT INTO friends (userid, friendid) VALUES (?, ?)", u.ID, f.ID)
	if err != nil {
		fmt.Println(err)
	}
	return true
}

func RemoveFriend(u *constants.User, f *constants.User) bool {
	db := global.DB

	if u == nil || f == nil {
		return false
	}

	_, err := db.Exec("DELETE FROM friends WHERE userid = ? AND friendid = ?", u.ID, f.ID)
	if err != nil {
		fmt.Println(err)
	}
	return true
}
