package main

import (
	"log"
	"fmt"
	"strconv"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/feeddageek/redstone.go/minecraft"
)

type Server struct {
	Jar	minecraft.Jar
	World	minecraft.World
	Port	uint16
	i	*minecraft.Instance
}

var s Server

//THE ROAD TO WISDOM
//Err and err and err again,
//but less and less and less.
func r2w(err error){
	if err != nil {
		log.Fatal(err)
	}
}

func ressourceHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a handler that serve static ressources
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	//TODO add a mechanism to handle authentication
}

func redstoneHandler(w http.ResponseWriter, req *http.Request) {
	if s.i!=nil  && s.i.Running(){
		w.Write([]byte("The server is running.\n"))
	} else {
		var err error
		s.i,err = minecraft.Start(s.Jar,s.World)
		if err != nil {
			w.Write([]byte("An error has occured.\n"))
			fmt.Fprintf(w, "Not started : %s", err.Error())
		} else {
			w.Write([]byte("Started.\n"))
		}
	}
}

func stopHandler(w http.ResponseWriter, req *http.Request) {
	if err := s.i.Stop(); err != nil {
		fmt.Fprintf(w, "Not stopped : %s", err.Error())
	} else {
		w.Write([]byte("Stopped.\n"))
	}
}

func main() {
	f,err:= ioutil.ReadFile("redstone.json")
	r2w(err)
	err = json.Unmarshal(f,&s)
	r2w(err)
	http.HandleFunc("/res/", ressourceHandler)
	http.HandleFunc("/redstone/", redstoneHandler)
	log.SetFlags(0)
	err = http.ListenAndServe(":"+strconv.Itoa(int(s.Port)), nil)
	r2w(err)
}
