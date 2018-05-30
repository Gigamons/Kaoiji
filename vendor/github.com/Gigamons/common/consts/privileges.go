package consts

// User Privileges.
const (
	Supporter         = 2 << 0
	BAT               = 2 << 1
	TournamentStaff   = 2 << 2
	AdminPannelAccess = 2 << 3
	AdminManageUsers  = 2 << 4
	AdminBanUsers     = 2 << 5
	AdminSilenceUsers = 2 << 6
	AdminWipeUsers    = 2 << 7
	AdminBeatmaps     = 2 << 8
	AdminDeveloper    = 2 << 9
	AdminSettings     = 2 << 10
	AdminReports      = 2 << 11
	AdminPrivileges   = 2 << 12
	AdminSendAlerts   = 2 << 13
	AdminChatMod      = 2 << 14
	AdminKickUsers    = 2 << 15
)
