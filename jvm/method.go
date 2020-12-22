package jvm

import "io"

type Method struct {
	Name   string
	Access AccessModifier
}

func (m Method) Compile(io.Writer) error {
	panic("implement me")
}
