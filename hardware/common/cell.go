package common

type Cell interface {
	Level() uint
}

type cell struct {
	level uint
}

func (c *cell) Level() uint {
	return c.level
}

func NewCell(level uint) *cell {
	return &cell{
		level,
	}
}
