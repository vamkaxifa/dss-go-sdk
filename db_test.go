package dss

import (
	"testing"
)

func TestGetDbProperties(t *testing.T) {
	const (
		url       = "https://test.example.com/dss/db"
		platKey   = "xxx"
		svcCode   = "xxx"
		profile   = "xxx"
		userAgent = svcCode
	)

	dsnObj, err := GetDbProperties(url, platKey, svcCode, profile, userAgent)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	for k, v := range dsnObj {
		t.Logf("DSN-N：%s, ip: %s, port: %d, sid: %s, username: %s, password: %s", k, v.IP, v.Port, v.Sid, v.UserName, v.Password)
	}
}
