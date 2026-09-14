//go:build linux

package crypt

/*
#cgo LDFLAGS: -L../target/release -lsec
#include <stdlib.h>

void gen_key(unsigned char* priv, unsigned char* pub);
*/
import "C"
import "unsafe"

func GenerateKeys() ([32]byte, [32]byte) {
	var privateKey [32]byte
	var publicKey [32]byte
	C.gen_key(
		(*C.uchar)(unsafe.Pointer(&privateKey[0])),
		(*C.uchar)(unsafe.Pointer(&publicKey[0])),
	)

	return privateKey, publicKey
}
