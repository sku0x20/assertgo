package pkg

type AssertionSuite struct {
	sink *TSink
}

func NewAssertionSuite(sink *TSink) *AssertionSuite {
	return &AssertionSuite{sink: sink}
}
