package cpu

import (
	"bufio"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	UNKNOW = "UNKNOW"
)

type CPUInfo struct {
	Arch  string
	Model string
}

func GetCPUInfo() CPUInfo {
	return CPUInfo{
		Arch:  runtime.GOARCH,
		Model: GetCPUModel(),
	}
}

func GetCPUModel() string {
	var CPUModel string
	switch runtime.GOOS {
	case "darwin":
		CPUModel = getDarwinCPUModel()
	case "linux":
		CPUModel = getLinuxCPUModel()
	case "windows":
		CPUModel = getWindowsCPUModel()
	}
	return CPUModel
}

func getDarwinCPUModel() string {
	cmd := exec.Command("sysctl", "-n", "machdep.cpu.brand_string")
	output, err := cmd.Output()
	if err != nil {
		return UNKNOW
	}
	return strings.TrimSpace(string(output))
}

func getLinuxCPUModel() string {
	file, err := os.Open("/proc/cpuinfo")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		params := strings.Split(text, ":")
		if strings.TrimSpace(params[0]) == "model name" {
			return strings.TrimSpace(params[1])
		}
	}
	return UNKNOW
}

func getWindowsCPUModel() string { return "" }
