package server

import "github.com/KonjacBot/go-mc/data/packetid"

//codec:gen
type Attack struct {
	EntityID int32 `mc:"VarInt"`
}

func (*Attack) PacketID() packetid.ServerboundPacketID {
	return packetid.ServerboundAttack
}

func init() {
	registerPacket(packetid.ServerboundAttack, func() ServerboundPacket {
		return &Attack{}
	})
}
