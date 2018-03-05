// server.go
package main

import (
        "fmt"
        "log"
        "net"
        "net/rpc"
        "net/rpc/jsonrpc"
        "os/exec"
)

type Args struct {
        Container1 string
        Container2 string
}

type ContainerStruct struct{}

func (t *ContainerStruct) CreateContainers(args *Args, reply *string) error {
        fmt.Printf("%s\n", args.Container1)
        *reply = "Started containers: " + args.Container1 + " and " + args.Container2
        container1 := "docker run -d -t -i " + args.Container1
        out, err := exec.Command("/bin/sh", "-c", container1).Output()
        if err != nil {
                panic(err)
        }
        fmt.Printf("Starting container: %s\n", out)

        fmt.Printf("%s\n", args.Container2)
        container2 := "docker run -d -t -i " + args.Container2
        out, err = exec.Command("/bin/sh", "-c", container2).Output()
        if err != nil {
                panic(err)
        }
        fmt.Printf("Starting container: %s\n", out)
        return nil
}

func main() {
        s := new(ContainerStruct)
        server := rpc.NewServer()
        server.Register(s)

        server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
        listener, e := net.Listen("tcp", ":12347")
        if e != nil {
                log.Fatal("listen error:", e)
        }
        for {
                if conn, err := listener.Accept(); err != nil {
                        log.Fatal("accept error: " + err.Error())
                } else {
                        log.Printf("new connection established\n")
                        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
                }
        }
}
