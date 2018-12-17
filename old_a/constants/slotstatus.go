package constants

// SlotStatus
const (
	SlotOpen = 1 << iota
	SlotLocked
	SlotNotReady
	SlotReady
	SlotNoMap
	SlotPlaying
	SlotComplete
	SlotQuit
)
