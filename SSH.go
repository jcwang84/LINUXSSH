package main

import (
	"bufio"
	"bytes"
	"fmt"
	//"log"
	"os"
	"os/exec"
	"strings"
	//"reflect" // for check type
	//"time"
)

func main() {
	var hosts string
	var i int = 0

	file, err := os.Open("hosts.txt")
	if err != nil {
		fmt.Print(err)
	}
	defer file.Close()
	// variable := readLine()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		hosts = strings.Trim(scanner.Text(), " ")
		output := ExecuteCMD(i, hosts)
		fmt.Println(output)
	}
	//fmt.Println(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hit enter to close the window...")
	reader.ReadString('\n')

}

func ExecuteCMD(i int, cmdstring ...string) string {
	cmdname := [...]string{"SSH","ping","systeminfo"}
	
	cmd := exec.Command(cmdname[i], strings.Join(cmdstring, " "))
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
    fmt.Println(cmd)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	outputText := string(cmdOutput.Bytes())
	//fmt.Println(outputText)
	return outputText
}


// func readLine () string {
// 	fmt.Print("Enter variable: ")
//     var input string
//     fmt.Scanln(&input)
//       return input
// }



func moveFile(oldLoaction,newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

