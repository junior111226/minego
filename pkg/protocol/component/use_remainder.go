package component

import (
	"github.com/KonjacBot/minego/pkg/protocol/slot"
)

//codec:gen
type UseRemainder struct {
	Remainder slot.ItemStackTemplate
}

func (*UseRemainder) ID() string {
	return "minecraft:use_remainder"
}
