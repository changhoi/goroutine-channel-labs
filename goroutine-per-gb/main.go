package main

import (
	"fmt"
	"runtime"
)

func main() {
	ch := make(chan struct{})
	var nch chan struct{}

	cnt := 0
	total := 0

	go func() {
		for {
			<-ch
			printMemUsage(total)
		}
	}()

	fmt.Println("Goroutine start")

	for {
		cnt++
		go func() {
			<-nch // block forever
		}()
		if cnt == 10000 {
			total += cnt
			cnt = 0
			ch<- struct{}{}
		}
	}
}

func printMemUsage(totalCall int) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalCall %d:\tStackInuse = %.3v MiB\tStackSys = %.3v MiB\tTotalAlloc = %.3v MiB\tTotalAlloc = %.3v MiB\tSys = %.3v MiB\n", totalCall, bToMb(m.StackInuse), bToMb(m.StackSys), bToMb(m.HeapAlloc), bToMb(m.TotalAlloc), bToMb(m.Sys))
}

func bToMb(b uint64) float64 {
	return float64(b) / 1024.0 / 1024.0
}
