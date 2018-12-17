package consts

type PacketId uint16

// PacketID list by Itsyuka https://github.com/Itsyuka/osu-packet under MIT License
const (
	ClientSendUserStatus             PacketId = iota
	ClientSendIrcMessage
	ClientExit
	ClientRequestStatusUpdate
	ClientPong
	ServerLoginReply
	ServerCommandError
	ServerSendMessage
	ServerPing
	ServerHandleIrcChangeUsername
	ServerHandleIrcQuit
	ServerHandleOsuUpdate
	ServerHandleUserQuit
	ServerSpectatorJoined
	ServerSpectatorLeft
	ServerSpectateFrames
	ClientStartSpectating
	ClientStopSpectating
	ClientSpectateFrames
	ServerVersionUpdate
	ClientErrorReport
	ClientCantSpectate
	ServerSpectatorCantSpectate
	ServerGetAttention
	ServerAnnounce
	ClientSendIrcMessagePrivate
	ServerMatchUpdate
	ServerMatchNew
	ServerMatchDisband
	ClientLobbyPart
	ClientLobbyJoin
	ClientMatchCreate
	ClientMatchJoin
	ClientMatchPart
	ServerMatchJoinSuccess
	ServerMatchJoinFail
	ClientMatchChangeSlot
	ClientMatchReady
	ClientMatchLock
	ClientMatchChangeSettings
	ServerFellowSpectatorJoined
	ServerFellowSpectatorLeft
	ClientMatchStart
	ServerMatchStart
	ClientMatchScoreUpdate
	ServerMatchScoreUpdate
	ClientMatchComplete
	ServerMatchTransferHost
	ClientMatchChangeMods
	ClientMatchLoadComplete
	ServerMatchAllPlayersLoaded
	ClientMatchNoBeatmap
	ClientMatchNotReady
	ClientMatchFailed
	ServerMatchPlayerFailed
	ServerMatchComplete
	ClientMatchHasBeatmap
	ClientMatchSkipRequest
	ServerMatchSkip
	ServerUnauthorised
	ClientChannelJoin
	ServerChannelJoinSuccess
	ServerChannelAvailable
	ServerChannelRevoked
	ServerChannelAvailableAutojoin
	ClientBeatmapInfoRequest
	ServerBeatmapInfoReply
	ClientMatchTransferHost
	ServerLoginPermissions
	ServerFriendsList
	ClientFriendAdd
	ClientFriendRemove
	ServerProtocolNegotiation
	ServerTitleUpdate
	ClientMatchChangeTeam
	ClientChannelLeave
	ClientReceiveUpdates
	ServerMonitor
	ServerMatchPlayerSkipped
	ClientSetIrcAwayMessage
	ServerUserPresence
	ClientUserStatsRequest
	ServerRestart
	ClientInvite
	ServerInvite
	ServerChannelListingComplete
	ClientMatchChangePassword
	ServerMatchChangePassword
	ServerBanInfo
	ClientSpecialMatchInfoRequest
	ServerUserSilenced
	ServerUserPresenceSingle
	ServerUserPresenceBundle
	ClientUserPresenceRequest
	ClientUserPresenceRequestAll
	ClientUserToggleBlockNonFriendPM
	ServerUserPMBlocked
	ServerTargetIsSilenced
	ServerVersionUpdateForced
	ServerSwitchServer
	ServerAccountRestricted
	ServerRTX
	ClientMatchAbort
	ServerSwitchTourneyServer
	ClientSpecialJoinMatchChannel
	ClientSpecialLeaveMatchChannel
)