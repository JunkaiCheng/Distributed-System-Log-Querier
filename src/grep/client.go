package main

import (
		"fmt"
		"net/rpc"
		"../types"
		"log"
		"os"
		"bufio"
		"strings"
)

func main(){
	var err error
	var reply types.Response
	var client *rpc.Client
	var error error
	var temp []string
	ans := ""
	//get input command
	fmt.Print("Please type in the grep command to run:\n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	fmt.Print("your command is:",text)

	for i:= 0; i<10;i++{
		client,error = rpc.DialHTTP("tcp", types.Addr_list[i]+":1111")
		if error != nil {
				fmt.Print("VM ",i," is offline\n")
		} else {
			err = client.Call("Idx.Socket",text,&reply)
			temp = strings.Split(reply.Foo,"\n")
			fmt.Print("VM ", i, " have ",len(temp)-1, " matches\n")
			ans = ans + reply.Foo
			reply.Foo = ""
			if err != nil {
				log.Fatal("response Fatal Error: ",err)
			}
		}
	}
	f, _ := os.Create("output.txt")
	_,er:=f.WriteString(ans)//l,err possibly
	if er != nil{
		fmt.Println(err)
		f.Close()
		return
	}
	f.Close()
	return;
}
