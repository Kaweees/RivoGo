package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alexflint/go-arg"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

type argsParsed struct {
	args
}

// Used to parse hexadecimal/decimal arguments as uint32
type HexUint uint32

func (h *HexUint) UnmarshalText(b []byte) error {
	input := string(b)
	var v uint64
	var err error

	// Check if the input string starts with "0x" or "0X" for hexadecimal
	if strings.HasPrefix(input, "0x") || strings.HasPrefix(input, "0X") {
		v, err = strconv.ParseUint(input, 0, 32)
	} else {
		// Assume decimal if not hexadecimal
		v, err = strconv.ParseUint(input, 10, 32)
	}

	if err != nil {
		return err
	}

	*h = HexUint(v)
	return nil
}

func (args) Epilogue() string {
	return "For more information visit github.com/Kaweees/RivoGo"
}

// CLI arguments
type args struct {
	// File config
	FileName string `arg:"required" help:"Image file to virtualize"`
	// Logging config
	Logging bool `arg:"-l,--logging" help:"Enable logging"`
	// Starting address
	Start HexUint `arg:"help:Program counter starting address"`
	// Memory length
	Length HexUint `arg:"-n,--length" help:"Memory length"`
}

// Returns a human-readable version string
func (args) Version() string {
	return fmt.Sprintf("Version: %v, commit: %v, built at: %v", version, commit, date)
}

// Returns a description of the program
func (args) Description() string {
	return "A simple assembler for the RISC-V architecture"
}

// Returns the parsed CLI arguments
func GetCliArgs() (cli argsParsed, err error) {
	rawCli := args{
		Logging: false,
		Start:   HexUint(PC_START),
		Length:  HexUint(MEM_MAX_SIZE),
	}

	arg.MustParse(&rawCli)
	cli.args = rawCli
	fmt.Printf("Parsed value: %d (hex: %x)\n", cli.Start, uint32(cli.Start))

	return cli, nil
}
