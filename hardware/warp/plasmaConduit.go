package warp

import "github.com/nolwn/starship/hardware/common"

type fuelConduit struct {
	common.Conduit
}

const fuelTransferLimit uint = 100

func NewFuelConduit(source tank, destination chan uint) *fuelConduit {
	return &fuelConduit{
		common.NewConduit(fuelTransferLimit, source, destination),
	}
}
