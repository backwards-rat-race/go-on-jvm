package opcodes

const (
	ICONST_0      = 0x03
	IADD          = 0x60
	ALOAD_0       = 0x2A
	ASTORE_0      = 0x3A
	IRETURN       = 0xAC
	LRETURN       = 0xAD
	FRETURN       = 0xAE
	DRETURN       = 0xAD
	ARETURN       = 0xB0
	RETURN        = 0xB1
	INVOKESPECIAL = 0xB7
	INVOKESTATIC  = 0xB8
	ARRAYLENGTH   = 0xBE
	IFNONNULL     = 0xC7
)

func NewIConst(i int) int {
	return ICONST_0 + i
}

func NewALoad(index int) int {
	return ALOAD_0 + index
}

func NewAStore(index int) int {
	return ASTORE_0 + index
}
