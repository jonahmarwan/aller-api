package api

type ConfigStruct struct {
	CORS            bool
	API_PRIVATE_KEY [32]byte
	API_PUBLIC_KEY  [32]byte
}
