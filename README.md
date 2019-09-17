##Part 1 Query Distributed Log Files
1. ``` cd cs425_grep_project```
1. Need to run server on all vms with <br>```go run src/grep/server.go```
2. run client on any server to perform grep (or other shell command) with <br>```go run src/grep/client.go```
    command to run (eg. grep a my.log)
3. All logfiles are stored as ```～/my.log``` 

Note:<br>
	1) Output will be saved in output.txt on the vm running client<br>
	2) The filename/filepath need to be pre-given by the client

##Part 2 Unit Test
1. Generate log files at every machine:<br>```go run src/test/generateUnitTest.go```<br>You need to enter your netID here
2. Check whether our program works: <br>```go test -v src/test/logQuerier_test.go src/test/logQuerier.go``` <br>If it says "PASS", then the program works.

Note:<br>
	You may store your ssh pub key in ```～/.ssh/authorized_keys``` in each VM to avoid repeating entering passward
	
	
##Contributor

This project is contributed by Junkai Cheng and Tianyi Tang.