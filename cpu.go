package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

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

// Represents the emulated RISC-V   processor
type CPU struct {
	registers [REG_COUNT]uint32    // Core registers, exposed publicly to make it easier to interface with
	memory    [MEM_MAX_SIZE]uint32 // Memory bus interface
}

type OpCode uint8

const (
	ADD   OpCode = 0b100000
	ADDU  OpCode = 0b100001
	ADDI  OpCode = 0b001000
	ADDIU OpCode = 0b001001
	AND   OpCode = 0b100100
	ANDI  OpCode = 0b001100
	DIV   OpCode = 0b011010
	DIVU  OpCode = 0b011011
	MULT  OpCode = 0b011000
	MULTU OpCode = 0b011001
	NOR   OpCode = 0b100111
	OR    OpCode = 0b100101
	ORI   OpCode = 0b001101
	SLL   OpCode = 0b000000
	SLLV  OpCode = 0b000100
	SRA   OpCode = 0b000011
	SRAV  OpCode = 0b000111
	SRL   OpCode = 0b000010
	SRLV  OpCode = 0b000110
	SUB   OpCode = 0b100010
	SUBU  OpCode = 0b100011
	XOR   OpCode = 0b100110
	XORI  OpCode = 0b001110
)

// Constructor to initialize memory for the CPU.
func NewCPU() (*CPU, error) {
	cpu := &CPU{}
	cpu.registers[REG_PC] = PC_START
	return cpu, nil
}

// Loads a binary image into memory
func (cpu *CPU) LoadImage(image string) error {
	file, err := os.Open(image)
	if err != nil {
		// Log.Fatalf("Error opening file: %v", err)
		return err
	}

	// Read the size of the binary image
	var memSize uint32
	err = binary.Read(file, binary.LittleEndian, &memSize)
	if err != nil {
		return fmt.Errorf("error reading binary image size: %v", err)
	}

	// Read the binary image into memory
	const maxReadSize = MEM_MAX_SIZE - int(memSize)
	// _, err = file.Read((*(*[MEM_MAX_SIZE]byte)(unsafe.Pointer(&cpu.memory[0])))[:])
	defer file.Close()
	return nil
}

// Decodes the instruction and returns the opcode

// Fetches the instruction at the current program counter
func (cpu *CPU) Fetch() uint32 {
	// Guard against out of bounds memory access
	return cpu.memory[(int(cpu.registers[REG_PC]))%MEM_MAX_SIZE]
}
