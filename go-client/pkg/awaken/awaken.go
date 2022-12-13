package awaken

import (
	"log"
	"os"
	"path/filepath"
)

type File struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

type SSHArgs struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Argument struct {
	Protocol string  `json:"protocol"`
	Args     SSHArgs `json:"args"`
}

func (a *Argument) HandleSSH() {
	currentPath := filepath.Dir(os.Args[0])

	// generate Cmd
	cmd := handleSSH(currentPath, a.Args.Username, a.Args.IP, a.Args.Password, a.Args.Port)
	// exec
	cmd.Start()
}

func (a *Argument) Run() {
	protocol := a.Protocol
	switch protocol {
	case "ssh":
		a.HandleSSH()
	default:
		log.Fatalf("error protocol:%v", protocol)
	}
}
