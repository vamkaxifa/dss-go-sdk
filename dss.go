package dss

import (
	"fmt"
	"strings"
)

const (
	//timeOutSeconds 3s
	timeOutSeconds = 3
	//profileRc1 test1
	profileRc1 = "rc1"
	//profileRc2 test2
	profileRc2 = "rc2"
	//profileProd prod
	profileProd = "prod"
)

// CheckParam
func CheckParam(url, platKey, svcCode, profile, userAgent string) error {
	if strings.Trim(url, " ") == "" {
		return fmt.Errorf("URL can't be empty string")
	}
	if !strings.HasPrefix(url, "https") {
		return fmt.Errorf("Only https protocol is supported ")
	}
	if strings.Trim(platKey, " ") == "" {
		return fmt.Errorf("platKey can't be empty string")
	}
	if strings.Trim(svcCode, " ") == "" {
		return fmt.Errorf("svcCode can't be empty string")
	}
	if strings.Trim(profile, " ") == "" {
		return fmt.Errorf("profile can't be empty string")
	}
	if profile != profileRc1 && profile != profileRc2 && profile != profileProd {
		return fmt.Errorf("profile error")
	}
	return nil
}
