# Makefile for compiling, linking, and building the program.
# Begin Variables Section

## Program Section: change these variables based on your program
# The name of the program to build.
TARGET := freestanding

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
LDFLAGS := -nostdlib -Wl -Ttext=0x0 -Wl --no-relax -Tdata=0x1400
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
# directory to place build artifacts
BUILD_DIR := $(TOP_DIR)/target/release/

# header files to preprocess
INCS := -I$(INC_DIR)
# source files to compile
SRCS := $(wildcard $(SRC_DIR)/*.c)
# object files to link
OBJS := $(patsubst $(SRC_DIR)/%.c, $(OBJ_DIR)/%.o, $(SRCS))
# executable file to build
BINS := $(BUILD_DIR)$(TARGET).bin

## Command Section: change these variables based on your commands
# -----------------------------------------------------------------------------
# Targets
.PHONY: all $(TARGET) debug

# Default target: build the program
all: $(BINS)

# Build the program
$(TARGET): $(BINS)

# Rule to build the program from linked object files
$(BINS): $(OBJS)
	@mkdir -p $(BUILD_DIR) # Create the build directory if it doesn't exist
	$(LD) $(LDFLAGS) $(OBJS) -o $(BINS)

# Rule to compile source files into object files 
$(OBJ_DIR)/%.o: $(SRC_DIR)/%.c
	@mkdir -p $(OBJ_DIR) # Create the object directory if it doesn't exist
	$(CC) $(CFLAGS) $(INCS) -c $< -o $@

# Clean target: remove build artifacts and non-essential files
clean:
	@echo "Cleaning $(TARGET)..."
	rm -rf $(OBJ_DIR) $(BUILD_DIR)
