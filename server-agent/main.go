package main

import (
	"fmt"
	cpu "server-agent/core/CPU"
)

func main() {
	fmt.Println(cpu.GetCPUInfo())
}
