# Cloudflare Tunnel
This module is for running cloudflare tunnels in the background along with other go code like a http server

## Example
In this example the program creates a tunnel that will route example.com to http://localhost:8080
```go
    package main

    import (
        "github.com/jeff9014223/tunnel"
    )

    func main() {
        tunnl, err := tunnel.New("CLOUDFLARE_TUNNEL_TOKEN", true) // set to false for no logging
        if err != nil {
            panic(err)
        }

        err = tunnl.Start()
        if err != nil {
            panic(err)
        }
    }
```
