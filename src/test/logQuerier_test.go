package test

import(
	"testing"
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)


func TestLogQuerier(t *testing.T){
	fileName := "output.txt"
	//execute the client
	LogQuerier()
	//Open the result file and read each line
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
				fmt.Println(err)
				return
	}
	rd := bufio.NewReader(file)
	//Check whether the result is correct
	i := 1
	flag := 0
	for {
		line, _:= rd.ReadString('\n')
		if i > 23 {
			break
		}
		line = strings.Replace(line, "\n", "", -1)
		line_int,_ := strconv.Atoi(line)
	  if i <= 10 && line_int != i {
			flag = 1
			break
		}
		if i > 10 && i < 21 && line_int != 100 {
			flag = 1
			break
		}
		if i == 21 && line_int != 1 {
			flag = 1
			break
		}
		if i == 22 && line_int != 1 {
			flag = 1
			break
		}
		if i == 23 && line_int != 2 {
			flag = 1
			break
		}
		i++
	}
	if flag == 1 {
		t.Error("Failed! \n" )
	}
}
