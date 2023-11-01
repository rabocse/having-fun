
# Having Fun

## Simple Golang Web Server

Find below the web server source code:

```go
package main

import (
        "fmt"
        "net/http"
)

func main() {
        http.HandleFunc("/tunnel", func(w http.ResponseWriter, r *http.Request) {
                for name, values := range r.Header {
                        for _, value := range values {
                                fmt.Fprintf(w, "%v: %v\n", name, value)
                        }
                }
        })

        err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
        if err != nil {
                fmt.Println(err)
        }
}
```


A test via CLI (curl) is performed and it can be observed how the web server returns the headers that are being passed in the HTTP Request.
```bash
curl -H 'Having-Fun: Oh-yes!' -H 'Hired: Hopefully!' http://localhost:8080/tunnel
Having-Fun: Oh-yes!
Hired: Hopefully!
User-Agent: curl/8.1.2
Accept: */*
```

Likewise, when a HTTP Request is sent via a browser, this is how it looks:
```bash
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36 Edg/118.0.2088.76
Accept-Language: en-US,en;q=0.9
Connection: keep-alive
Sec-Fetch-Site: none
Accept-Encoding: gzip, deflate, br
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Sec-Fetch-Mode: navigate
Sec-Fetch-User: ?1
Sec-Ch-Ua: "Chromium";v="118", "Microsoft Edge";v="118", "Not=A?Brand";v="99"
Sec-Ch-Ua-Mobile: ?0
Sec-Ch-Ua-Platform: "macOS"
Sec-Fetch-Dest: document
``` 