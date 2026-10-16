//go:build windows

package crypt

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/windows_amd64 -lsec -lntdll -lws2_32 -luserenv
#include <stdlib.h>

void gen_key(unsigned char* priv, unsigned char* pub);
*/
import "C"
import (
	"errors"
	"unsafe"
)

func GenerateKeys() ([32]byte, [32]byte, error) {
	var privateKey [32]byte
	var publicKey [32]byte
	C.gen_key(
		(*C.uchar)(unsafe.Pointer(&privateKey[0])),
		(*C.uchar)(unsafe.Pointer(&publicKey[0])),
	)
	if (privateKey == [32]byte{0} || publicKey == [32]byte{0}) {
		return privateKey, publicKey, errors.New("Failed to generate API keys.")
	}

	return privateKey, publicKey, nil
}
