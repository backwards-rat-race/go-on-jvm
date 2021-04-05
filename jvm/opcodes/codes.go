package opcodes

import jvmio "go-on-jvm/jvm/io"

const (
	NOP           = 0x00
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
	ALOAD_1       = 0x2B
	ALOAD_2       = 0x2C
	ALOAD_3       = 0x2D
	AALOAD        = 0x32
	ASTORE        = 0x3A
	ASTORE_0      = 0x4B
	ASTORE_1      = 0x4C
	ASTORE_2      = 0x4D
	ASTORE_3      = 0x4E
	AASTORE       = 0x53
	POP           = 0x57
	POP2          = 0x58
	DUP           = 0x59
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
	GETSTATIC     = 0xB2
	PUTSTATIC     = 0xB3
	GETFIELD      = 0xB4
	PUTFIELD      = 0xB5
	INVOKEVIRTUAL = 0xB6
	INVOKESPECIAL = 0xB7
	INVOKESTATIC  = 0xB8
	NEW           = 0xBB
	ARRAYLENGTH   = 0xBE
	IFNULL        = 0xC6
	IFNONNULL     = 0xC7
	GOTOW         = 0xC8

	GotoSize = 3
	IfSize   = 3
)

func GetALoadInstruction(index int) []byte {
	switch index {
	case 0:
		return []byte{ALOAD_0}
	case 1:
		return []byte{ALOAD_1}
	case 2:
		return []byte{ALOAD_2}
	case 3:
		return []byte{ALOAD_3}
	default:
		return []byte{byte(ALOAD), byte(index)}
	}
}

func GetAStoreInstruction(index int) []byte {
	switch index {
	case 0:
		return []byte{ASTORE_0}
	case 1:
		return []byte{ASTORE_1}
	case 2:
		return []byte{ASTORE_2}
	case 3:
		return []byte{ASTORE_3}
	default:
		return []byte{byte(ASTORE), byte(index)}
	}
}

func GetGotoInstruction(line uint16) []byte {
	instruction := []byte{byte(GOTO)}
	instruction = jvmio.AppendPaddedBytes(instruction, uint(line), 2)
	return instruction
}

func GetGotoInstructionI(line int) []byte {
	return GetGotoInstruction(uint16(line))
}
