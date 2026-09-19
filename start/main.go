package main

import (
	api "aller-api"
	"aller-api/crypt"
	"aller-api/multi"
	"aller-api/route"
	"fmt"
	"log"
	"net/http"
	"strings"
	// "log"
	// "net/http"
)

var EnableCORS = true
var Config = api.ConfigStruct{CORS: EnableCORS, API_PRIVATE_KEY: [32]byte{}, API_PUBLIC_KEY: [32]byte{}}

func main() {
	// Init router
	var router route.Router
	router.InitRouter(&Config)
	router.EnableCORS()

	//Init API keys
	Config.API_PRIVATE_KEY, Config.API_PUBLIC_KEY = crypt.GenerateKeys()

	// Init polyglot
	var polyglot multi.Polyglot
	polyglot.InitPolyglot()
	polyglot.ReadServices()

	// To be replaced by polyglot comprehension service
	polyglot.PathTree.Walk(func(s string, v interface{}) bool {
		if v == false {
			pathElem := strings.Split(s, "\\")
			if strings.Contains(pathElem[len(pathElem)-1], "_") {
				meta, err := multi.ReadMeta(s)
				if err != nil {
					fmt.Println(err)
					return true
				}
				err = multi.GenerateTOMLfromMeta(meta, s)
				if err != nil {
					fmt.Println(err)
				}

			}
		}
		return false // return true to stop the walk early
	})
	//Add router paths
	router.AddRoute("/api/idiot", route.Idiothandler)
	router.AddRoute("/api/users", route.Usershandler)

	fmt.Println(router.Config.API_PUBLIC_KEY)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router.Self()))
}
