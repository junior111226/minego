package component

//codec:gen
type CatSoundVariant struct {
	Variant int32 `mc:"VarInt"`
}

func (*CatSoundVariant) ID() string {
	return "minecraft:cat/sound_variant"
}
