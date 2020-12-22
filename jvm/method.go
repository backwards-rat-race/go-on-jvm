package jvm

import (
	"go-on-jvm/jvm/constants"
	"io"
)

type Method struct {
	Name   string
	Access []AccessModifier
}

func (m Method) Compile(w io.Writer) error {
	return nil
}

func (m Method) fillConstantsPool(pool *constants.ConstantPool) {

}
