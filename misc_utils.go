package dss

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

// GetMiscProperties get queue connection parameters
func GetMiscProperties(url, platKey, svcCode, profile, userAgent string) (map[string]MiscProperties, error) {
	if err := CheckParam(url, platKey, svcCode, profile, userAgent); err != nil {
		return nil, err
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
	resObj := make(map[string]map[string]MiscProperties)
	err = yaml.Unmarshal(bodyBytes, resObj)
	if err != nil {
		// It indicates that the server does not return misc.yml file normally, either the PLATKEY is wrong,
		// or the SVCCODE is wrong, or the PROFILE is wrong.
		// If there is no problem in self-examination, contact the data source management system administrator
		return nil, fmt.Errorf("%s: %s", err.Error(), string(bodyBytes))
	}
	return resObj["misc"], nil
}