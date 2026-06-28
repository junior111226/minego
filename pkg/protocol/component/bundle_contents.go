package component

import (
	"github.com/KonjacBot/minego/pkg/protocol/slot"
)

//codec:gen
type BundleContents struct {
	Items []slot.ItemStackTemplate
}

func (*BundleContents) ID() string {
	return "minecraft:bundle_contents"
}
