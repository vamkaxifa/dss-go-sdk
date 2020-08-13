# dss-go-sdk

dss-go-sdk is a go sdk for data source management system. 

1.Install 

`go get github.com/vamkaxifa/dss-go-sdk`


2.Usage

2.1 database

```go
import (
    "fmt"
    "github.com/vamkaxifa/dss-go-sdk"
)
func test(){
   const (
        url     = "https://test.example.com/dss/db"
        platKey = "xxxx"
        svcCode = "xxxx"
        profile = "xxx"
        userAgent = svcCode
   	)
    dsnObj, err := dss.GetDbProperties(url, platKey, svcCode, profile, userAgent)
    if err != nil {
        panic(err)
    }
    for k, v := range dsnObj {
        fmt.Printf("DSN-N：%s, ip: %s, port: %d, sid: %s, username: %s, password: %s", k, v.IP, v.Port, v.Sid,v.UserName, v.Password)
    }
}
```

2.2 queue

```go
import (
    "fmt"
    "github.com/vamkaxifa/dss-go-sdk"
)
func test(){
   const (
        url     = "https://test.example.com/dss/queue"
        platKey = "xxxx"
        svcCode = "xxxx"
        profile = "xxx"
        userAgent = svcCode
   	)
   queueObj, err := GetQueueProperties(url, platKey, svcCode, profile, userAgent)
   	if err != nil {
   		panic(err)
   	}
   	for k, v := range queueObj {
   		t.Logf("mq-n：%s, ip: %s, port: %d, vhost: %s, username: %s, password: %s", k, v.IP, v.Port, v.VHost, v.UserName, v.Password)
   	}
}
```

2.1 database

```go
import (
    "fmt"
    "github.com/vamkaxifa/dss-go-sdk"
)
func test(){
   const (
        url     = "https://test.example.com/dss/misc"
        platKey = "xxxx"
        svcCode = "xxxx"
        profile = "xxx"
        userAgent = svcCode
   	)
    miscObj, err := GetMiscProperties(url, platKey, svcCode, profile, userAgent)
    if err != nil {
        panic(err)
    }
    for k, v := range miscObj {
        t.Logf("mq-n：%s, ip: %s, port: %d, password: %s", k, v.IP, v.Port, v.Password)
    }
}
```
