package dss

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// HTTPGet HTTP Get
func HTTPGet(url, userAgent string, prams map[string]string, timeOut time.Duration) ([]byte, error) {
	//Skip certificate verification
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	//http cookie
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Timeout:   timeOut * time.Second,
		Jar:       cookieJar,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	for k, v := range prams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
