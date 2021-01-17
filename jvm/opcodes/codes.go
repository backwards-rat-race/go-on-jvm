package opcodes

const (
	ALOAD_0       = 0x2A
	RETURN        = 0xB1
	INVOKESPECIAL = 0xB7
)

func NewALoad(index int) int {
	return ALOAD_0 + index
}
