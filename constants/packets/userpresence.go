package packetconst

type UserPresence struct {
	UserID      int32
	Username    string
	Timezone    int8
	CountryID   int8
	Permissions int8
	Lon         float64
	Lat         float64
	Rank        int32
}
