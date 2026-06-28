package component

import (
	"github.com/KonjacBot/go-mc/net/packet"
)

// As of 26.2 this component is a plain Holder<TrimMaterial>: a VarInt where 0
// means an inline (direct) TrimMaterial follows, otherwise it is a registry
// reference id. This matches the OptID wire format.
//
//codec:gen
type ProvidesTrimMaterial struct {
	Material packet.OptID[TrimMaterial, *TrimMaterial]
}

func (*ProvidesTrimMaterial) ID() string {
	return "minecraft:provides_trim_material"
}
