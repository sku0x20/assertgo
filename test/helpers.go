package test

import "github.com/sku0x20/assertgo/pkg/sink"

func NewSink() (*MockT, *sink.TSink) {
	mock := &MockT{}
	return mock, sink.NewTSink(mock)
}
