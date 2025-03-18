package core

import "runtime"

func GetSysNam() string {
	return runtime.GOOS
}
