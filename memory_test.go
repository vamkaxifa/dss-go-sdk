package dss

import (
	"testing"
)

func TestGetMemProperties(t *testing.T) {
	const (
		url       = "https://test.example.com/dss/mem"
		platKey   = "xxx"
		svcCode   = "xxx"
		profile   = "xxx"
		userAgent = svcCode
	)

	memObj, err := GetMemProperties(url, platKey, svcCode, profile, userAgent)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	for k, v := range memObj {
		t.Logf("mem-n：%s, server_path: %s, mount_path: %s", k, v.MountPath, v.ServerPath)
	}
}
