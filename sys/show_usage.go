package sys

import (
	"fmt"
	"runtime"
)

type MemUsage struct {
	Alloc      float64
	TotalAlloc float64
	Sys        float64
	NumGC      uint32
}

func (m MemUsage) String() string {
	return fmt.Sprintf("Alloc = %.2fMB TotalAlloc = %.2fMB Sys = %.2fMB NumGC = %v",
		m.Alloc,
		m.TotalAlloc,
		m.Sys,
		m.NumGC)
}

func ShowMemUsage() MemUsage {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)
	mb := 1024 * 1024.0

	return MemUsage{
		Alloc:      float64(m.Alloc) / mb,
		TotalAlloc: float64(m.TotalAlloc) / mb,
		Sys:        float64(m.Sys) / mb,
		NumGC:      m.NumGC,
	}
}
