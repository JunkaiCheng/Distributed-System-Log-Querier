package types

import (
        "fmt"
		"os/exec"
		//"log"
		"strings"
)

type Response struct{
	Foo string
}

func init(){
    fmt.Print("Types Package init...\n")
}
type Idx int

func (t *Idx) Socket(request string, response *Response) error {
	var s Response
    fmt.Print("received request: ",request)
	x := strings.Split(request," ")
	length := len(x)-1
	x[length] = strings.Trim(x[length],"\n")
	cmd := exec.Command (x[0],x[1:]...)
	out,err := cmd.CombinedOutput()
	if err != nil{
		//fmt.Print(fmt.Sprint(err) +":" + string(out))
		fmt.Print("No match\n")
		s.Foo = ""
		*response = s
		return nil
	}
	s.Foo = string(out)
    *response = s
    return nil
}
var Addr_list = [10]string{"172.22.152.196","172.22.154.192","172.22.156.192","172.22.152.197","172.22.154.193","172.22.156.193","172.22.152.198","172.22.154.194","172.22.156.194","172.22.152.199"}
