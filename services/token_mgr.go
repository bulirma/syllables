package services

import (
	//"fmt"
	"bytes"
	"os/exec"
	"time"
)

const (
	tokenMgrCmd = "./token_mgr.sh"

	removeTokenOption = "--remove"
	expirationDateOption = "--info"
)

type TokenManager struct {
	DbFile string
}

// Token is invalid, when it's not present in database or
// the validity date has expired.
func (tokenMgr *TokenManager) IsTokenValid(token string) bool {
	cmd := exec.Command(tokenMgrCmd, expirationDateOption, token, tokenMgr.DbFile)
	var output bytes.Buffer
	cmd.Stdout = &output
	// todo: error handling
	cmd.Run()
	if output.Len() == 0 {
		return false
	}
	location, _ := time.LoadLocation("UTC")
	now := time.Now().In(location)
	dateOutput := string(output.String()[:output.Len() - 1])
	date, err := time.ParseInLocation(time.RFC3339, dateOutput, location)
	if err != nil {
		return false
	}
	if now.After(date) {
		return false
	}
	return true
}

func (tokenMgr *TokenManager) RemoveToken(token string) {
	cmd := exec.Command(tokenMgrCmd, removeTokenOption, token, tokenMgr.DbFile)
	cmd.Run()
}
