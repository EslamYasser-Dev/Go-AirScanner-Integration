package scanner

/*
#cgo CXXFLAGS: -std=c++11 -I./twain-dsm/Include
#cgo LDFLAGS: -L./lib -lTWAINDSM -lole32 -loleaut32 -luser32
#include "twain_wrapper.h"
*/
import "C"
import (
	"fmt"
)

func Init() error {
	if C.twain_init() != 1 {
		return fmt.Errorf("failed to init TWAIN DSM")
	}
	return nil
}

func SelectScanner() error {
	if C.twain_select_source() != 1 {
		return fmt.Errorf("failed to select scanner")
	}
	return nil
}

func OpenSource() error {
	if C.twain_open_source() != 1 {
		return fmt.Errorf("failed to open source")
	}
	return nil
}

func SetADF(enable bool) {
	enableInt := 0
	if enable {
		enableInt = 1
	}
	C.twain_enable_adf(C.int(enableInt))
}

func SetDPI(dpi int) {
	C.twain_set_dpi(C.int(dpi))
}

func SetColor(color bool) {
	mode := 0
	if color {
		mode = 1
	}
	C.twain_set_color(C.int(mode))
}

func Close() {
	C.twain_exit()
}
