// client.go
package main

import (
        "fmt"
        "log"
        "net"
        "net/rpc/jsonrpc"
)

type Args struct {
        Container1 string
        Container2 string
}

func main() {
	url := "172.31.99.166:12347"
        client, err := net.Dial("tcp", url)
        if err != nil {
                log.Fatal("dialing:", err)
        }
        // Synchronous call

        args := &Args{"ubuntu", "debian"}
        var sreply string
        c := jsonrpc.NewClient(client)
        err = c.Call("ContainerStruct.CreateContainers", args, &sreply)
        if err != nil {
                log.Fatal("arith error:", err)
        }
        fmt.Printf("Result: %s\n", sreply)
}
