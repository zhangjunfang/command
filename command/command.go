package command

import (
	"errors"
	"os/exec"
	"strings"
)

const ()

var ()

func init() {

}

func ParseCommand(command string) (cmd *exec.Cmd, e error) {
	command = strings.TrimSpace(command)
	if len(command) > 0 {
		slice := strings.Split(command, " ")
		cmd = exec.Command(slice[0], slice[1:]...)
	} else {
		return nil, errors.New(" Command execution error ")
	}
	return cmd, e
}
func RemoveMiddleSpace(s string) (ss []string) {
	s = strings.TrimSpace(s)
	slice := strings.Split(s, " ")
	for _, v := range slice {
		if v != "" {
			ss = append(ss, v)
		}
	}
	return ss
}
