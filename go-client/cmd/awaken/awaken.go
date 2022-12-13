package main

import (
	"encoding/base64"
	"encoding/json"
	"go-client/pkg/awaken"
	"log"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		a := parse(os.Args[1])
		a.Run()
	} else {
		log.Printf("usage: %v kiwi://<base64>", os.Args[0])
	}

}

func parse(arg1 string) *awaken.Argument {
	u, err := url.Parse(arg1)
	if err != nil {
		log.Fatalf("args parse error:%v", err.Error())
	}
	// remove xxx://
	base64String := u.Hostname()

	// decode
	decoded, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		log.Fatalf("args parse error:%v", err.Error())
	}

	// unmarshal
	a := &awaken.Argument{}
	if err = json.Unmarshal(decoded, a); err != nil {
		log.Fatalf("args parse error:%v", err.Error())
	}

	return a
}

//go build -o KiwiClient awaken.go
//CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-H windowsgui" -o KiwiClient.exe awaken.go
//CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o KiwiClient awaken.go
