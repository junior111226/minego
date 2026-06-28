package component

import (
	"github.com/KonjacBot/minego/pkg/protocol/slot"
)

//codec:gen
type ChargedProjectiles struct {
	Projectiles []slot.ItemStackTemplate
}

func (*ChargedProjectiles) ID() string {
	return "minecraft:charged_projectiles"
}
