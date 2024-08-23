package main

// An enum containing all the possible registers of the processor
const (
	REG_ZERO = iota // Always 0
	REG_AT          // Assembler temporary
	REG_V0          // Return value
	REG_V1
	REG_A0 // Function parameters
	REG_A1
	REG_A2
	REG_A3
	REG_T0 // Function temporary values
	REG_T1
	REG_T2
	REG_T3
	REG_T4
	REG_T5
	REG_T6
	REG_T7
	REG_S0 // Saved registers
	REG_S1
	REG_S2
	REG_S3
	REG_S4
	REG_S5
	REG_S6
	REG_S7
	REG_T8 // Function temporary values
	REG_T9
	REG_K0 // Reserved for interrupt handler
	REG_K1
	REG_GP    // Global pointer
	REG_SP    // Stack pointer
	REG_S8    // Saved registers
	REG_RA    // Return address
	REG_HI    // Multiplication result
	REG_LO    // Division result
	REG_PC    // Program counter
	REG_EPC   // Exception program counter
	REG_COUNT // Number of registers
)

// RISC-V Constants
const XLEN uint8 = 32 // Width of a register in bits
const BYTES_PER_WORD uint8 = 4
const MEM_MAX_SIZE int = (1 << 32)
const PC_START uint32 = 0x00400000
