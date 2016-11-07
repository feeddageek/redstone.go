package main

import (
	"fmt"
	"github.com/feeddageek/redstone.go/minecraft"
	"log"
	"net/http"
)

var i *minecraft.Instance
var jar minecraft.Jar
var world minecraft.World

func ressourceHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a handler that serve static ressources
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a mechanism to handle authentication
}

func redstoneHandler(w http.ResponseWriter, req *http.Request) {
	if i!=nil  && i.Running(){
		w.Write([]byte("The server is running.\n"))
	} else {
		var err error
		i,err = minecraft.Start(jar,world)
		if err != nil {
			w.Write([]byte("An error has occured.\n"))
			fmt.Fprintf(w, "Not started : %s", err.Error())
		} else {
			w.Write([]byte("Started.\n"))
		}
	}
}

func stopHandler(w http.ResponseWriter, req *http.Request) {
	if err := i.Stop(); err != nil {
		fmt.Fprintf(w, "Not stopped : %s", err.Error())
	} else {
		w.Write([]byte("Stopped.\n"))
	}
}

func main() {
	jar = minecraft.NewJar("minecraft_server.jar","-Xmx1700M", "-Xms1700M")
	world = minecraft.NewWorld("./minecraft/")
	http.HandleFunc("/res/", ressourceHandler)
	http.HandleFunc("/redstone/", redstoneHandler)
	log.SetFlags(0)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}
