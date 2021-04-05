package statements

import (
	"go-on-jvm/jvm/constantpool"
)

type Statement interface {
	GetInstructions(stack *Stack, pool *constantpool.ConstantPool) []byte
	FillConstantsPool(pool *constantpool.ConstantPool)
	MaxStack() int
}
