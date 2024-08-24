package main

// Represents an instruction type in RISC-V
type InstructionType uint8

// An enum containing all the possible formats of an instruction
const (
	R_TYPE       InstructionType = 0b0110011 // Register (R-format) instructions
	I_TYPE_ARITH InstructionType = 0b0010011 // Arithmetic Immediate (I-format) instructions
	I_TYPE_LOAD  InstructionType = 0b0000011 // Load Immediate (I-format) instructions
	I_TYPE_SYS   InstructionType = 0b1110011 // System Immediate (I-format) instructions
	S_TYPE       InstructionType = 0b0100011 // Store (S-format) instructions
	B_TYPE       InstructionType = 0b1100011 // Branch (B-format) instructions
	U_TYPE       InstructionType = 0b0110111 // Upper immediate (U-format) instructions
	J_TYPE       InstructionType = 0b1101111 // Jump (J-format) instructions
)

// Represents a MIPS assembly instruction
type AssemblyInstruction struct {
	instructionType InstructionType   // The type of instruction
	rType           *RTypeInstruction // The R-type instruction
	iType           *ITypeInstruction // The I-type instruction
	sType           *STypeInstruction // The S-type instruction
	bType           *BTypeInstruction // The B-type instruction
	uType           *UTypeInstruction // The U-type instruction
	jType           *JTypeInstruction // The J-type instruction
}

// Represents a R-type instruction
type RTypeInstruction struct {
	rd  uint8 // The destination register
	rs1 uint8 // The first source register
	rs2 uint8 // The second source register
}

// Represents a I-type instruction
type ITypeInstruction struct {
	rd  uint8  // The destination register
	rs1 uint8  // The first source register
	imm uint16 // The immediate value
}

// Represents a S-type instruction
type STypeInstruction struct {
	imm uint16 // The immediate value
	rs1 uint8  // The first source register
	rs2 uint8  // The second source register
}

// Represents a B-type instruction
type BTypeInstruction struct {
	imm uint16 // The immediate value
	rs1 uint8  // The first source register
	rs2 uint8  // The second source register
}

// Represents a U-type instruction
type UTypeInstruction struct {
	imm uint16 // The immediate value
	rd  uint8  // The destination register
}

// Represents a J-type instruction
type JTypeInstruction struct {
	imm uint16 // The immediate value
	rd  uint8  // The destination register
}
