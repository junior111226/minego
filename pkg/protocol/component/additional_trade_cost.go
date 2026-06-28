package component

//codec:gen
type AdditionalTradeCost struct {
	Cost int32 `mc:"VarInt"`
}

func (*AdditionalTradeCost) ID() string {
	return "minecraft:additional_trade_cost"
}
