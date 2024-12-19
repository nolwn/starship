package power

import "github.com/nolwn/starship/hardware/common"

const transferLimit uint = 100

type conduit struct {
	common.Conduit
}

func NewConduit(source powerCell, destination chan uint) *conduit {
	return &conduit{
		Conduit: common.NewConduit(transferLimit, source, destination),
	}
}
