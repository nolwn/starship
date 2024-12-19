package internalsensors

type section struct {
	hull   int
	number int
}

func NewHullIntegrity(number int) *section {
	return &section{
		hull:   100,
		number: number,
	}
}

func (s *section) Value() int {
	return s.number
}
