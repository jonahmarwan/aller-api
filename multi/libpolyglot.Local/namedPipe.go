//go:build windows

package multi

import (
	"net"
)

func hi[T any](test T) T{
    net.Listen("unix", "./hisocket")
    return test 
}



