package warp

type core struct {
	antideuterium      chan int
	deuterium          chan int
	dilithiumAlignment uint
	isBreached         bool

	// Below 15, containment will fail an a warp core breach will occur.
	matterAntimatterContainment uint
}

func (c core) producePlasma() {
	if c.dilithiumAlignment > 100 {
		return
	}
}
