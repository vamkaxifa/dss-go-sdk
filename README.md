# dss-go-sdk

dss-go-sdk is a go sdk for data source management system. 

1.Install 

`go get github.com/vamkaxifa/dss-go-sdk`


2.Usage

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
