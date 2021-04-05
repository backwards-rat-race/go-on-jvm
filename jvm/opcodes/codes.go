package opcodes

import jvmio "go-on-jvm/jvm/io"

const (
	ICONST_M1     = 0x02
	ICONST_0      = 0x03
	ICONST_1      = 0x04
	ICONST_2      = 0x05
	ICONST_3      = 0x06
	ICONST_4      = 0x07
	ICONST_5      = 0x08
	LDC           = 0x12
	ALOAD         = 0x19
	ALOAD_0       = 0x2A
	ASTORE        = 0x3A
	ASTORE_0      = 0x4B
	IADD          = 0x60
	LADD          = 0x61
	FADD          = 0x62
	DADD          = 0x63
	ISUB          = 0x64
	LSUB          = 0x65
	FSUB          = 0x66
	DSUB          = 0x67
	IMUL          = 0x68
	LMUL          = 0x69
	FMUL          = 0x6A
	DMUL          = 0x6B
	IFEQ          = 0x99
	IFNE          = 0x9A
	IFLT          = 0x9B
	IFGE          = 0x9C
	IFGT          = 0x9D
	IFLE          = 0x9E
	GOTO          = 0xA7
	IRETURN       = 0xAC
	LRETURN       = 0xAD
	FRETURN       = 0xAE
	DRETURN       = 0xAF
	ARETURN       = 0xB0
	RETURN        = 0xB1
	INVOKESPECIAL = 0xB7
	INVOKESTATIC  = 0xB8
	ARRAYLENGTH   = 0xBE
	IFNULL        = 0xC6
	IFNONNULL     = 0xC7
	GOTOW         = 0xC8

	GotoSize = 3
	IfSize   = 3

	maxGotoSize = 65535
)

func GetALoadInstruction(index int) []byte {
	if index > 3 {
		return []byte{byte(ALOAD), byte(index)}
	} else {
		return []byte{byte(ALOAD_0 + index)}
	}
}

func GetAStoreInstruction(index int) []byte {
	if index > 3 {
		return []byte{byte(ASTORE), byte(index)}
	} else {
		return []byte{byte(ASTORE_0 + index)}
	}
}

func GetGotoInstruction(line int) []byte {
	if line > maxGotoSize {
		panic("wide goto not yet supported. max method size reached")
	} else {
		instruction := []byte{byte(GOTO)}
		instruction = jvmio.AppendPaddedBytes(instruction, line, 2)
		return instruction
	}
}
