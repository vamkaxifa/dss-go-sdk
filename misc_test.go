package dss

import (
	"testing"
)

func TestGetMiscProperties(t *testing.T) {
	const (
		url       = "https://test.example.com/dss/queue"
		platKey   = "xx"
		svcCode   = "xx"
		profile   = "xx"
		userAgent = svcCode
	)

	miscObj, err := GetMiscProperties(url, platKey, svcCode, profile, userAgent)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	for k, v := range miscObj {
		t.Logf("mq-n：%s, ip: %s, port: %d, password: %s", k, v.IP, v.Port, v.Password)
	}
}
