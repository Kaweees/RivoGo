# Makefile for compiling, linking, and building the program.
# Begin Variables Section

## Program Section: change these variables based on your program
# The name of the program to build.
TARGET := test

## Compiler Section: change these variables based on your compiler
# -----------------------------------------------------------------------------
# The target architecture.
ARCH=riscv64-unknown-elf
# The compiler executable.
CC := $(ARCH)-gcc
# The compiler flags.
CFLAGS := -ffreestanding -fno-pic -march=rv32i -mabi=ilp32
# The linker executable.
LD := $(ARCH)-ld
# The linker flags.
LDFLAGS := -nostdlib -Wl -Ttext=0x0 -Wl --no-relax -Tdata=0x1000
# The libraries to link.
LDLIBS :=
# Copies and translates object files to another format.
OBJCOPY=$(ARCH)-objcopy
# Displays information about object files.
OBJDUMP=$(ARCH)-objdump
# Displays the size of the object file.
SIZE=$(ARCH)-size
# Creates, modifies, and extracts from archives.
AR=$(ARCH)-ar
# Creates an index (or updates an existing one) for an archive.
RANLIB=$(ARCH)-ranlib

# The shell executable.
SHELL := /bin/bash

## Output Section: change these variables based on your output
# -----------------------------------------------------------------------------
# top directory of project
TOP_DIR := $(shell pwd)
# directory to locate source files
SRC_DIR := $(TOP_DIR)/src
# directory to locate header files
INC_DIR := $(TOP_DIR)/include
# directory to locate object files
OBJ_DIR := $(TOP_DIR)/obj
# directory to place assembly files
ASM_DIR := $(TOP_DIR)/asm
# directory to place build artifacts
BUILD_DIR := $(TOP_DIR)/target/release/

# header files to preprocess
INCS := -I$(INC_DIR)
# source files to compile
SRCS := $(wildcard $(SRC_DIR)/*.c)
# object files to link
OBJS := $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(SRCS))
# assembly files to compile
ASMS := $(patsubst $(SRC_DIR)/%.c, $(ASM_DIR)/%.s, $(SRCS))
# executable file to build
BINS := $(BUILD_DIR)$(TARGET).bin
ELFS := $(BUILD_DIR)$(TARGET).elf
LSTS := $(BUILD_DIR)$(TARGET).lst

## Command Section: change these variables based on your commands
# -----------------------------------------------------------------------------
# Targets
.PHONY: all $(TARGET) clean

# Default target: build the program
all: $(BINS)

# Build the program
$(TARGET): $(BINS)

# Rule to build the program from 
$(BINS): $(ELFS)
	@mkdir -p $(BUILD_DIR) # Create the build directory if it doesn't exist
	$(OBJCOPY) $(ELFS) -O binary $(BINS)
	$(OBJDUMP) -D $(ELFS) > $(LSTS)

# Rule to build the program from linked object files
$(ELFS): $(OBJ_DIR)/%.o
	@mkdir -p $(BUILD_DIR) # Create the build directory if it doesn't exist
	$(LD) $(LDFLAGS) $(OBJS) -o $(ELFS) $(LDLIBS)

# Rule to compile source files into object files 
$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	@mkdir -p $(OBJ_DIR) # Create the object directory if it doesn't exist
	$(CC) $(CFLAGS) $(INCS) -c $< -o $@

# Rule to compile source files into object files 
$(OBJ_DIR)/%.o: $(SRC_DIR)/%.S
	@mkdir -p $(OBJ_DIR) # Create the object directory if it doesn't exist
	$(CC) $(CFLAGS) $(INCS) -c $< -o $@

# Clean target: remove build artifacts and non-essential files
clean:
	@echo "Cleaning $(TARGET)..."
	rm -rf $(OBJ_DIR) $(BUILD_DIR)
