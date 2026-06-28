package component

//codec:gen
type Dye struct {
	Color DyeColor
}

func (*Dye) ID() string {
	return "minecraft:dye"
}
