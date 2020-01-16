package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/json-iterator/go"
)

type Data struct {
	AccountName      string
	Extranonce1      string
	ExtranonceSuffix string
	VersionMask      uint
	Difficulty       float64
	Target           string
	TargetDifficulty float64
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	data := &Data{
		"AccountNameAccountNameAccountNameAccountNameAccountNameAccountName",
		"Extranonce1Extranonce1Extranonce1Extranonce1Extranonce1",
		"ExtranonceSuffixExtranonceSuffixExtranonceSuffixExtranonceSuffixExtranonceSuffixExtranonceSuffix",
		1,
		100000,
		"TargetTargetTargetTargetTargetTarget",
		100000,
	}

	times := 1
	for i := 0; i < 100000; i++ {
		//fmt.Printf("the %d times json marshal and unmarshal at %s\n", times, time.Now().Format("2006/01/02 15:04:05"))

		d, _ := json.Marshal(&data)
		json.Unmarshal(d, data)
		times++

		//time.Sleep(time.Second)
	}
	fmt.Printf("end %d times json marshal and unmarshal at %s\n", times, time.Now().Format("2006/01/02 15:04:05"))

	c := make(chan bool)
	<-c

}
