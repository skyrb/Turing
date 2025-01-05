
package main

import (
    "fmt"
)

// Computer represents a complex computer configuration
type Computer struct {
    CPU         string
    GPU         string
    Memory      int
    Storage     string
    OperatingSystem string
    Monitor     string
}

// ComputerBuilder is a builder for the Computer object
type ComputerBuilder struct {
    computer *Computer
}

// NewComputerBuilder creates a new ComputerBuilder
func NewComputerBuilder() *ComputerBuilder {
    return &ComputerBuilder{computer: &Computer{}}
}

// WithCPU sets the CPU of the computer
func (b *ComputerBuilder) WithCPU(cpu string) *ComputerBuilder {
    b.computer.CPU = cpu
    return b
}

// WithGPU sets the GPU of the computer
func (b *ComputerBuilder) WithGPU(gpu string) *ComputerBuilder {
    b.computer.GPU = gpu
    return b
}

// WithMemory sets the memory of the computer in GB
func (b *ComputerBuilder) WithMemory(memory int) *ComputerBuilder {
    b.computer.Memory = memory
    return b
}

// WithStorage sets the storage of the computer
func (b *ComputerBuilder) WithStorage(storage string) *ComputerBuilder {
    b.computer.Storage = storage
    return b
}

// WithOperatingSystem sets the operating system of the computer
func (b *ComputerBuilder) WithOperatingSystem(os string) *ComputerBuilder {
    b.computer.OperatingSystem = os
    return b
}

// WithMonitor sets the monitor of the computer
func (b *ComputerBuilder) WithMonitor(monitor string) *ComputerBuilder {
    b.computer.Monitor = monitor
    return b
}

// Build constructs and returns the Computer object
func (b *ComputerBuilder) Build() *Computer {
    return b.computer
}

func main() {
    // Various ways to build computers using the builder pattern

    // Building a basic computer with 8GB RAM and a 256GB SSD
    basicComputer := NewComputerBuilder().
        WithMemory(8).
        WithStorage("256GB SSD").
        Build()

    fmt.Printf("Basic Computer: %+v\n", basicComputer)

    // Building a high-end gaming computer with a powerful CPU, GPU, lots of RAM, and a big screen monitor
    highEndGamingComputer := NewComputerBuilder().
        WithCPU("Intel Core i9-10900K").
        WithGPU("NVIDIA GeForce RTX 3080 Ti").
        WithMemory(32).
        WithStorage("2TB NVMe SSD").
        WithOperatingSystem("Windows 10").
        WithMonitor("ASUS ROG Swift PG279QM").
        Build()

    fmt.Printf("High-End Gaming Computer: %+v\n", highEndGamingComputer)

    // Building a lightweight notebook with basic specifications
    lightweightNotebook := NewComputerBuilder().
        WithCPU("Intel Core i5-1135G7").
        WithGPU("Intel Iris Xe Graphics").
        WithMemory(8).
        WithStorage("512GB PCIe NVMe M.2 SSD").
        WithOperatingSystem("Windows 11 Home").
        Build()

    fmt.Printf("Lightweight Notebook: %+v\n", lightweightNotebook)

    // Building a headless server for high-performance computation
    headlessServer := NewComputerBuilder().
        WithCPU("Intel Xeon E-2278G").
        WithGPU("NVIDIA Tesla V100-SXM3-32GB").
        WithMemory(128).
        WithStorage("2TB NVMe RAID 0 Array").
        WithOperatingSystem("Ubuntu Server 20.04 LTS").
        Build()
    
    fmt.Printf("Headless Server: %+v\n", headlessServer)
}  
