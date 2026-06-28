package client

//codec:gen
type ClockNetworkState struct {
	TotalTicks  int64 `mc:"VarLong"`
	PartialTick float32
	Rate        float32
}

//codec:gen
type ClockUpdate struct {
	// Clock is the Holder<WorldClock> registry id.
	Clock int32 `mc:"VarInt"`
	State ClockNetworkState
}

//codec:gen
type SetTime struct {
	GameTime     int64
	ClockUpdates []ClockUpdate
}
