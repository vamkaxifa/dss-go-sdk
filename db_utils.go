package dss

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
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

// GetDbProperties get database connection parameters
func GetDbProperties(url, platKey, svcCode, profile, userAgent string) (map[string]DsnProperties, error) {
	if strings.Trim(url, " ") == "" {
		return nil, fmt.Errorf("URL can't be empty string")
	}
	if !strings.HasPrefix(url, "https") {
		return nil, fmt.Errorf("Only https protocol is supported ")
	}
	if strings.Trim(platKey, " ") == "" {
		return nil, fmt.Errorf("platKey can't be empty string")
	}
	if strings.Trim(svcCode, " ") == "" {
		return nil, fmt.Errorf("svcCode can't be empty string")
	}
	if strings.Trim(profile, " ") == "" {
		return nil, fmt.Errorf("profile can't be empty string")
	}
	if profile != profileRc1 && profile != profileRc2 && profile != profileProd {
		return nil, fmt.Errorf("profile error")
	}
	prams := make(map[string]string)
	prams["platkey"] = platKey
	prams["svccode"] = svcCode
	prams["profile"] = profile

	if userAgent == "" {
		userAgent = svcCode
	}

	bodyBytes, err := HTTPGet(url, userAgent, prams, timeOutSeconds)
	if err != nil {
		return nil, err
	}
	resObj := make(map[string]map[string]DsnProperties)
	err = yaml.Unmarshal(bodyBytes, resObj)
	if err != nil {
		// It indicates that the server does not return db.yml file normally, either the PLATKEY is wrong,
		// or the SVCCODE is wrong, or the PROFILE is wrong.
		// If there is no problem in self-examination, contact the data source management system administrator
		return nil, fmt.Errorf("%s: %s", err.Error(), string(bodyBytes))
	}
	return resObj["database"], nil
}
