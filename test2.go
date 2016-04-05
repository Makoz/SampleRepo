package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	// "os/exec"
)

type CFReturnVal struct {
	JsonArgString  string // Json struct htat'll be passed to next hop
	Ip             string
	FunName        string
	ServiceFunName string
	CFInfo         ChainingFunctionInfo
	ReturnToOrigin bool
}

type ChainingFunctionInfo struct {
	GitRepo       string
	RepoName      string
	FileName      string
	CFName        string
	DebuggingPort string
	ClientIpPort  string
}

type ValReply struct {
	Reply string
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

	argNext, _ := json.Marshal(args)

	var returnVal CFReturnVal
	returnVal.Ip = "localhost:4000"
	returnVal.FunName = "1"
	returnVal.JsonArgString = string(argNext)
	returnVal.ServiceFunName = "SampleServ.Reverse"
	returnVal.CFInfo = cfInfo
	buff, _ := json.Marshal(returnVal)
	str := string(buff)
	fmt.Println(str)
}

func func2() {

	fmt.Println("func2 from test2")
}

func formatJsonStringInput(str string) string {
	str = strings.Replace(jsonString, "\\", "", -1)
	last := len(jsonString) - 1
	str = str[1:last]
	return str
}


func main() {

	usage := fmt.Sprintf("Usage: %s CF Id\n", os.Args[0])
	if len(os.Args) != 3 {
		fmt.Printf(usage)
		os.Exit(1)
	}
	jsonString := os.Args[1]
	jsonString = strings.Replace(jsonString, "\\", "", -1)
	last := len(jsonString) - 1
	jsonString = jsonString[1:last]
	// jsonString = strings.Replace(jsonString, "\\", "", -1)

	cfInfoOrigString := os.Args[2]
	// fmt.Println("jsonstring is1234: ", jsonString)
	var vReply ValReply
	err := json.Unmarshal([]byte(jsonString), &vReply)
	if err != nil {
		fmt.Println("ERROR 123!!: ", err)
		return
		// fmt.Println(err)
	}
	// fmt.Println(vReply)

	var origCFInfo ChainingFunctionInfo
	// fmt.Println("input CF: ", cfInfoOrigString)
	cfInfoOrigString = strings.Replace(cfInfoOrigString, "\\", "", -1)
	last = len(cfInfoOrigString) - 1
	cfInfoOrigString = cfInfoOrigString[1:last]
	// fmt.Println("input CF: ", cfInfoOrigString)

	err = json.Unmarshal([]byte(cfInfoOrigString), &origCFInfo)
	if err != nil {
		fmt.Println("cf info, ", err)
		return
	}
	var cfInfo ChainingFunctionInfo
	cfInfo.GitRepo = "https://github.com/Makoz/SampleRepo.git"
	cfInfo.FileName = "test3.go"
	cfInfo.CFName = "2"
	cfInfo.RepoName = "SampleRepo"
	var args GetArgs
	args.Key = vReply.Reply
	// fmt.Println("args is: ", args)
	b, err := json.Marshal(args)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	var returnVal CFReturnVal
	returnVal.Ip = "localhost:4000"
	returnVal.FunName = "1"
	returnVal.JsonArgString = string(b)
	// fmt.Println("string b", returnVal.JsonArgString)
	// returnVal.FunArgs = args
	returnVal.ReturnToOrigin = false
	returnVal.ServiceFunName = "SampleServ.Reverse"
	returnVal.CFInfo = cfInfo
	buff, err := json.Marshal(returnVal)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(string(buff))
	}
	// str := string(buff)
	// fmt.Println(string(buff))

	// cfId := os.Args[1]
	// if cfId == "1" {
	//  func1()
	// } else {
	//  func2()
	// }

}
