package awaken

import (
	"fmt"
	"path/filepath"
	"strconv"

	"os/exec"
)

// 将来如果支持其他终端软件，在这个函数加一个参数用来做分支判断
func handleSSH(cwd, username, ip, password string, port int) *exec.Cmd {
	command := puttySSH(username, ip, password, port)
	exepath := filepath.Join(cwd, "putty.exe")
	return exec.Command(exepath, command...)
}

// func xshellSSH(c string) string {
// 	return c
// }

// putty.exe -ssh KIWI-544254ac-d8ef-49fa-bfe4-f0f002fe3bfe@192.168.88.10 -P 2222 -pw empty
func puttySSH(username, ip, password string, port int) []string {
	// return fmt.Sprintf("%v -ssh %v@%v -P %v -pw %v", exepath, username, ip, port, password)
	return []string{
		"-ssh", fmt.Sprintf("%v@%v", username, ip),
		"-P", strconv.Itoa(port),
		"-pw", password,
	}
}
