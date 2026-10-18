package main

import (
	api "github.com/jonahmarwan/aller-api"
	"github.com/jonahmarwan/aller-api/multi"
	"github.com/jonahmarwan/aller-api/route"
	"fmt"
	"log"
	"net/http"
	"strings"
	// "log"
	// "net/http"
)

var EnableCORS = true

func main() {

	// Configuration
	Config := api.ConfigStruct{CORS: EnableCORS, AutoKeys: true}

	/*
		Alternatively:

		privateKey, publicKey, err := crypt.GenerateKeys()
		if err != nil {
			//handle error
		}
		Config := api.ConfigStruct{CORS: EnableCORS, ApiPrivateKey: privateKey, ApiPublicKey: publicKey}

	*/

	// Init router
	router, err := route.NewRouter(&Config)
	router.EnableCORS()
	
	if err != nil {
		panic(err)
	}

	// Init server

	server, err := NewServer(&Config, &router)

	if err != nil {
		panic(err)
	}

	// Init polyglot
	polyglot := multi.NewPolyglot()
	err = polyglot.ReadServices()

	if err != nil {
		panic(err)
	}


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

	fmt.Println(router.Config.ApiPublicKey)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", server.Self()))
}
