package component

import (
	pk "github.com/KonjacBot/go-mc/net/packet"
	"github.com/KonjacBot/minego/pkg/protocol/slot"
)

//codec:gen
type Container struct {
	Items []pk.Option[slot.ItemStackTemplate, *slot.ItemStackTemplate]
}

func (*Container) ID() string {
	return "minecraft:container"
}
