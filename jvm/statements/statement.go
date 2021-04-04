package statements

import (
	"go-on-jvm/jvm/constantpool"
)

type Statement interface {
	GetInstructions(writeIndex int, stack *Stack, pool *constantpool.ConstantPool) []byte
	FillConstantsPool(pool *constantpool.ConstantPool)
}
