package particle

import (
	"io"

	pk "github.com/KonjacBot/go-mc/net/packet"
)

type ParticleData interface {
	io.ReaderFrom
	io.WriterTo
	ParticleID() int32
}

type Particle struct {
	ID   int32
	Data ParticleData
}

func (p *Particle) ReadFrom(r io.Reader) (int64, error) {
	n, err := (*pk.VarInt)(&p.ID).ReadFrom(r)
	if err != nil {
		return n, err
	}

	switch p.ID {
	case 1: // block
		data := &Block{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 2: // block_marker
		data := &BlockMarker{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 7, 10: // geyser, geyser_plume
		data := &Geyser{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 8, 9: // geyser_base, geyser_poof
		data := &GeyserBase{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 15: // dragon_breath
		data := &DragonBreth{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 21: // dust
		data := &Dust{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 22: // dust_color_transition
		data := &DustColorTransition{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 23: // effect
		data := &Effect{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 28: // entity_effect
		data := &EntityEffect{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 36: // falling_dust
		data := &FallingDust{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 43: // tinted_leaves
		data := &TintedLeaves{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 45: // sculk_charge
		data := &SculkCharge{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 49: // flash
		data := &Flash{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 53: // instant_effect
		data := &InstantEffect{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 54: // item
		data := &Item{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 55: // vibration
		data := &Vibration{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 56: // trail
		data := &Trail{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 112: // shriek
		data := &Shriek{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 118: // dust_pillar
		data := &DustPillar{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	case 122: // block_crumble
		data := &BlockCrumble{}
		n2, err := data.ReadFrom(r)
		p.Data = data
		return n + n2, err
	default:
		// BasicParticle - no additional data
		p.Data = nil
		return n, nil
	}
}

func (p Particle) WriteTo(w io.Writer) (int64, error) {
	n, err := (*pk.VarInt)(&p.ID).WriteTo(w)
	if err != nil {
		return n, err
	}

	if p.Data == nil {
		return n, nil
	}

	n2, err := p.Data.WriteTo(w)
	return n + n2, err
}
