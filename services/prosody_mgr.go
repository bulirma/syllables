package services

import (
	//"fmt"
	"os/exec"

	"github.com/bulirma/syllables/config"
)

const (
	prosodyCtlCmd = "prosodyctl"

	registerSubCmd = "register"
)

func ProsodyRegister(username, password string) bool {
	cmd := exec.Command(prosodyCtlCmd, registerSubCmd, username, config.Host, password)
	err := cmd.Run()
	return err == nil
}
