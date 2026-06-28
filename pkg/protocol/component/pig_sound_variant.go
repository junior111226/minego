package component

//codec:gen
type PigSoundVariant struct {
	Variant int32 `mc:"VarInt"`
}

func (*PigSoundVariant) ID() string {
	return "minecraft:pig/sound_variant"
}
