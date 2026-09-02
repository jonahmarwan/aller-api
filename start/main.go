package main

import (
	api "aller-api"
	"aller-api/multi"
	"aller-api/route"
	"fmt"
	// "log"
	// "net/http"
)

var EnableCORS = true
var Config = api.ConfigStruct{CORS: EnableCORS, API_PRIVATE_KEY: [32]byte{}, API_PUBLIC_KEY: [32]byte{}}

func main() {
	var router route.Router
	router.InitRouter(&Config)
	router.EnableCORS()
	router.AddRoute("/api/idiot", route.Idiothandler)

	Config.API_PRIVATE_KEY, Config.API_PUBLIC_KEY = api.GenerateKeys()

	var polyglot multi.Polyglot
	polyglot.InitPolyglot()
	polyglot.ReadServices()
	polyglot.PathTree.WalkPrefix("../services/", func(s string, v interface{}) bool {
		fmt.Println(s, v)
		return false
	})

	fmt.Println(Config.API_PRIVATE_KEY, Config.API_PUBLIC_KEY)
	// log.Fatal(http.ListenAndServe("localhost:8080", router.Self()))
}
