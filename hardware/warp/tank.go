package warp

import "github.com/nolwn/starship/hardware/common"

const (
	Antideuterium = iota
	Deuterium
)

type tank struct {
	common.Cell
	contents int
}

func (t tank) Level() uint {
	return t.Cell.Level()
}

func (t tank) Contents() int {
	return t.contents
}
