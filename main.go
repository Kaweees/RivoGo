package main

func main() {
	var err error
	var cli argsParsed

	// Parse the arguments
	cli, _ = GetCliArgs()

	// Initialize the logger
	initalizeLogger()

	var cpu *CPU

	// Initialize the CPU
	cpu, err = NewCPU(uint32(cli.Start), uint32(cli.Length))
	if err != nil {
		return
	}

	// Load the image into memory
	cpu.LoadImage(cli.FileName)
	cpu.DisplayRegisters()
	cpu.DisplayMemory(cpu.pc, 200)
	var running = true
	for running {
		// Fetch the instruction
		instruction := cpu.Fetch()
		// Execute the instruction
		err = cpu.Execute(instruction)
		if err != nil {
			return
		}
	}
}
