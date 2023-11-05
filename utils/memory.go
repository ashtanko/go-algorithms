package utils

import (
	"runtime"
)

func MeasureMemoryUsage(codeToMeasure func()) uint64 {
	runtime.GC() // Run garbage collection to get a clean slate

	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1) // Read memory statistics before running the code

	codeToMeasure() // Execute the code to be measured

	runtime.GC() // Run garbage collection again

	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2) // Read memory statistics after running the code

	// Calculate the difference in memory usage
	// Note that this calculation might not be 100% accurate due to Go's memory management
	memoryUsage := m2.TotalAlloc - m1.TotalAlloc

	return memoryUsage
}
