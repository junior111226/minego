package component

//codec:gen
type AttackRange struct {
	MinReach float32
	MaxReach float32

	MinCreativeReach float32
	MaxCreativeReach float32

	HitboxMargin float32

	MobFactor float32
}

func (*AttackRange) ID() string {
	return "minecraft:attack_range"
}
