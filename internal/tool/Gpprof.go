package utils

import (
	"os"
	"runtime/pprof"
	"sago/internal/logger"
)

func CpuPprof() {
	filename := GetCurrentPath() + "/cpu.pprof"

	if f, err := os.Create(filename); err != nil {
		logger.Error(err)
	} else {
		if err := pprof.StartCPUProfile(f); err != nil {
			logger.Error(err)
		} else {
			defer pprof.StopCPUProfile()
			defer func() {
				if err := f.Close(); err != nil {
					logger.Error(err)
				}
			}()
		}
	}

}
