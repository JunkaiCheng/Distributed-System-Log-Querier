package test

import (
		"fmt"
		"net/rpc"
		"../types"
		"log"
		"os"
		"strconv"
)



func LogQuerier(){
	var err error
	var reply types.Response
	var client *rpc.Client
	var error error
	ans := ""
//===================================
//rare and somewhat frequent patterns
  // execute on each VM
	for i := 0; i < 10;i++{
		//get the VM ID and file name
 		var vmID string
		if i < 9 {
			vmID = "0" + strconv.Itoa(i + 1)
		} else {
			vmID = strconv.Itoa(i + 1)
		}
		text := "grep -c 425 " + "../unit_test" + vmID + ".log\n"
		client,error = rpc.DialHTTP("tcp", types.Addr_list[i]+":1111")
		if error != nil {
				fmt.Print("VM ",i," is offline\n")
		} else {
			err = client.Call("Idx.Socket",text,&reply)
			ans = ans + reply.Foo
			reply.Foo = ""
			if err != nil {
				log.Fatal("response Fatal Error: ",err)
			}
		}
	}


	//===================================
	//frequent patterns
	// execute on each VM
	for i := 0; i < 10;i++{
		//get the VM ID and file name
		var vmID string
		if i < 9 {
			vmID = "0" + strconv.Itoa(i + 1)
		} else {
			vmID = strconv.Itoa(i + 1)
		}
		text := "grep -c frequent " + "../unit_test" + vmID + ".log\n"
		client,error = rpc.DialHTTP("tcp", types.Addr_list[i]+":1111")
		if error != nil {
				fmt.Print("VM ",i," is offline\n")
		} else {
			err = client.Call("Idx.Socket",text,&reply)
			ans = ans + reply.Foo
			reply.Foo = ""
			if err != nil {
				log.Fatal("response Fatal Error: ",err)
			}
		}
	}


	//===================================
	//pattern occurring in one log
	// execute on each VM
	for i := 0; i < 10;i++{
		//get the VM ID and file name
		var vmID string
		if i < 9 {
			vmID = "0" + strconv.Itoa(i + 1)
		} else {
			vmID = strconv.Itoa(i + 1)
		}
		text := "grep -c VM010203 " + "../unit_test" + vmID + ".log\n"
		client,error = rpc.DialHTTP("tcp", types.Addr_list[i]+":1111")
		if error != nil {
				fmt.Print("VM ",i," is offline\n")
		} else {
			err = client.Call("Idx.Socket",text,&reply)
			ans = ans + reply.Foo
			reply.Foo = ""
			if err != nil {
				log.Fatal("response Fatal Error: ",err)
			}
		}
	}




	//===================================
	//pattern occurring in some logs
	// execute on each VM
	for i := 0; i < 10;i++{
		//get the VM ID and file name
		var vmID string
		if i < 9 {
			vmID = "0" + strconv.Itoa(i + 1)
		} else {
			vmID = strconv.Itoa(i + 1)
		}
		text := "grep -c VM01 " + "../unit_test" + vmID + ".log\n"
		client,error = rpc.DialHTTP("tcp", types.Addr_list[i]+":1111")
		if error != nil {
				fmt.Print("VM ",i," is offline\n")
		} else {
			err = client.Call("Idx.Socket",text,&reply)
			ans = ans + reply.Foo
			reply.Foo = ""
			if err != nil {
				log.Fatal("response Fatal Error: ",err)
			}
		}
	}

	//Output the result
	f, _ := os.Create("output.txt")
	_,er:=f.WriteString(ans)
	if er != nil{
		fmt.Println(err)
		f.Close()
		return
	}



	f.Close()
	return;
}
