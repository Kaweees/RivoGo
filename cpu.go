package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Represents the emulated RISC-V   processor
type CPU struct {
	pc        uint32            // Program counter
	memSize   uint32            // Size of the memory
	registers [REG_COUNT]uint32 // Core registers, exposed publicly to make it easier to interface with
	memory    []uint8           // Memory bus interface
}

// Constructor to initialize memory for the CPU.
func NewCPU(memoryStart uint32, memoryLength uint32) (*CPU, error) {
	cpu := &CPU{}
	cpu.pc = memoryStart
	cpu.memSize = memoryLength
	cpu.memory = make([]uint8, memoryLength)
	cpu.registers[REG_SP] = memoryLength
	return cpu, nil
}

// Displays the contents of the registers
func (cpu *CPU) DisplayRegisters() {
	for i := 0; i < REG_COUNT; {
		fmt.Printf("x%02d: ", i)
		for j := 0; j < 8; j++ {
			fmt.Printf("%08x ", cpu.registers[i])
			if j == 3 {
				fmt.Print(" ")
			}
			i++
		}
		fmt.Println()
	}
	fmt.Printf(" pc: %08x\n", cpu.pc)
}

// Displays the contents of the memory
func (cpu *CPU) DisplayMemory(addr uint32, count uint32) {
	// Pading to align the memory address
	if addr%16 != 0 {
		fmt.Printf("%08x: ", addr)
		for i := uint32(0); i < count; i++ {
			fmt.Printf("%02x ", cpu.memory[addr+i])
		}
	}
	for i := uint32(addr); i < addr+count; i++ {
		if i%16 == 0 && i != 0 {
			fmt.Println()
		}
		if i%16 == 0 {
			fmt.Printf("0x%08x: ", i)
		}
		fmt.Printf("%02x ", cpu.memory[i])
		if i%8 == 0 && i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

// Loads a binary image into memory
func (cpu *CPU) LoadImage(image string) error {
	file, err := os.Open(image)
	if err != nil {
		// Log.Fatalf("Error opening file: %v", err)
		return err
	}

	// Read the size of the binary image
	var binMemSize uint32
	err = binary.Read(file, binary.LittleEndian, &binMemSize)
	if err != nil {
		return fmt.Errorf("error reading binary image size: %v", err)
	}

	// Read the binary image into memory
	maxReadSize := cpu.memSize - binMemSize
	err = binary.Read(file, binary.LittleEndian, cpu.memory[:maxReadSize])
	if err != nil {
		return fmt.Errorf("error reading binary image: %v", err)
	}
	defer file.Close()
	return nil
}

// Read a byte from memory
func (cpu *CPU) FetchByte(addr uint32) (byte, error) {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return 0, fmt.Errorf("invalid address: %d", addr)
	}
	return cpu.memory[addr], nil
}

// Write a byte to memory
func (cpu *CPU) StoreByte(addr uint32, byte uint8) error {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return fmt.Errorf("invalid address: %d", addr)
	}
	cpu.memory[addr] = byte
	return nil
}

// Read a halfword from memory
func (cpu *CPU) FetchHalfWord(addr uint32) (uint16, error) {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return 0, fmt.Errorf("invalid address: %d", addr)
	}
	return binary.LittleEndian.Uint16([]byte(cpu.memory[addr : addr+2])), nil
}

// Write a halfword to memory
func (cpu *CPU) StoreHalfWord(addr uint32, halfWord uint16) error {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return fmt.Errorf("invalid address: %d", addr)
	}
	binary.LittleEndian.PutUint16(cpu.memory[addr:addr+2], halfWord)
	return nil
}

// Read a word from memory
func (cpu *CPU) FetchWord(addr uint32) (uint32, error) {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return 0, fmt.Errorf("invalid address: %d", addr)
	}
	return binary.LittleEndian.Uint32([]byte(cpu.memory[addr : addr+4])), nil
}

// Writes a word to memory
func (cpu *CPU) StoreWord(addr uint32, word uint32) error {
	// Guard against invalid addresses
	if addr >= cpu.memSize {
		return fmt.Errorf("invalid address: %d", addr)
	}
	binary.LittleEndian.PutUint32(cpu.memory[addr:addr+4], word)
	return nil
}

// Fetches the instruction at the current program counter
func (cpu *CPU) Fetch() uint32 {
	// Ignore overflow and wrap around
	instruction, err := cpu.FetchWord(cpu.pc)
	if err != nil {
		Log.Fatalf("Error reading instruction at address %08x: %v", cpu.pc, err)
	}
	cpu.pc += BYTES_PER_WORD
	return instruction
}

// Decodes and executes the instruction given by its opcode
func (cpu *CPU) Execute(instruction uint32) error {
	var opcode InstructionType
	var funct3 uint8

	// Extract the opcode amd funct3 from the instruction
	opcode = InstructionType(instruction & 0x3F)
	funct3 = uint8((instruction >> 12) & 0x3)
	funct7 := uint8((instruction >> 25) & 0x3)

	// Decode the instruction based on the opcode and funct3
	switch opcode {
	case R_TYPE:
		return cpu.ExecuteRType(funct3, funct7, &RTypeInstruction{
			rd:  uint8((instruction >> 7) & 0x3F),
			rs1: uint8((instruction >> 15) & 0x1F),
			rs2: uint8((instruction >> 20) & 0x1F),
		})
	case I_TYPE_ARITH:
		return cpu.ExecuteIArithType(funct3, funct7, &ITypeInstruction{
			rd:  uint8((instruction >> 7) & 0x3F),
			rs1: uint8((instruction >> 15) & 0x1F),
			imm: uint16((instruction >> 20) & 0xFFF),
		})
	case I_TYPE_LOAD:
		return cpu.ExecuteILoadType(funct3, funct7, &ITypeInstruction{
			rd:  uint8((instruction >> 7) & 0x3F),
			rs1: uint8((instruction >> 15) & 0x1F),
			imm: uint16((instruction >> 20) & 0xFFF),
		})
	case I_TYPE_SYS:
		return cpu.ExecuteISysType(funct3, funct7, &ITypeInstruction{
			rd:  uint8((instruction >> 7) & 0x3F),
			rs1: uint8((instruction >> 15) & 0x1F),
			imm: uint16((instruction >> 20) & 0xFFF),
		})
	case S_TYPE:
		return cpu.ExecuteSType(funct3, &STypeInstruction{
			imm: uint16(((instruction >> 25 & 0xEF) << 5) & ((instruction >> 7) & 0xF)),
			rs1: uint8((instruction >> 15) & 0x1F),
			rs2: uint8((instruction >> 20) & 0x1F),
		})
	case B_TYPE:
		return fmt.Errorf("b-type instructions not implemented yet")
		// return cpu.ExecuteBType(funct3, &BTypeInstruction{
		// 	imm: uint16(((instruction >> 31 & 0x1) << 12) & ((instruction >> 25 & 0x3F) << 5) & ((instruction >> 8 & 0xF) << 1) & ((instruction >> 7 & 0x1) << 11)),
		// 	rs1: uint8((instruction >> 15) & 0x1F),
		// 	rs2: uint8((instruction >> 20) & 0x1F),
		// })
	case U_TYPE:
		return cpu.ExecuteUType(funct3, funct7, &UTypeInstruction{
			imm: uint16(instruction >> 12),
			rd:  uint8((instruction >> 7) & 0x3F),
		})
	case J_TYPE:
		return fmt.Errorf("j-type instructions not implemented yet")
	default:
		return fmt.Errorf("unknown instruction type: %v", opcode)
	}
}

// Executes the corresponding R-type instruction based on the funct3 and funct7 fields
func (cpu *CPU) ExecuteRType(funct3 uint8, funct7 uint8, instruction *RTypeInstruction) error {
	if funct3 == 0x0 && funct7 == 0x00 {
		return cpu.ADD(instruction)
	} else if funct3 == 0x0 && funct7 == 0x20 {
		return cpu.SUB(instruction)
	} else if funct3 == 0x4 && funct7 == 0x00 {
		return cpu.XOR(instruction)
	} else if funct3 == 0x6 && funct7 == 0x00 {
		return cpu.OR(instruction)
	} else if funct3 == 0x7 && funct7 == 0x00 {
		return cpu.AND(instruction)
	} else if funct3 == 0x1 && funct7 == 0x00 {
		return cpu.SLL(instruction)
	} else if funct3 == 0x5 && funct7 == 0x00 {
		return cpu.SRL(instruction)
	} else if funct3 == 0x5 && funct7 == 0x20 {
		return cpu.SRA(instruction)
	} else if funct3 == 0x2 && funct7 == 0x00 {
		return cpu.SLT(instruction)
	} else if funct3 == 0x3 && funct7 == 0x00 {
		return cpu.SLTU(instruction)
	} else {
		return fmt.Errorf("unknown r-type instruction: %v", instruction)
	}
}

// Adds two registers and stores the result in a third register
func (cpu *CPU) ADD(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] + cpu.registers[instruction.rs2]
	return nil
}

// Subtracts two registers and stores the result in a third register
func (cpu *CPU) SUB(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] - cpu.registers[instruction.rs2]
	return nil
}

// Bitwise XORs two registers and stores the result in a third register
func (cpu *CPU) XOR(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] ^ cpu.registers[instruction.rs2]
	return nil
}

// Bitwise ORs two registers and stores the result in a third register
func (cpu *CPU) OR(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] | cpu.registers[instruction.rs2]
	return nil
}

// Bitwise ANDs two registers and stores the result in a third register
func (cpu *CPU) AND(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] & cpu.registers[instruction.rs2]
	return nil
}

// Shifts the bits in a register left by a certain amount and stores the result in a third register
func (cpu *CPU) SLL(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] << cpu.registers[instruction.rs2]
	return nil
}

// Shifts the bits in a register right by a certain amount and stores the result in a third register
func (cpu *CPU) SRL(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = cpu.registers[instruction.rs1] >> cpu.registers[instruction.rs2]
	return nil
}

// Shifts the bits in a register right by a certain amount, filling the leftmost bits with the sign bit
func (cpu *CPU) SRA(instruction *RTypeInstruction) error {
	cpu.registers[instruction.rd] = uint32(int32(cpu.registers[instruction.rs1]) >> cpu.registers[instruction.rs2])
	return nil
}

// Sets a register to 1 if the first register is less than the second, 0 otherwise
func (cpu *CPU) SLT(instruction *RTypeInstruction) error {
	if int32(cpu.registers[instruction.rs1]) < int32(cpu.registers[instruction.rs2]) {
		cpu.registers[instruction.rd] = 1
	} else {
		cpu.registers[instruction.rd] = 0
	}
	return nil
}

// Sets a register to 1 if the first register is less than the second, 0 otherwise (unsigned)
func (cpu *CPU) SLTU(instruction *RTypeInstruction) error {
	if uint32(cpu.registers[instruction.rs1]) < uint32(cpu.registers[instruction.rs2]) {
		cpu.registers[instruction.rd] = 1
	} else {
		cpu.registers[instruction.rd] = 0
	}
	return nil
}

// Executes the corresponding I-type arithmetic instruction based on the funct3 field
func (cpu *CPU) ExecuteIArithType(funct3 uint8, funct7 uint8, instruction *ITypeInstruction) error {
	if funct3 == 0x0 {
		return cpu.ADDI(instruction)
	} else if funct3 == 0x4 {
		return cpu.XORI(instruction)
	} else if funct3 == 0x6 {
		return cpu.ORI(instruction)
	} else if funct3 == 0x7 {
		return cpu.ANDI(instruction)
	} else if funct3 == 0x1 && funct7 == 0x00 {
		return cpu.SLLI(instruction)
	} else if funct3 == 0x5 && funct7 == 0x00 {
		return cpu.SRLI(instruction)
	} else if funct3 == 0x5 && funct7 == 0x20 {
		return cpu.SRAI(instruction)
	} else if funct3 == 0x2 {
		return cpu.SLTI(instruction)
	} else if funct3 == 0x3 {
		return cpu.SLTIU(instruction)
	} else {
		return fmt.Errorf("unknown i-type instruction: %v", instruction)
	}
}

func (cpu *CPU) ADDI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) XORI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) ORI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) ANDI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) SLLI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) SRLI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) SRAI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) SLTI(instruction *ITypeInstruction) error {
	return nil
}

func (cpu *CPU) SLTIU(instruction *ITypeInstruction) error {
	return nil
}

// Executes the corresponding I-type load instruction based on the funct3 field
func (cpu *CPU) ExecuteILoadType(funct3 uint8, funct7 uint8, instruction *ITypeInstruction) error {
	return nil
}

// Executes the corresponding I-type system instruction based on the funct3 field
func (cpu *CPU) ExecuteISysType(funct3 uint8, funct7 uint8, instruction *ITypeInstruction) error {
	return nil
}

// Executes the corresponding S-type instruction based on the funct3 field
func (cpu *CPU) ExecuteSType(funct3 uint8, instruction *STypeInstruction) error {
	return nil
}

// Executes the corresponding B-type instruction based on the funct3 field
func (cpu *CPU) ExecuteBType(funct3 uint8, instruction *BTypeInstruction) error {
	return nil
}

// Executes the corresponding U-type instruction based on the funct3 field
func (cpu *CPU) ExecuteUType(funct3 uint8, funct7 uint8, instruction *UTypeInstruction) error {
	return nil
}

// Executes the corresponding J-type instruction based on the funct3 field
func (cpu *CPU) ExecuteJType(funct3 uint8, funct7 uint8, instruction *JTypeInstruction) error {
	return nil
}
