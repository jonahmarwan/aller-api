package api

/*
#cgo LDFLAGS: -L./target/release -lsec
#include <stdlib.h>

void gen_key(unsigned char* priv, unsigned char* pub);
*/
import "C"
import "unsafe"

type ConfigStruct struct {
	CORS            bool
	API_PRIVATE_KEY [32]byte
	API_PUBLIC_KEY  [32]byte
}

func GenerateKeys() ([32]byte, [32]byte) {
	var privateKey [32]byte
	var publicKey [32]byte
	C.gen_key(
		(*C.uchar)(unsafe.Pointer(&privateKey[0])),
		(*C.uchar)(unsafe.Pointer(&publicKey[0])),
	)

	return privateKey, publicKey
}
