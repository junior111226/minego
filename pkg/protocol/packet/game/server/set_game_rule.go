package server

import "github.com/KonjacBot/go-mc/data/packetid"

//codec:gen
type SetGameRule struct {
	Entries []SetGameRuleEntry
}

//codec:gen
type SetGameRuleEntry struct {
	GameRuleKey string `mc:"Identifier"`
	Value       string
}

func (*SetGameRule) PacketID() packetid.ServerboundPacketID {
	return packetid.ServerboundSetGameRule
}

func init() {
	registerPacket(packetid.ServerboundSetGameRule, func() ServerboundPacket {
		return &SetGameRule{}
	})
}
