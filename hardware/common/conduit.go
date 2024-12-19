package common

type Conduit interface {
	Transfer(amount uint)
}

type conduit struct {
	destination   chan uint
	source        Cell
	transferLimit uint
}

func NewConduit(
	transferLimit uint,
	source Cell,
	destination chan uint,
) *conduit {
	return &conduit{
		destination,
		source,
		transferLimit,
	}
}

func (c *conduit) Transfer(amount uint) {
	source := c.source.(*cell)

	// The amount cannot exceed the transfer limit. If it does, set it to be the transfer
	// limit.
	if amount > c.transferLimit {
		amount = c.transferLimit
	}

	// The amount cannot exceed the available energy level. If it does, set it to be the
	// available amount.
	if amount > source.level {
		amount = source.level
	}

	// Remove amount from the source's energy level and feed it into the destination.
	source.level -= amount
	c.destination <- amount
}
