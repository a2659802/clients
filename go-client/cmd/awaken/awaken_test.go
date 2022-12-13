package main

import (
	"go-client/pkg/awaken"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		arg1 string
	}
	tests := []struct {
		name string
		args args
		want *awaken.Argument
	}{
		{
			name: `{"protocol":"ssh","args":{"ip":"1.1.1.1","port":22,"username":"user","password":"pwd"}}`,
			args: args{arg1: "kiwi://ewoJInByb3RvY29sIjogInNzaCIsCgkiYXJncyI6IHsKCQkiaXAiOiAiMS4xLjEuMSIsCgkJInBvcnQiOiAyMiwKCQkidXNlcm5hbWUiOiAidXNlciIsCgkJInBhc3N3b3JkIjogInB3ZCIKCX0KfQ=="},
			want: &awaken.Argument{
				Protocol: "ssh",
				Args: awaken.SSHArgs{
					IP:       "1.1.1.1",
					Port:     22,
					Username: "user",
					Password: "pwd",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.arg1); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCommand(t *testing.T) {
	exepath := filepath.Join(`C:\Users\sawyer\Desktop\work\clients\go-client\build\windows`, "putty.exe")
	c := exec.Command(exepath, strings.Split("-ssh user@1.1.1.1 -P 22 -pw empty", " ")...)
	err := c.Run()
	if err != nil {
		t.Error(err)
	}
}
