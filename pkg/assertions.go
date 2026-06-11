package pkg

type Assertions struct {
	sink *TSink
}

func NewAssertions(sink *TSink) *Assertions {
	return &Assertions{sink: sink}
}
