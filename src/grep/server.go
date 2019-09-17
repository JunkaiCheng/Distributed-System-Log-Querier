package main

import (
        "log"
        "fmt"
        "net"
        "net/http"
        "net/rpc"
        "../types"
)


func main(){
	var err error
    idx := new(types.Idx)
    rpc.Register(idx)
    rpc.HandleHTTP()
    l, e := net.Listen("tcp", ":1111")
    if e != nil{
        log.Fatal ("server listen error: ",e)
    }
    log.Printf("Listening...\n")
    err = http.Serve(l,nil)
	if err != nil{
		log.Fatal("reply error: ",err)
	}

    //reader := bufio.NewReader(os.Stdin)
    //fmt.Print("-> %s",a+b)
    //text,_ := reader.ReadString('\n')
    fmt.Printf("should never get here")
}
