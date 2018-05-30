package helpers

import "github.com/Gigamons/common/consts"

// HasPrivileges if user has those permissions else not!
func HasPrivileges(p int, u consts.User) bool {
	return u.Privileges&int32(p) > 0
}
