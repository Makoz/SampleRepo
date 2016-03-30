package main

import (
	"encoding/json"
	"fmt"
	"os"
	// "os/exec"
)

type CFReturnVal struct {
	FunArgs        interface{}
	Ip             string
	FunName        string
	ServiceFunName string
	CFInfo         ChainingFunctionInfo
}

type ChainingFunctionInfo struct {
	GitRepo  string
	RepoName string
	FileName string
	CFName   string
}

type GetArgs struct {
	Key string
}

func func1() {

	var cfInfo ChainingFunctionInfo
	cfInfo.GitRepo = "https://github.com/Makoz/SampleRepo.git"
	cfInfo.FileName = "test2.go"
	cfInfo.CFName = "2"
	cfInfo.RepoName = "SampleRepo"
	var args GetArgs
	args.Key = "hi"
	var returnVal CFReturnVal
	returnVal.Ip = "localhost:4000"
	returnVal.FunName = "1"
	returnVal.FunArgs = args
	returnVal.ServiceFunName = "SampleServ.Reverse"
	returnVal.CFInfo = cfInfo
	buff, _ := json.Marshal(returnVal)
	str := string(buff)
	fmt.Println(str)
}

func func2() {

	fmt.Println("func2 from test2")
}

func main() {

	usage := fmt.Sprintf("Usage: %s CF Id\n", os.Args[0])
	if len(os.Args) != 2 {
		fmt.Printf(usage)
		os.Exit(1)
	}

	cfId := os.Args[1]
	if cfId == "1" {
		func1()
	} else {
		func2()
	}

}
