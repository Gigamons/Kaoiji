package helpers

import "github.com/Gigamons/Kaoiji/constants"

// HasPrivileges if user has those permissions else not!
func HasPrivileges(p int, u constants.User) bool {
	return u.Privileges&int32(p) > 0
}
