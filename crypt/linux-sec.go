//go:build linux

package crypt

/*
#cgo LDFLAGS: -L${SRCDIR}/../lib/linux_amd64 -lsec
#include <stdlib.h>

void gen_key(unsigned char* priv, unsigned char* pub);
*/
import "C"
import "unsafe"
import "errors"

func GenerateKeys() ([32]byte, [32]byte) {
	var privateKey [32]byte
	var publicKey [32]byte
	C.gen_key(
		(*C.uchar)(unsafe.Pointer(&privateKey[0])),
		(*C.uchar)(unsafe.Pointer(&publicKey[0])),
	)
	if (privateKey == [32]byte{0} || publicKey == [32]byte{0}) {
		return privateKey, publicKey, errors.New("Failed to generate API keys.")
	}

	return privateKey, publicKey
}
