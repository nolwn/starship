package power

import "github.com/nolwn/starship/hardware/common"

type powerCell struct {
	common.Cell
}

func NewCell(level uint) *powerCell {
	return &powerCell{
		Cell: common.NewCell(level),
	}
}
