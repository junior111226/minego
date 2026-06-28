package server

import "github.com/KonjacBot/go-mc/data/packetid"

//codec:gen
type SpectatorAction struct {
	// SpectateEntityID is an OptionalInt: 0 means no target, otherwise the
	// spectated entity id encoded as id+1.
	SpectateEntityID int32 `mc:"VarInt"`
}

func (*SpectatorAction) PacketID() packetid.ServerboundPacketID {
	return packetid.ServerboundSpectatorAction
}

func init() {
	registerPacket(packetid.ServerboundSpectatorAction, func() ServerboundPacket {
		return &SpectatorAction{}
	})
}
