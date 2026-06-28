package component

//codec:gen
type ChickenSoundVariant struct {
	Variant int32 `mc:"VarInt"`
}

func (*ChickenSoundVariant) ID() string {
	return "minecraft:chicken/sound_variant"
}
