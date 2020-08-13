package dss

import (
	"testing"
)

func TestGetQueueProperties(t *testing.T) {
	const (
		url       = "https://test.example.com/dss/queue"
		platKey   = "xx"
		svcCode   = "xx"
		profile   = "xx"
		userAgent = svcCode
	)

	queueObj, err := GetQueueProperties(url, platKey, svcCode, profile, userAgent)
	if err != nil {
		t.Errorf(err.Error())
		t.FailNow()
	}
	for k, v := range queueObj {
		t.Logf("mq-n：%s, ip: %s, port: %d, vhost: %s, username: %s, password: %s", k, v.IP, v.Port, v.VHost, v.UserName, v.Password)
	}
}
