package main

import (
	"os"
	"runtime/pprof"
)

func main() {
	// Open a file to write the CPU profile.
	f, err := os.Create("cpu.prof")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Start CPU profiling.
	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	// Rest of the code...
}