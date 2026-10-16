package api

type ConfigStruct struct {
	CORS            bool
	AutoKeys	bool
	ApiPrivateKey [32]byte
	ApiPublicKey  [32]byte
}
