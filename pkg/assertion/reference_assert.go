package assertion

import (
	"fmt"

	"github.com/sku0x20/assertgo/pkg/sink"
)

type ReferenceAssert struct {
	sink  *sink.TSink
	value any
	other any
}

func NewReferenceAssert(s *sink.TSink, value, other any) *ReferenceAssert {
	return &ReferenceAssert{sink: s, value: value, other: other}
}

func (r *ReferenceAssert) IsSameAs() {
	if r.value != r.other {
		r.sink.Fail(fmt.Sprintf("expected same reference as %v but got %v", r.other, r.value))
	}
}

func (r *ReferenceAssert) IsNotSameAs() {
	if r.value == r.other {
		r.sink.Fail(fmt.Sprintf("expected different reference but got same %v", r.value))
	}
}
