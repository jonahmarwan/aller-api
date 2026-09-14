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
	// Config.API_PRIVATE_KEY, Config.API_PUBLIC_KEY = crypt.GenerateKeys()

	var polyglot multi.Polyglot
	polyglot.InitPolyglot()
	polyglot.ReadServices()
	polyglot.PathTree.Walk(func(s string, v interface{}) bool {
		fmt.Printf("Key: %s, Value: %v\n", s, v)
		return false // return true to stop the walk early
	})
	meta, err := multi.ReadMeta("../services/python/db.py")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(meta)
	// log.Fatal(http.ListenAndServe("127.0.0.1:80", router.Self()))
}
