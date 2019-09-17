package main

import(
  "fmt"
  "os"
  "math/rand"
  "time"
  "strings"
  "os/exec"
  "strconv"
  "bufio"
)

func ifFileExist(filename string) bool {
//check whether the file exists
  _, err := os.Stat(filename)
	if err == nil {
		return true
	}else{
    return false
  }
}


func removeFile(filename string){
//delete the existed file
  os.Remove(filename)
}


func generateRandomLine() string{
//generate line of random characters with random length
  numOfChar := rand.Intn(50) + 20
  line := make([]byte, numOfChar)
    for i := 0; i < numOfChar; i++ {
        char := rand.Intn(26) + 65
        line[i] = byte(char)
    }
    return string(line) + "\n"
}

func main(){
  //Get the username
  fmt.Print("Please type in your NetID:\n")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
  user := strings.Replace(text, "\n", "", -1)
  //Get the slice for all VM IDs
  VMs := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}

  for _, vmID := range VMs {
    id, _ := strconv.Atoi(vmID)
    fileName := "./unit_test" + vmID + ".log"
    serverName := "fa19-cs425-g58-" + vmID + ".cs.illinois.edu"
    //To check whether the file exsited. If so, remove it.
    if ifFileExist(fileName){
      removeFile(fileName)
    }

    //To create a new file
    file, err := os.Create(fileName)
    defer file.Close()
    if err != nil {
          fmt.Println(err)
    }

    //To generate the log file
    rand.Seed(time.Now().UnixNano())
    numOfLine := rand.Intn(40) + 10
    //frequent pattern
    for i := 0; i < numOfLine; i++ {
      if i == 0 {
        for j := 0; j < 100; j++{
          file.Write([]byte("frequent\n"))
        }
      }
    //rare and somewhat frequent patterns
      if i < id {
      file.Write([]byte("Cs425 Distributed Systems Fall 2019 \n"))
    }
      file.Write([]byte(generateRandomLine()))
    }

    //pattern occurring in one/some logs
    if id == 1{
      file.Write([]byte("VM01\n"))
    }

    if id == 2{
      file.Write([]byte("VM0102\n"))
    }

    if id == 2{
      file.Write([]byte("VM010203\n"))
    }

    //scp to the correponding VM
    command := "scp -o StrictHostKeyChecking=no " + fileName + " " + user + "@" + serverName + ":./mp1"
    list := strings.Split(command, " ")
    cmd := exec.Command(list[0], list[1:]...)
    output, err := cmd.CombinedOutput()
    removeFile(fileName)
    if err != nil {
      fmt.Println(fmt.Sprint(err) + ": " + string(output))
      return
    }
  }
}
