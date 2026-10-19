package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	api "github.com/jonahmarwan/aller-api"
	serve "github.com/jonahmarwan/aller-api/server"
	"github.com/jonahmarwan/aller-api/multi"
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
	router, err := serve.NewRouter(&Config)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	
	router.EnableCORS()
	

	// Init server

	server, err := serve.NewServer(&Config, &router)

	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	// Init polyglot

	os.Setenv("API_SERVICES", "../../services") //set environment variable for service folder
	polyglot, err := multi.NewPolyglot()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}
	
	err = polyglot.ReadServices()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = polyglot.ExecServices()

	if err != nil {
		log.Fatal(err.Error())
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
	router.AddRoute("/api/idiot", Idiothandler)
	router.AddRoute("/api/users", Usershandler)

	fmt.Println(router.Config.ApiPublicKey)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", server.Self()))
}



func Usershandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("dfdsfd"))
}

func Idiothandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("fuck u"))
}