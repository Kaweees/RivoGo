package main

import "fmt"

func main() {
	var err error
	var cli argsParsed

	// Parse the arguments
	// cli, _ = GetCliArgs()

	// Initialize the logger
	initalizeLogger()

	var cpu *CPU

	// Initialize the CPU
	cpu, err = NewCPU()
	if err != nil {
		return
	}

	// Load the image into memory
	cpu.LoadImage(cli.FileName)

	fmt.Println("CPU initialized")
	var running = true
	for running {
		// Fetch the instruction
		instruction := cpu.Fetch()
		fmt.Println(instruction)
		// Execute the instruction
		err = cpu.Execute(instruction)
		if err != nil {
			return
		}
	}

}
