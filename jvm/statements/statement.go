package statements

import (
	"go-on-jvm/jvm/constantpool"
	jvmio "go-on-jvm/jvm/io"
)

type Statement interface {
	NewSerialiser(stack Stack, pool *constantpool.ConstantPool) jvmio.Serialisable
	Variables() []Variable
	fillConstantsPool(pool *constantpool.ConstantPool)
}
