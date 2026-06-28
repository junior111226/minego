package component

//codec:gen
type CowSoundVariant struct {
	Variant int32 `mc:"VarInt"`
}

func (*CowSoundVariant) ID() string {
	return "minecraft:cow/sound_variant"
}
