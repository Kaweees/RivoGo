package main

// An enum containing all the possible registers of the processor
const (
	REG_ZERO = iota // Hard-wired zero
	REG_RA          // Return address
	REG_SP          // Stack pointer
	REG_GP          // Global pointer
	REG_TP          // Thread pointer
	REG_T0          // Temporary/alternate link register
	REG_T1          // Temporaries
	REG_T2
	REG_S0 // Saved register/frame pointer
	REG_S1 // Saved register
	REG_A0 // Function arguments/return value
	REG_A1
	REG_A2 // Function arguments
	REG_A3
	REG_A4
	REG_A5
	REG_A6
	REG_A7
	REG_S2 // Saved registers
	REG_S3
	REG_S4
	REG_S5
	REG_S6
	REG_S7
	REG_S8
	REG_S9
	REG_S10
	REG_S11
	REG_T3 // Temporaries
	REG_T4
	REG_T5
	REG_T6
	REG_COUNT // Number of registers
)

// RISC-V Constants
const XLEN uint8 = 32 // Width of a register in bits
const BYTES_PER_WORD uint8 = 4
const MEM_MAX_SIZE int = (1 << 32)
const PC_START uint32 = 0x00400000

// Memory-mapped Registers
// const (
// 	MMREG_MTIMECMP = iota

// const REG_MTIMECMP uint32 = 0x02004000
// const REG_MTIME uint32 = 0x0200BFF8
// const REG_MTIME_BASE uint32 = 0x0200B000
// const REG_MTIME_SIZE uint32 = 0x4000

// Memory-mapped I/O
// const MMIO_BASE uint32 = 0x30000000
// const MMIO_SIZE uint32 = 0x1000
